package bacnet

import (
	"testing"
)

func TestNPDUEncodingAndDecoding(t *testing.T) {
	encodedService := &NPDU{Length: 0, PDU: make([]byte, 50)}
	decodedService := &NPDU{Length: 0}

	encodedService.Encode()

	decodedService.PDU = encodedService.PDU
	decodedService.Decode()

	if encodedService.Length != decodedService.Length {
		t.Errorf("Length mismatch: %d %d", encodedService.Length, decodedService.Length)
	} else if encodedService.Priority != decodedService.Priority {
		t.Errorf("Priority mismatch: %d %d", encodedService.Priority, decodedService.Priority)
	} else if encodedService.ProtocolVersion != decodedService.ProtocolVersion {
		t.Errorf("ProtocolVersion mismatch: %d %d", encodedService.ProtocolVersion, decodedService.ProtocolVersion)
	} else if encodedService.NetworkLayer != decodedService.NetworkLayer {
		t.Errorf("NetworkLayer mismatch: %v %v", encodedService.NetworkLayer, decodedService.NetworkLayer)
	} else if encodedService.HopCount != decodedService.HopCount {
		t.Errorf("HopCount mismatch: %d %d", encodedService.HopCount, decodedService.HopCount)
	}
}
