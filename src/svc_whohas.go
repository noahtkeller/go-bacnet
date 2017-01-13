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
	return self
}