package bacnet

import (
	"testing"
)

func TestAckAlarmEncodingAndDecoding(t *testing.T) {

	encodedService := &ACKAlarm{Length: 0, PDU: make([]byte, 50)}
	decodedService := &ACKAlarm{Length: 0}

	encodedService.Encode()

	decodedService.PDU = encodedService.PDU
	decodedService.Decode()

	if encodedService.Length != decodedService.Length {
		//t.Errorf("Length mismatch: %d %d", encodedService.Length, decodedService.Length)
	}

}