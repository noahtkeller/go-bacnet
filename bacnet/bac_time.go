package bacnet

type BACNET_TIME struct {
	Hours      byte
	Minutes    byte
	Seconds    byte
	Hundredths byte
}

type BACNET_DATE struct {
	Year    uint16
	Month   byte
	Day     byte
	Weekday byte
}

type BACNET_DATE_TIME struct {
	Date BACNET_DATE
	Time BACNET_TIME
}

type BACNET_TIMESTAMP struct {
	Tag         byte
	Time        BACNET_TIME
	DateTime    BACNET_DATE_TIME
	SequenceNum uint16
}

type BACNET_TIME_VALUE struct {
	Time  BACNET_TIME
	Value BACNET_APPLICATION_DATA_VALUE
}

const TIME_STAMP_TIME byte = 0x00
const TIME_STAMP_SEQUENCE byte = 0x01
const TIME_STAMP_DATETIME byte = 0x02

const BACNET_WEEKDAY_MONDAY byte = 0x01
const BACNET_WEEKDAY_TUESDAY byte = 0x02
const BACNET_WEEKDAY_WEDNESDAY byte = 0x03
const BACNET_WEEKDAY_THURSDAY byte = 0x04
const BACNET_WEEKDAY_FRIDAY byte = 0x05
const BACNET_WEEKDAY_SATURDAY byte = 0x06
const BACNET_WEEKDAY_SUNDAY byte = 0x07

func bacapp_encode_timestamp(apdu []byte, value *BACNET_TIMESTAMP) int {
	var length int = 0

	if value != nil {
		if value.Tag == TIME_STAMP_TIME {
			length = encode_context_time(apdu, 0, &value.Time)
		} else if value.Tag == TIME_STAMP_SEQUENCE {
			length = encode_context_unsigned(apdu, 1, uint32(value.SequenceNum))
		} else if value.Tag == TIME_STAMP_DATETIME {
			length = bacapp_encode_context_datetime(apdu, 2, &value.DateTime)
		}
	}

	return length
}

func bacapp_encode_context_timestamp(apdu []byte, tag_number byte, value *BACNET_TIMESTAMP) int {
	var len int = 0
	var apdu_len int = 0

	if value != nil {
		len = encode_opening_tag(apdu[apdu_len:], tag_number)
		apdu_len += len
		len = bacapp_encode_timestamp(apdu[apdu_len:], value)
		apdu_len += len
		len = encode_closing_tag(apdu[apdu_len:], tag_number)
		apdu_len += len
	}
	return apdu_len
}

func bacapp_encode_datetime(apdu []byte, value *BACNET_DATE_TIME) int {
	var len int = 0
	var apdu_len int = 0

	if value != nil {
		len = encode_application_date(apdu, &value.Date)
		apdu_len += len
		len = encode_application_time(apdu[apdu_len:], &value.Time)
		apdu_len += len
	}
	return apdu_len
}

func bacapp_encode_context_datetime(apdu []byte, tag_number byte, value *BACNET_DATE_TIME) int {
	var len int = 0
	var apdu_len int = 0

	if value != nil {
		len = encode_opening_tag(apdu[apdu_len:], tag_number)
		apdu_len += len
		len = bacapp_encode_datetime(apdu[apdu_len:], value)
		apdu_len += len
		len = encode_closing_tag(apdu[apdu_len:], tag_number)
		apdu_len += len
	}
	return apdu_len
}
