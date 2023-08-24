package bacnet

func encode_unsigned16(apdu []byte, value uint16) int {
	apdu[0] = byte((value & 0xff00) >> 8)
	apdu[1] = byte(value & 0x00ff)
	return 2
}

func decode_unsigned16(apdu []byte) (int, uint16) {
	var value uint16 = 0
	value = (uint16(apdu[0]) << 8) & 0xff00
	value |= uint16(apdu[1]) & 0x00ff
	return 2, value
}

func encode_unsigned24(apdu []byte, value uint32) int {
	apdu[0] = byte((value & 0xff0000) >> 16)
	apdu[1] = byte((value & 0x00ff00) >> 8)
	apdu[2] = byte(value & 0x0000ff)
	return 3
}

func decode_unsigned24(apdu []byte) (int, uint32) {
	var value uint32 = 0
	value = (uint32(apdu[0]) << 16) & 0x00ff0000
	value |= (uint32(apdu[1]) << 8) & 0x0000ff00
	value |= uint32(apdu[2]) & 0x000000ff
	return 3, value
}

func encode_unsigned32(apdu []byte, value uint32) int {
	apdu[0] = byte((value & 0xff000000) >> 24)
	apdu[1] = byte((value & 0x00ff0000) >> 16)
	apdu[2] = byte((value & 0x0000ff00) >> 8)
	apdu[3] = byte(value & 0x000000ff)
	return 4
}

func decode_unsigned32(apdu []byte) (int, uint32) {
	var value uint32 = 0
	value = (uint32(apdu[0]) << 24) & 0xff000000
	value |= (uint32(apdu[1]) << 16) & 0x00ff0000
	value |= (uint32(apdu[2]) << 8) & 0x0000ff00
	value |= uint32(apdu[3]) & 0x000000ff
	return 4, value
}

func encode_signed8(apdu []byte, value int8) int {
	apdu[0] = byte(value)
	return 1
}

// TODO: I had to change value to int64 to compile
func decode_signed8(apdu []byte, value *int64) int {
	if *value != 0 {
		if apdu[0] & 0x80 != 0 {
			*value = 0xFFFFFF00
		} else {
			*value = 0
		}
		*value |= int64(apdu[0]) & 0x000000ff
	}
	return 1
}

func encode_signed16(apdu []byte, value int32) int {
	apdu[0] = byte((value & 0xff00) >> 8)
	apdu[1] = byte(value & 0x00ff)
	return 2
}

// TODO: I had to change value to int64 to compile
func decode_signed16(apdu []byte, value *int64) int {
	if *value != 0 {
		if apdu[0] & 0x80 != 0 {
			*value = 0xFFFF0000
		} else {
			*value = 0
		}
		*value |= (int64(apdu[0]) << 8) & 0x0000ff00
		*value |= int64(apdu[1]) & 0x000000ff
	}
	return 2
}

func encode_signed24(apdu []byte, value int32) int {
	apdu[0] = byte((value & 0xff0000) >> 16)
	apdu[1] = byte((value & 0x00ff00) >> 8)
	apdu[2] = byte(value & 0x0000ff)
	return 3
}

// TODO: I had to change value to int64 to compile
func decode_signed24(apdu []byte, value *int64) int {
	if *value != 0 {
		if apdu[0] & 0x80 != 0 {
			*value = 0xFF000000
		} else {
			*value = 0
		}
		*value |= (int64(apdu[0]) << 16) & 0x00ff0000
		*value |= (int64(apdu[1]) << 8) & 0x0000ff00
		*value |= int64(apdu[2]) & 0x000000ff
	}
	return 3
}

func encode_signed32(apdu []byte, value int64) int {
	apdu[0] = byte((value & 0xff000000) >> 24)
	apdu[1] = byte((value & 0x00ff0000) >> 16)
	apdu[2] = byte((value & 0x0000ff00) >> 8)
	apdu[3] = byte(value & 0x000000ff)
	return 4
}

// TODO: I had to change value to int64 to compile
func decode_signed32(apdu []byte, value *int64) int {
	if *value != 0 {
		*value = (int64(apdu[0]) << 24) & 0xff000000
		*value |= (int64(apdu[1]) << 16) & 0x00ff0000
		*value |= (int64(apdu[2]) << 8) & 0x0000ff00
		*value |= int64(apdu[3]) & 0x000000ff
	}
	return 4
}
