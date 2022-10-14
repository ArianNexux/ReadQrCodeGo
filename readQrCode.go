package main

import (
	"bytes"
	"fmt"
	"github.com/liyue201/goqr"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	//"net/http"
	"os"
)

func recognizeFile(path string) {
	imgdata, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%v\n", err)
		return;
	}
	img, _, err := image.Decode(bytes.NewReader(imgdata))
	if err != nil {
		fmt.Printf("image.Decode error: %v\n", err)
		return;
	}
	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		fmt.Printf("Recognize failed", err)
		return;
	}
	for _, qrCode := range qrCodes {
		fmt.Printf("%s", qrCode.Payload)
	}
	return;
}


func main() {
	recognizeFile(os.Args[1])
}