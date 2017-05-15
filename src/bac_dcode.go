package bacnet

import "fmt"

/* from clause 20.2.1.2 Tag Number */
/* true if extended tag numbering is used */
func IS_EXTENDED_TAG_NUMBER(x byte) bool {
	return (x & 0xF0) == 0xF0
}

/* from clause 20.2.1.3.1 Primitive Data */
/* true if the extended value is used */
func IS_EXTENDED_VALUE(x byte) bool {
	return (x & 0x07) == 5
}

/* from clause 20.2.1.1 Class */
/* true if the tag is context specific */
func IS_CONTEXT_SPECIFIC(x byte) bool {
	fmt.Printf("%d %d %d   is context specific: %t\n", x, x & BIT3, BIT3, (x & BIT3) == BIT3)
	return (x & BIT3) == BIT3
}

/* from clause 20.2.1.3.2 Constructed Data */
/* true if the tag is an opening tag */
func IS_OPENING_TAG(x byte) bool {
	return (x & 0x07) == 6
}

/* from clause 20.2.1.3.2 Constructed Data */
/* true if the tag is a closing tag */
func IS_CLOSING_TAG(x byte) bool {
	return (x & 0x07) == 7
}

/* from clause 20.1.2.4 max-segments-accepted */
/* and clause 20.1.2.5 max-APDU-length-accepted */
/* returns the encoded octet */
func encode_max_segs_max_apdu(max_segs int, max_apdu uint32) byte {
	var octet byte = 0

	if max_segs < 2 {
		octet = 0
	} else if max_segs < 4 {
		octet = 0x10
	} else if max_segs < 8 {
		octet = 0x20
	} else if max_segs < 16 {
		octet = 0x30
	} else if max_segs < 32 {
		octet = 0x40
	} else if max_segs < 64 {
		octet = 0x50
	} else if max_segs == 64 {
		octet = 0x60
	} else {
		octet = 0x70
	}

	if max_apdu <= 50 {
		octet |= 0x00
	} else if max_apdu <= 128 {
		octet |= 0x01
	} else if max_apdu <= 206 {
		octet |= 0x02
	} else if max_apdu <= 480 {
		octet |= 0x03
	} else if max_apdu <= 1024 {
		octet |= 0x04
	} else if max_apdu <= 1476 {
		octet |= 0x05
	}

	return octet
}

/* from clause 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func encode_tag(apdu []byte, tag_number byte, context_specific bool, len_value_type uint32) int {
	var length int = 1        /* return value */

	apdu[0] = 0
	if context_specific {
		apdu[0] = BIT3
	}

	/* additional tag byte after this byte */
	/* for extended tag byte */
	if tag_number <= 14 {
		apdu[0] |= (tag_number << 4)
	} else {
		apdu[0] |= 0xF0
		apdu[1] = tag_number
		length++
	}

	/* NOTE: additional len byte(s) after extended tag byte */
	/* if larger than 4 */
	if (len_value_type <= 4) {
		apdu[0] |= byte(len_value_type)
	} else {
		apdu[0] |= 5
		length++
		if (len_value_type <= 253) {
			apdu[length] = byte(len_value_type)
		} else if (len_value_type <= 65535) {
			apdu[length] = 254
			length += encode_unsigned16(apdu[length:], uint16(len_value_type))
		} else {
			apdu[length] = 255
			length += encode_unsigned32(apdu[length:], len_value_type)
		}
	}

	return length
}

/* from clause 20.2.1.3.2 Constructed Data */
/* returns the number of apdu bytes consumed */
func encode_opening_tag(apdu []byte, tag_number byte) int {
	var len int = 1

	/* set class field to context specific */
	apdu[0] = BIT3
	/* additional tag byte after this byte for extended tag byte */
	if tag_number <= 14 {
		apdu[0] |= (tag_number << 4)
	} else {
		apdu[0] |= 0xF0
		apdu[1] = tag_number
		len++
	}
	/* set type field to opening tag */
	apdu[0] |= 6

	return len
}

/* from clause 20.2.1.3.2 Constructed Data */
/* returns the number of apdu bytes consumed */
func encode_closing_tag(apdu []byte, tag_number byte) int {
	var len int = 1

	/* set class field to context specific */
	apdu[0] = BIT3
	/* additional tag byte after this byte for extended tag byte */
	if tag_number <= 14 {
		apdu[0] |= (tag_number << 4)
	} else {
		apdu[0] |= 0xF0
		apdu[1] = tag_number
		len++
	}
	/* set type field to closing tag */
	apdu[0] |= 7

	return len
}

func decode_tag_number(apdu []byte) (int, byte) {
	var length int = 1
	var tag_number byte

	if IS_EXTENDED_TAG_NUMBER(apdu[0]) {
		tag_number = apdu[1]
		length++
	} else {
		tag_number = apdu[0] >> 4
	}

	return length, tag_number
}

// value *uint32
func decode_tag_number_and_value(apdu []byte) (int, byte, uint32) {
	var length int = 1
	var len_tmp int = 0
	var value16 uint16 = 0
	var value32 uint32 = 0
	var value uint32 = 0

	var tag_number byte

	length, tag_number = decode_tag_number(apdu)
	if IS_EXTENDED_VALUE(apdu[0]) {
		/* tagged as uint32_t */
		if apdu[length] == 255 {
			length++
			len_tmp, value32 = decode_unsigned32(apdu[length:])
			length += len_tmp
			value = value32
		} else if (apdu[length] == 254) {
			length++
			len_tmp, value16 = decode_unsigned16(apdu[length:])
			length += len_tmpxo
			value = uint32(value16)
		} else {
			value = uint32(apdu[length])
			length++
		}
	} else if IS_OPENING_TAG(apdu[0]) {
		value = 0
	} else if IS_CLOSING_TAG(apdu[0]) {
		/* closing tag */
		value = 0
	} else {
		/* small value */
		value = uint32(apdu[0] & 0x07)
	}

	return length, tag_number, value
}

/* from clause 20.2.1.3.2 Constructed Data */
/* returns true if the tag is context specific and matches */
func decode_is_context_tag(apdu []byte, tag_number byte) bool {
	_, my_tag_number := decode_tag_number(apdu)
	fmt.Printf("need %d, found %d\n", tag_number, my_tag_number)
	return IS_CONTEXT_SPECIFIC(apdu[0]) && tag_number == my_tag_number
}

/* from clause 20.2.14 Encoding of an Object Identifier Value */
/* returns the number of apdu bytes consumed */
func decode_object_id(apdu []byte) (int, uint16, uint32) {
	var value uint32 = 0
	var len int = 0
	var object_type uint16
	var instance uint32

	len, value = decode_unsigned32(apdu)
	object_type = uint16(((value >> uint32(BACNET_INSTANCE_BITS)) & BACNET_MAX_OBJECT))
	instance = (value & uint32(BACNET_MAX_INSTANCE))

	return len, object_type, instance
}

/* from clause 20.2.14 Encoding of an Object Identifier Value */
/* returns the number of apdu bytes consumed */
func encode_bacnet_object_id(apdu []byte, object_type int, instance uint32) int {
	var value uint32 = 0
	var objType uint32 = 0
	var len int = 0

	objType = uint32(object_type)
	value = ((objType & uint32(BACNET_MAX_OBJECT)) << uint32(BACNET_INSTANCE_BITS)) | (instance & uint32(BACNET_MAX_INSTANCE))
	len = encode_unsigned32(apdu, value)

	return len
}

/* from clause 20.2.14 Encoding of an Object Identifier Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func encode_context_object_id(apdu []byte, tag_number byte, object_type int, instance uint32) int {
	var length int = 0

	/* length of object id is 4 octets, as per 20.2.14 */
	length = encode_tag(apdu, tag_number, true, 4)
	length += encode_bacnet_object_id(apdu[length:], object_type, instance)

	return length
}

/* from clause 20.2.14 Encoding of an Object Identifier Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func encode_application_object_id(apdu []byte, object_type int, instance uint32) int {
	var len int = 0

	/* assumes that the tag only consumes 1 octet */
	len = encode_bacnet_object_id(apdu[1:], object_type, uint32(instance))
	len += encode_tag(apdu, BACNET_APPLICATION_TAG_OBJECT_ID, false, uint32(len))

	return len
}

/* from clause 20.2.9 Encoding of a Character String Value */
/* returns the number of apdu bytes consumed, or zero if failed */
func encode_bacnet_character_string_safe(apdu []byte, max_apdu uint32, encoding byte, pString []byte, length uint32) uint32 {
	var apdu_len uint32 = 1 /*encoding */

	var i uint32

	apdu_len += length
	if apdu_len <= max_apdu {
		apdu[0] = encoding
		for i = 0; i < length; i++ {
			apdu[1 + i] = pString[i]
		}
	} else {
		apdu_len = 0
	}

	return apdu_len
}

func encode_bacnet_character_string(apdu []byte, char_string *BACNET_CHARACTER_STRING) int {
	cse := characterstring_encoding(char_string)
	csv := characterstring_value(char_string)
	csl := characterstring_length(char_string)
	return int(encode_bacnet_character_string_safe(apdu, MAX_APDU, cse, csv, csl))
}

/* from clause 20.2.9 Encoding of a Character String Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func encode_application_character_string(apdu []byte, char_string *BACNET_CHARACTER_STRING) int {
	var length int = 0
	var len_str int = 0

	len_str = int(characterstring_length(char_string)) + 1 /* for encoding */
	length = encode_tag(apdu, BACNET_APPLICATION_TAG_CHARACTER_STRING, false, uint32(len_str))
	if length + len_str < MAX_APDU {
		length += encode_bacnet_character_string(apdu[length:], char_string)
	} else {
		length = 0
	}

	return length
}

func encode_context_character_string(apdu []byte, tag_number byte, char_string *BACNET_CHARACTER_STRING) int {
	var len_tmp int = 0
	var string_len int = 0

	string_len = int(characterstring_length(char_string)) + 1 /* for encoding */
	len_tmp += encode_tag(apdu, tag_number, true, uint32(string_len))
	if len_tmp + string_len < MAX_APDU {
		len_tmp += encode_bacnet_character_string(apdu[len_tmp:], char_string)
	} else {
		len_tmp = 0
	}

	return len_tmp
}

/* from clause 20.2.9 Encoding of a Character String Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func decode_character_string(apdu []byte, len_value uint32) (int, BACNET_CHARACTER_STRING) {
	var len_tmp int = 0
	var status bool = false
	var char_string BACNET_CHARACTER_STRING

	status = characterstring_init(&char_string, apdu[0], apdu[1:], len_value - 1)
	if status {
		len_tmp = int(len_value)
	}

	return len_tmp, char_string
}

/* from clause 20.2.4 Encoding of an Unsigned Integer Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func decode_unsigned(apdu []byte, len_value uint32) (int, uint32) {
	var unsigned16_value uint16 = 0
	var value uint32 = 0
	switch (len_value) {
	case 1:
		value = uint32(apdu[0])
	case 2:
		_, unsigned16_value = decode_unsigned16(apdu)
		value = uint32(unsigned16_value)
	case 3:
		_, value = decode_unsigned24(apdu)
	case 4:
		_, value = decode_unsigned32(apdu)
	default:
		value = 0
	}
	return int(len_value), value
}

/* from clause 20.2.4 Encoding of an Unsigned Integer Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func encode_bacnet_unsigned(apdu []byte, value uint32) int {
	var len int = 0        /* return value */
	if (value < 0x100) {
		apdu[0] = byte(value)
		len = 1
	} else if (value < 0x10000) {
		len = encode_unsigned16(apdu, uint16(value))
	} else if (value < 0x1000000) {
		len = encode_unsigned24(apdu, value)
	} else {
		len = encode_unsigned32(apdu, value)
	}
	return len
}

/* from clause 20.2.4 Encoding of an Unsigned Integer Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func encode_context_unsigned(apdu []byte, tag_number uint8, value uint32) int {
	var length int = 0

	/* length of unsigned is variable, as per 20.2.4 */
	if (value < 0x100) {
		length = 1
	} else if (value < 0x10000) {
		length = 2
	} else if (value < 0x1000000) {
		length = 3
	} else {
		length = 4
	}

	length = encode_tag(apdu, tag_number, true, uint32(length))
	length += encode_bacnet_unsigned(apdu[length:], value)

	return length
}

/* from clause 20.2.4 Encoding of an Unsigned Integer Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func encode_application_unsigned(apdu []byte, value uint32) int {
	var length int = 0

	length = encode_bacnet_unsigned(apdu[1:], value)
	length += encode_tag(apdu, BACNET_APPLICATION_TAG_UNSIGNED_INT, false, uint32(length))

	return length
}

/* from clause 20.2.11 Encoding of an Enumerated Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func decode_enumerated(apdu []byte, len_value uint32) (int, uint32) {
	var unsigned_value uint32 = 0
	var length int

	length, unsigned_value = decode_unsigned(apdu, len_value)

	return length, unsigned_value
}

/* from clause 20.2.11 Encoding of an Enumerated Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func encode_bacnet_enumerated(apdu []byte, value uint32) int {
	return encode_bacnet_unsigned(apdu, value)
}

/* from clause 20.2.11 Encoding of an Enumerated Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func encode_application_enumerated(apdu []byte, value uint32) int {
	var length int = 0        /* return value */

	/* assumes that the tag only consumes 1 octet */
	length = encode_bacnet_enumerated(apdu[1:], value)
	length += encode_tag(apdu, BACNET_APPLICATION_TAG_ENUMERATED, false, uint32(length))

	return length
}

/* from clause 20.2.11 Encoding of an Enumerated Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func encode_context_enumerated(apdu []byte, tag_number byte, value uint32) int {
	var length int = 0

	if value < 0x100 {
		length = 1
	} else if value < 0x10000 {
		length = 2
	} else if value < 0x1000000 {
		length = 3
	} else {
		length = 4
	}

	length = encode_tag(apdu, tag_number, true, uint32(length))
	length += encode_bacnet_enumerated(apdu[length:], value)

	return length
}

/* from clause 20.2.13 Encoding of a Time Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func encode_bacnet_time(apdu []byte, btime *BACNET_TIME) int {
	apdu[0] = btime.Hours
	apdu[1] = btime.Minutes
	apdu[2] = btime.Seconds
	apdu[3] = btime.Hundredths
	return 4
}

/* from clause 20.2.13 Encoding of a Time Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func encode_application_time(apdu []byte, btime *BACNET_TIME) int {
	var length int = 0
	/* assumes that the tag only consumes 1 octet */
	length = encode_bacnet_time(apdu[1:], btime)
	length += encode_tag(apdu, BACNET_APPLICATION_TAG_TIME, false, uint32(length))
	return length
}

func encode_context_time(apdu []byte, tag_number byte, btime *BACNET_TIME) int {
	var length int = 0 /* return value */

	/* length of time is 4 octets, as per 20.2.13 */
	length = encode_tag(apdu, tag_number, true, 4)
	length += encode_bacnet_time(apdu[length:], btime)

	return length
}


/* BACnet Date */
/* year = years since 1900 */
/* month 1=Jan */
/* day = day of month */
/* wday 1=Monday...7=Sunday */

/* from clause 20.2.12 Encoding of a Date Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func encode_bacnet_date(apdu []byte, bdate *BACNET_DATE) int {
	/* allow 2 digit years */
	if (bdate.Year >= 1900) {
		apdu[0] = byte(bdate.Year - 1900)
	} else if (bdate.Year < 0x100) {
		apdu[0] = byte(bdate.Year)
	} else {
		/*
		 ** Don't try and guess what the user meant here. Just fail
		 */
		return BACNET_STATUS_ERROR
	}

	apdu[1] = bdate.Month
	apdu[2] = bdate.Day
	apdu[3] = bdate.Weekday

	return 4
}

/* from clause 20.2.12 Encoding of a Date Value */
/* and 20.2.1 General Rules for Encoding BACnet Tags */
/* returns the number of apdu bytes consumed */
func encode_application_date(apdu []byte, bdate *BACNET_DATE) int {
	var len int = 0

	/* assumes that the tag only consumes 1 octet */
	len = encode_bacnet_date(apdu[1:], bdate)
	len += encode_tag(apdu, BACNET_APPLICATION_TAG_DATE, false, uint32(len))
	return len
}

func encode_context_date(apdu []byte, tag_number byte, bdate *BACNET_DATE) int {
	var len int = 0 /* return value */

	/* length of date is 4 octets, as per 20.2.12 */
	len = encode_tag(apdu, tag_number, true, 4)
	len += encode_bacnet_date(apdu[len:], bdate)
	return len
}
