package main

import (
	"fmt"
	"log"
	"os"

	"github.com/otiai10/gosseract"
	"gopkg.in/gographics/imagick.v2/imagick"
)

func main() {
	pdfName := "data/Aegon-V32009832.pdf"
	imageName := "dataout/test.jpg"
	var f, err = os.Create("test.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	if err := ConvertPdfToJpg(pdfName, imageName); err != nil {
		log.Fatal(err)
	}

	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("dataout/test.jpg")
	text, _ := client.Text()
	f.WriteString(text)
	f.Close()
	fmt.Println(text)

}

func ConvertPdfToJpg(pdfName string, imageName string) error {

	// Setup
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// Must be *before* ReadImageFile
	// Make sure our image is high quality
	if err := mw.SetResolution(400, 400); err != nil {
		return err
	}

	// Load the image file into imagick
	if err := mw.ReadImage(pdfName); err != nil {
		return err
	}

	// Must be *after* ReadImageFile
	// Flatten image and remove alpha channel, to prevent alpha turning black in jpg
	if err := mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_FLATTEN); err != nil {
		return err
	}

	// Set any compression (100 = max quality)
	if err := mw.SetCompressionQuality(100); err != nil {
		return err
	}

	// Select only first page of pdf
	mw.SetIteratorIndex(0)

	// Convert into JPG
	if err := mw.SetFormat("jpg"); err != nil {
		return err
	}

	// Save File
	return mw.WriteImage(imageName)
}
