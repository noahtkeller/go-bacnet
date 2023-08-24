package bacnet

func NewIAmAPDU(pdu []byte, deviceId, maxAPDU uint32, vendorId uint16) *IAm {
	if pdu == nil {
		pdu = make([]byte, 100)
	}
	apdu := &IAm{
		PDU: pdu,
		DeviceId: deviceId,
		VendorId: vendorId,
	}
	return apdu
}

type IAm struct {
	PDU          []byte
	DeviceId     uint32
	MaxAPDU      uint32
	Segmentation int
	VendorId     uint16
	Length       int
}

func (self *IAm) Encode() *IAm {
	var encodeIdx int = 2
	self.PDU[0] = PDU_TYPE_UNCONFIRMED_SERVICE_REQUEST
	self.PDU[1] = SERVICE_UNCONFIRMED_I_AM

	encodeIdx += encode_application_object_id(self.PDU[encodeIdx:], int(OBJECT_DEVICE), self.DeviceId)
	encodeIdx += encode_application_unsigned(self.PDU[encodeIdx:], self.MaxAPDU)
	encodeIdx += encode_application_enumerated(self.PDU[encodeIdx:], uint32(self.Segmentation))
	encodeIdx += encode_application_unsigned(self.PDU[encodeIdx:], uint32(self.VendorId))

	self.PDU = self.PDU[:encodeIdx]
	self.Length = encodeIdx
	return self
}

func (self *IAm) Decode() *IAm {

	var len_tmp int = 0
	var decodeIdx int = 0   /* total length of the apdu, return value */
	var object_type uint16 = 0   /* should be a Device Object */
	var object_instance uint32 = 0
	var tag_number byte = 0
	var len_value uint32 = 0
	var decoded_value uint32 = 0
	var offset int

	var slice []byte

	if self.PDU[0] == PDU_TYPE_UNCONFIRMED_SERVICE_REQUEST && self.PDU[1] == SERVICE_UNCONFIRMED_I_AM {
		slice = self.PDU[2:]
		offset = 2
	} else {
		slice = self.PDU
		offset = 0
	}

	/* OBJECT ID - object id */
	len_tmp, tag_number, len_value = decode_tag_number_and_value(slice[decodeIdx:])
	decodeIdx += len_tmp
	if (tag_number == BACNET_APPLICATION_TAG_OBJECT_ID) {
		len_tmp, object_type, object_instance = decode_object_id(slice[decodeIdx:])
		decodeIdx += len_tmp
		if (object_type == uint16(OBJECT_DEVICE)) {
			self.DeviceId = object_instance
			/* MAX APDU - unsigned */
			len_tmp, tag_number, len_value = decode_tag_number_and_value(slice[decodeIdx:])
			decodeIdx += len_tmp
			if (tag_number == BACNET_APPLICATION_TAG_UNSIGNED_INT) {
				len_tmp, decoded_value = decode_unsigned(slice[decodeIdx:], len_value)
				decodeIdx += len_tmp
				self.MaxAPDU = uint32(decoded_value)
				/* Segmentation - enumerated */
				len_tmp, tag_number, len_value = decode_tag_number_and_value(slice[decodeIdx:])
				decodeIdx += len_tmp
				if (tag_number == BACNET_APPLICATION_TAG_ENUMERATED) {
					len_tmp, decoded_value = decode_enumerated(slice[decodeIdx:], len_value)
					decodeIdx += len_tmp
					if (decoded_value < uint32(MAX_BACNET_SEGMENTATION)) {
						self.Segmentation = int(decoded_value)
						/* Vendor ID - unsigned16 */
						len_tmp, tag_number, len_value = decode_tag_number_and_value(slice[decodeIdx:])
						decodeIdx += len_tmp
						if (tag_number == BACNET_APPLICATION_TAG_UNSIGNED_INT) {
							len_tmp, decoded_value = decode_unsigned(slice[decodeIdx:], len_value)
							decodeIdx += len_tmp
							if (decoded_value <= 0xFFFF) {
								self.VendorId = uint16(decoded_value)
								self.Length = decodeIdx + offset
								return self
							}
						}
					}
				}
			}
		}
		panic("There was a problem decoding the IAm Service Request")
	}

	self.Length = offset
	return self
}