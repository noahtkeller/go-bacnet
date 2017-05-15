package bacnet

func NewWhoHasAPDU(pdu []byte, high, low int32, isObjectName bool, identifier BACNET_OBJECT_ID, name BACNET_CHARACTER_STRING) *WhoHas {
	if pdu == nil {
		pdu = make([]byte, 50)
	}
	apdu := &WhoHas{
		PDU: pdu,
		LowLimit: low,
		HighLimit: high,
		IsObjectName: isObjectName,
		Identifier: identifier,
		Name: name,
	}
	return apdu
}

type WhoHas struct {
	LowLimit     int32
	HighLimit    int32
	IsObjectName bool
	Identifier   BACNET_OBJECT_ID
	Name         BACNET_CHARACTER_STRING
	PDU          []byte
	Length       int
}

func (self *WhoHas) SetLimits(low_limit, high_limit int32) *WhoHas {
	self.LowLimit = low_limit
	self.HighLimit = high_limit
	return self
}

func (self *WhoHas) SetIdentifier(identifier BACNET_OBJECT_ID) *WhoHas {
	self.Identifier = identifier
	return self
}

func (self *WhoHas) SetName(name BACNET_CHARACTER_STRING) *WhoHas {
	self.Name = name
	return self
}

func (self *WhoHas) Encode() Service {
	var encodeIdx int = 2
	self.PDU[0] = PDU_TYPE_UNCONFIRMED_SERVICE_REQUEST
	self.PDU[1] = SERVICE_UNCONFIRMED_WHO_HAS

	if self.LowLimit >= 0 && self.LowLimit <= BACNET_MAX_INSTANCE && self.HighLimit >= 0 && self.HighLimit <= BACNET_MAX_INSTANCE {
		encodeIdx += encode_context_unsigned(self.PDU[encodeIdx:], 0, uint32(self.LowLimit))
		encodeIdx += encode_context_unsigned(self.PDU[encodeIdx:], 1, uint32(self.HighLimit))
	}
	if self.IsObjectName {
		encodeIdx += encode_context_character_string(self.PDU[encodeIdx:], 3, &self.Name)
	} else {
		encodeIdx += encode_context_object_id(self.PDU[encodeIdx:], 2, int(self.Identifier.Type), self.Identifier.Instance)
	}

	self.PDU = self.PDU[:encodeIdx]
	self.Length = encodeIdx
	return self
}

// TODO: WhoHas Service Request Decode
func (self *WhoHas) Decode() Service {

	var decodeIdx int = 0
	var len_tmp int = 0
	var len_value uint32 = 0
	var decoded_value uint32 = 0

	if self.PDU[0] == PDU_TYPE_UNCONFIRMED_SERVICE_REQUEST && self.PDU[1] == SERVICE_UNCONFIRMED_WHO_HAS {
		decodeIdx = 2
	}

	var apdu_len int = len(self.PDU[decodeIdx:])

	if apdu_len > 0 {
		// optional limits - must be used as a pair
		if (decode_is_context_tag(self.PDU[decodeIdx:], 0)) {

			len_tmp, _, len_value = decode_tag_number_and_value(self.PDU[decodeIdx:])
			decodeIdx += len_tmp

			len_tmp, decoded_value = decode_unsigned(self.PDU[decodeIdx:], len_value)
			decodeIdx += len_tmp

			if (int32(decoded_value) <= BACNET_MAX_INSTANCE) {
				self.LowLimit = int32(decoded_value)
			}

			// Decode HighLimit if available
			if (!decode_is_context_tag(self.PDU[decodeIdx:], 1)) {
				self.Length = -1
				return self
			}

			len_tmp, _, len_value = decode_tag_number_and_value(self.PDU[decodeIdx:])
			decodeIdx += len_tmp

			len_tmp, decoded_value = decode_unsigned(self.PDU[decodeIdx:], len_value)
			if (int32(decoded_value) <= BACNET_MAX_INSTANCE) {
				self.HighLimit = int32(decoded_value)
			}

		} else {

			self.LowLimit = -1
			self.HighLimit = -1

		}

		/* object id */
		if (decode_is_context_tag(self.PDU[decodeIdx:], 2)) {

			self.IsObjectName = false
			len_tmp, _, len_value = decode_tag_number_and_value(self.PDU[decodeIdx:])
			decodeIdx += len_tmp

			len_tmp, self.Identifier.Type, self.Identifier.Instance = decode_object_id(self.PDU[decodeIdx:])
			decodeIdx += len_tmp

		} else if (decode_is_context_tag(self.PDU[decodeIdx:], 3)) {

			self.IsObjectName = true
			len_tmp, _, len_value = decode_tag_number_and_value(self.PDU[decodeIdx:])
			decodeIdx += len_tmp

			len_tmp, self.Name = decode_character_string(self.PDU[decodeIdx:], len_value)
			decodeIdx += len_tmp

		} else {

			//fmt.Println("test")
			//self.Length = -1
			//return self

		}
	}

	self.Length = decodeIdx
	return self

}