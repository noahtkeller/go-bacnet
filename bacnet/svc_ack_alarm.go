package bacnet

func NewACKAlarmAPDU(pdu []byte) *ACKAlarm {
	ack := &ACKAlarm{
		PDU: pdu,
		AckSource: BACNET_CHARACTER_STRING{Length: 0, Value: make([]byte, MAX_CHARACTER_STRING_BYTES)},
	}
	return ack
}

type ACKAlarm struct {
	PDU                   []byte
	MaxAPDU               uint32
	InvokeId              byte
	Length                int
	AckProcessIdentifier  uint32
	EventObjectIdentifier BACNET_OBJECT_ID
	EventStateAcked       byte
	EventTimeStamp        BACNET_TIMESTAMP
	AckSource             BACNET_CHARACTER_STRING
	AckTimeStamp          BACNET_TIMESTAMP
}

func (self *ACKAlarm) Encode() *ACKAlarm {
	var encodeIdx int = 4

	self.PDU[0] = PDU_TYPE_CONFIRMED_SERVICE_REQUEST
	self.PDU[1] = encode_max_segs_max_apdu(0, self.MaxAPDU)
	self.PDU[2] = self.InvokeId
	self.PDU[3] = SERVICE_CONFIRMED_ACKNOWLEDGE_ALARM

	encodeIdx += encode_context_unsigned(self.PDU[encodeIdx:], 0, self.AckProcessIdentifier)
	encodeIdx += encode_context_object_id(self.PDU[encodeIdx:], 1, int(self.EventObjectIdentifier.Type), self.EventObjectIdentifier.Instance)
	encodeIdx += encode_context_enumerated(self.PDU[encodeIdx:], 2, uint32(self.EventStateAcked))
	encodeIdx += bacapp_encode_context_timestamp(self.PDU[encodeIdx:], 3, &self.EventTimeStamp)
	encodeIdx += encode_context_character_string(self.PDU[encodeIdx:], 4, &self.AckSource)
	encodeIdx += bacapp_encode_context_timestamp(self.PDU[encodeIdx:], 5, &self.AckTimeStamp)

	self.PDU = self.PDU[:encodeIdx]
	self.Length = encodeIdx

	return self
}

func (self *ACKAlarm) Decode() *ACKAlarm {
	return self
}
