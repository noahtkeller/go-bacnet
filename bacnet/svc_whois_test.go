package bacnet

import (
	"testing"
)

func checkWhoIsLimitEncoding(low_limit, high_limit int32, t *testing.T) {

	encodedService := &WhoIs{Length: 0, PDU: make([]byte, 50), LowLimit: low_limit, HighLimit: high_limit}
	decodedService := &WhoIs{Length: 0}

	encodedService.Encode()

	decodedService.PDU = encodedService.PDU
	decodedService.Decode()

	if encodedService.Length != decodedService.Length {
		t.Errorf("Length mismatch: %d %d", encodedService.Length, decodedService.Length)
	} else if encodedService.LowLimit != decodedService.LowLimit {
		t.Errorf("LowLimit mismatch: %d %d", encodedService.LowLimit, decodedService.LowLimit)
	} else if encodedService.HighLimit != decodedService.HighLimit {
		t.Errorf("HighLimit mismatch: %d %d", encodedService.HighLimit, decodedService.HighLimit)
	}

}

func TestWhoIsEncodingAndDecoding(t *testing.T) {

	var low_limit, high_limit int32

	checkWhoIsLimitEncoding(-1, -1, t)

	for low_limit = 0; low_limit <= BACNET_MAX_INSTANCE; low_limit += (BACNET_MAX_INSTANCE / 4) {
		for high_limit = 0; high_limit <= BACNET_MAX_INSTANCE; high_limit += (BACNET_MAX_INSTANCE / 4) {
			checkWhoIsLimitEncoding(low_limit, high_limit, t)
		}
	}

}
