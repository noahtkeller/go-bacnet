package bacnet

import (
	"testing"
)

func checkWhoHasLimitEncoding(low_limit, high_limit int32, t *testing.T) {

	encodedService := &WhoHas{Length: 0, PDU: make([]byte, 50), LowLimit: low_limit, HighLimit: high_limit}
	decodedService := &WhoHas{Length: 0}

	encodedService.Encode()

	decodedService.PDU = encodedService.PDU
	decodedService.Decode()

	if encodedService.Length != decodedService.Length {
		//t.Errorf("Length mismatch: %d %d", encodedService.Length, decodedService.Length)
	} else if encodedService.LowLimit != decodedService.LowLimit {
		//t.Errorf("LowLimit mismatch: %d %d", encodedService.LowLimit, decodedService.LowLimit)
	} else if encodedService.HighLimit != decodedService.HighLimit {
		//t.Errorf("HighLimit mismatch: %d %d", encodedService.HighLimit, decodedService.HighLimit)
	}

}

func TestWhoHasEncodingAndDecoding(t *testing.T) {

	var low_limit, high_limit int32

	checkWhoHasLimitEncoding(-1, -1, t)

	for low_limit = 0; low_limit <= BACNET_MAX_INSTANCE; low_limit += (BACNET_MAX_INSTANCE / 4) {
		for high_limit = 0; high_limit <= BACNET_MAX_INSTANCE; high_limit += (BACNET_MAX_INSTANCE / 4) {
			checkWhoHasLimitEncoding(low_limit, high_limit, t)
		}
	}

}
