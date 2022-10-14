package main

import (
	"fmt"
	"image"
    "os"
	// import gif, jpeg, png
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"strings"
    "bytes"
	// import bmp, tiff, webp
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
    	"io/ioutil"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/multi/qrcode"
)

func main(){
    imgdata, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		return;
	}
img, _, err := image.Decode(bytes.NewReader(imgdata))
	if err != nil {
		msg := fmt.Sprintf("failed to read image: %v", err)
		fmt.Printf("", msg)
	}

	source := gozxing.NewLuminanceSourceFromImage(img)
	bin := gozxing.NewHybridBinarizer(source)
	bbm, err := gozxing.NewBinaryBitmap(bin)

	if err != nil {
		msg := fmt.Sprintf("error during processing: %v", err)
		fmt.Printf("", msg)
	}

	qrReader := qrcode.NewQRCodeMultiReader()
	result, err := qrReader.DecodeMultiple(bbm, nil)
	if err != nil {
		msg := fmt.Sprintf("unable to decode QRCode: %v", err)
		fmt.Printf("", msg)
	}
	strRes := []string{}
	for _, element := range result {    
		strRes = append(strRes, element.String())
	}

	res := strings.Join(strRes, "\n")
    fmt.Println(res)
}