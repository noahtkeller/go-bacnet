package bacnet

type BACNET_CHARACTER_STRING struct {
	Encoding byte
	Length   uint32
	Value    []byte
}

/* returns false if the string exceeds capacity
   initialize by using value=NULL */
func characterstring_init(char_string *BACNET_CHARACTER_STRING, encoding byte, value []byte, length uint32) bool {
	var status bool = false
	var i uint32

	if char_string != nil {
		char_string.Length = 0
		char_string.Encoding = encoding
		if length <= CHARACTER_STRING_CAPACITY {
			if value != nil {
				for i = 0; i < MAX_CHARACTER_STRING_BYTES; i++ {
					if i < length {
						char_string.Value[char_string.Length] = value[i]
						char_string.Length++
					} else {
						char_string.Value[i] = 0
					}
				}
			} else {
				for i = 0; i < MAX_CHARACTER_STRING_BYTES; i++ {
					char_string.Value[i] = 0
				}
			}
			status = true
		}
	}

	return status
}

func characterstring_init_ansi(char_string *BACNET_CHARACTER_STRING, value []byte) bool {
	return characterstring_init(char_string, CHARACTER_ANSI_X34, value, uint32(len(value)))
}

/* Returns the value. */
func characterstring_value(char_string *BACNET_CHARACTER_STRING) []byte {
	var value []byte = nil

	if char_string != nil {
		value = char_string.Value
	}

	return value
}

/* returns the length. */
func characterstring_length(char_string *BACNET_CHARACTER_STRING) uint32 {
	var length uint32 = 0

	if char_string != nil {
		/* FIXME: validate length is within bounds? */
		length = char_string.Length
	}

	return length
}

func characterstring_capacity(char_string *BACNET_CHARACTER_STRING) uint32 {
	var length uint32 = 0

	if char_string != nil {
		length = CHARACTER_STRING_CAPACITY
	}

	return length
}

/* returns the encoding. */
func characterstring_encoding(char_string *BACNET_CHARACTER_STRING) byte {
	var encoding byte = 0
	if char_string != nil {
		encoding = char_string.Encoding
	}
	return encoding
}
