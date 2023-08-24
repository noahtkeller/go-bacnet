package bacnet

import (
	"testing"
)

func TestIAmEncodingAndDecoding(t *testing.T) {
	encodedService := &IAm{Length: 0, PDU: make([]byte, 50)}
	decodedService := &IAm{Length: 0}

	encodedService.Encode()

	decodedService.PDU = encodedService.PDU
	decodedService.Decode()

	if encodedService.Length != decodedService.Length {
		t.Errorf("Length mismatch: %d %d", encodedService.Length, decodedService.Length)
	} else if encodedService.VendorId != decodedService.VendorId {
		t.Errorf("VendorId mismatch: %d %d", encodedService.VendorId, decodedService.VendorId)
	} else if encodedService.MaxAPDU != decodedService.MaxAPDU {
		t.Errorf("MaxAPDU mismatch: %d %d", encodedService.MaxAPDU, decodedService.MaxAPDU)
	} else if encodedService.Segmentation != decodedService.Segmentation {
		t.Errorf("Segmentation mismatch: %d %d", encodedService.Segmentation, decodedService.Segmentation)
	}
}
