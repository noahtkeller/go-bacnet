package bacnet

func NewNPDU(pdu []byte, dest, src *BACNET_ADDRESS, expectReply bool, priority byte) *NPDU {
	if pdu == nil {
		pdu = make([]byte, 50)
	}
	npdu := &NPDU{
		PDU: pdu,
		Dest: dest,
		Source: src,
		ExpectingReply: expectReply,
		Priority: priority,
		ProtocolVersion: 0x01,
		NetworkLayer: false,
		HopCount: HOP_COUNT_DEFAULT,
	}
	return npdu
}

func NewNetworkLayerNPDU(pdu []byte, dest, src *BACNET_ADDRESS, expectReply bool, priority byte) *NPDU {
	npdu := NewNPDU(pdu, dest, src, expectReply, priority)
	npdu.NetworkLayer = true
	return npdu
}

type NPDU struct {
	PDU             []byte
	ProtocolVersion byte
	ExpectingReply  bool
	NetworkLayer    bool
	Priority        byte
	MessageType     byte
	VendorId        uint16
	HopCount        byte
	Dest            *BACNET_ADDRESS
	Source          *BACNET_ADDRESS
	Length          int
}

func (self *NPDU) Encode() *NPDU {
	var encodeIdx int = 0
	var i byte = 0

	encodeIdx = 2
	self.PDU[0] = self.ProtocolVersion
	self.PDU[1] = 0
	if self.NetworkLayer {
		self.PDU[1] |= BIT7
	}
	if self.Dest != nil && self.Dest.Net != 0 {
		self.PDU[1] |= BIT5
	}
	if self.Source != nil && self.Source.Net != 0 && self.Source.Len != 0 {
		self.PDU[1] |= BIT3
	}
	if self.ExpectingReply {
		self.PDU[1] |= BIT2
	}
	self.PDU[1] |= (self.Priority & 0x03)
	if self.Dest != nil && self.Dest.Net != 0 {
		encodeIdx += encode_unsigned16(self.PDU[encodeIdx:], self.Dest.Net)
		self.PDU[encodeIdx] = self.Dest.Len
		encodeIdx++
		if self.Dest.Len != 0 {
			for i = 0; i < self.Dest.Len; i, encodeIdx = i + 1, encodeIdx + 1 {
				self.PDU[encodeIdx] = self.Dest.Adr[i]
			}
		}
	}
	if self.Source != nil && self.Source.Net != 0 && self.Source.Len != 0 {
		encodeIdx += encode_unsigned16(self.PDU[encodeIdx:], self.Source.Net)
		self.PDU[encodeIdx] = self.Source.Len
		encodeIdx++
		if self.Source.Len != 0 {
			for i = 0; i < self.Source.Len; i, encodeIdx = i + 1, encodeIdx + 1 {
				self.PDU[encodeIdx] = self.Source.Adr[i]
			}
		}
	}
	if self.Dest != nil && self.Dest.Net != 0 {
		self.PDU[encodeIdx] = self.HopCount
		encodeIdx++
	}
	if self.NetworkLayer {
		self.PDU[encodeIdx] = self.MessageType
		encodeIdx++
		if self.MessageType >= 0x80 {
			encodeIdx += encode_unsigned16(self.PDU[encodeIdx:], self.VendorId)
		}
	}

	self.Length = encodeIdx

	return self

}

func (self *NPDU) Decode() *NPDU {
	var len_pdu int = 0
	var len_tmp = 0
	var i byte = 0
	var src_net uint16 = 0
	var dest_net uint16 = 0
	var address_len byte = 0
	var mac_octet byte = 0

	if self.Dest == nil {
		self.Dest = NewBACNET_ADDRESS()
	}
	if self.Source == nil {
		self.Source = NewBACNET_ADDRESS()
	}

	self.ProtocolVersion = self.PDU[0]
	if self.PDU[1] & BIT7 != 0 {
		self.NetworkLayer = true
	} else {
		self.NetworkLayer = false
	}
	if self.PDU[1] & BIT2 != 0 {
		self.ExpectingReply = true
	} else {
		self.ExpectingReply = false
	}
	self.Priority = self.PDU[1] & 0x03
	len_pdu = 2
	if self.PDU[1] & BIT5 != 0 {
		len_tmp, dest_net = decode_unsigned16(self.PDU[len_pdu:])
		len_pdu += len_tmp
		address_len = self.PDU[len_pdu]
		len_pdu++
		if self.Dest != nil {
			self.Dest.Net = dest_net
			self.Dest.Len = address_len
		}
		if address_len != 0 {
			if address_len > MAX_MAC_LEN {
				panic("Dest address_len greater than MAX_MAC_LEN " + string(MAX_MAC_LEN))
			}
			for i = 0; i < address_len; len_pdu, i = len_pdu + 1, i + 1 {
				mac_octet = self.PDU[len_pdu]
				if self.Dest != nil {
					self.Dest.Adr[i] = mac_octet
				}
			}
		}
	} else if self.Dest != nil {
		self.Dest.Net = 0
		self.Dest.Len = 0
		for i = 0; i < MAX_MAC_LEN; i++ {
			self.Dest.Adr[i] = 0
		}
	}
	if self.PDU[1] & BIT3 != 0 {
		len_tmp, src_net = decode_unsigned16(self.PDU[len_pdu:])
		len_pdu += len_tmp
		address_len = self.PDU[len_pdu]
		len_pdu++
		if self.Source != nil {
			self.Source.Net = src_net
			self.Source.Len = address_len
		}
		if address_len != 0 {
			if address_len > MAX_MAC_LEN {
				panic("Source address_len greater than MAX_MAC_LEN " + string(MAX_MAC_LEN))
			}
			for i = 0; i < address_len; len_pdu, i = len_pdu + 1, i + 1 {
				mac_octet = self.PDU[len_pdu]
				if self.Source != nil {
					self.Source.Adr[i] = mac_octet
				}
			}
		}
	} else if self.Source != nil {
		if self.Source.Net != BACNET_BROADCAST_NETWORK {
			self.Source.Net = 0
		}
		self.Source.Len = 0
		for i = 0; i < MAX_MAC_LEN; i++ {
			self.Source.Adr[i] = 0
		}
	}
	if dest_net != 0 {
		self.HopCount = self.PDU[len_pdu]
		len_pdu++
	} else {
		self.HopCount = 0
	}
	if self.NetworkLayer {
		self.MessageType = self.PDU[len_pdu]
		len_pdu++
		if self.MessageType >= 0x80 {
			len_tmp, self.VendorId = decode_unsigned16(self.PDU[len_pdu:])
			len_pdu += len_tmp
		}
	} else {
		/* Since self.Network_layer_message is false,
		 * it doesn't much matter what we set here this is safe: */
		//self.Network_message_type = NETWORK_MESSAGE_INVALID
	}

	self.Length = len_pdu

	return self
}
