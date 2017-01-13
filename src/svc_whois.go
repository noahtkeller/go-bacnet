package bacnet

func NewWhoIsAPDU(pdu []byte, high, low int32) *WhoIs {
	if pdu == nil {
		pdu = make([]byte, 50)
	}
	apdu := &WhoIs{
		PDU: pdu,
		LowLimit: low,
		HighLimit: high,
	}
	return apdu
}

type WhoIs struct {
	LowLimit  int32
	HighLimit int32
	PDU       []byte
	Length    int
}

func (self *WhoIs) SetLimits(low_limit, high_limit int32) *WhoIs {
	self.LowLimit = low_limit
	self.HighLimit = high_limit
	return self
}

func (self *WhoIs) Encode() Service {
	var encodeIdx int = 2
	self.PDU[0] = PDU_TYPE_UNCONFIRMED_SERVICE_REQUEST
	self.PDU[1] = SERVICE_UNCONFIRMED_WHO_IS

	if self.LowLimit >= 0 && self.LowLimit <= BACNET_MAX_INSTANCE && self.HighLimit >= 0 && self.HighLimit <= BACNET_MAX_INSTANCE {
		encodeIdx += encode_context_unsigned(self.PDU[encodeIdx:], 0, uint32(self.LowLimit))
		encodeIdx += encode_context_unsigned(self.PDU[encodeIdx:], 1, uint32(self.HighLimit))
	}

	self.PDU = self.PDU[:encodeIdx]
	self.Length = encodeIdx
	return self
}

func (self *WhoIs) Decode() Service {

	var decodeIdx int = 0
	var len_tmp int = 0
	var tag_number byte = 0
	var len_value uint32 = 0
	var decoded_value uint32 = 0
	var apdu_len int = 0
	var offset int

	var slice []byte

	if self.PDU[0] == PDU_TYPE_UNCONFIRMED_SERVICE_REQUEST && self.PDU[1] == SERVICE_UNCONFIRMED_WHO_IS {
		slice = self.PDU[2:]
		offset = 2
	} else {
		slice = self.PDU
		offset = 0
	}

	apdu_len = len(slice)

	if apdu_len > 0 {
		len_tmp, tag_number, len_value = decode_tag_number_and_value(slice[decodeIdx:])
		decodeIdx += len_tmp
		if tag_number == 0 && apdu_len > decodeIdx {
			len_tmp, decoded_value = decode_unsigned(slice[decodeIdx:], len_value)
			decodeIdx += len_tmp
			if decoded_value <= uint32(BACNET_MAX_INSTANCE) {
				self.LowLimit = int32(decoded_value)
				if apdu_len > decodeIdx {
					len_tmp, tag_number, len_value = decode_tag_number_and_value(slice[decodeIdx:])
					decodeIdx += len_tmp
					if tag_number == 1 && apdu_len > decodeIdx {
						len_tmp, decoded_value = decode_unsigned(slice[decodeIdx:], len_value)
						decodeIdx += len_tmp
						if decoded_value <= uint32(BACNET_MAX_INSTANCE) {
							self.HighLimit = int32(decoded_value)
							self.Length = decodeIdx + offset
							return self
						}
					}
				}
			}
		}
		// if any of the above conditions are false return status error
		panic("There was a problem decoding the WhoIs Service Request")
	}

	self.LowLimit = -1
	self.HighLimit = -1
	self.Length = offset
	return self
}