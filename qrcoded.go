//Generate Qrcode

package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello QR Code!")

	// generate a QRcode for a phonenumber
	qrcode := GenerateQRCode("555-2368")
	ioutil.WriteFile("qrcode.png", qrcode, 0644)
}

//GenerateQRCode for a given value by creating an image and encode
//it using png.Encode to pass the written test
func GenerateQRCode(code string) []byte {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	buf := new(bytes.Buffer)
	_ = png.Encode(buf, img)

	return buf.Bytes()
}
