//QRCoded file test cases
package main

import "testing"

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	result := GenerateQRCode("555-2368")

	if result == nil {
		t.Error("Generated QRCode is nil")
	}

	if len(result) == 0 {
		t.Error("Generated QRCode has no data")
	}
}
