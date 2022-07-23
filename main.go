package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/pschlump/filelib"
)

var out *os.File = os.Stdout

var raw = flag.Bool("raw", false, "just print out the text with no extra stuff.")
var output = flag.String("output", "", "file to encode")
var help = flag.Bool("help", false, "print out the help message.")

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "qr-decode: Usage: %s [flags]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse() // Parse CLI arguments to this, --cfg <name>.json

	fns := flag.Args()

	if *help {
		flag.Usage()
		os.Exit(1)
	}

	if len(fns) == 0 {
		fmt.Printf("Missing arguments (list of images to decode)\n")
		flag.Usage()
		os.Exit(1)
	}

	if *output != "" {
		var err error
		out, err = filelib.Fopen(*output, "w")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to open %s for output, error: %s\n", *output, err)
			os.Exit(1)
		}
	}

	// for ii, fn := range os.Args {
	for _, fn := range fns {
		//if ii == 0 {
		//	continue
		//}

		file, err := filelib.Fopen(fn, "r")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid file: %s error:%s\n", fn, err)
			continue
		}
		img, _, err := image.Decode(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid QR code, file: %s error:%s\n", fn, err)
			continue
		}

		// prepare BinaryBitmap
		bmp, err := gozxing.NewBinaryBitmapFromImage(img)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to convert to QR-code to bitmap, file: %s error:%s\n", fn, err)
			continue
		}

		// decode image
		qrReader := qrcode.NewQRCodeReader()
		result, err := qrReader.Decode(bmp, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to decode QR, file : %s error:%s\n", fn, err)
			continue
		}

		if *raw {
			fmt.Fprintf(out, "%s\n", result)
		} else {
			fmt.Fprintf(out, "%s: %s\n", fn, result)
		}
	}
}
