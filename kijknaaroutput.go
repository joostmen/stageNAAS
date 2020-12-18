package main

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func g() {
	output := "output.png"
	fSrc, err := os.Open("bol000.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer fSrc.Close()
	// src, err := png.Decode(fSrc)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	Mask, err := os.Open("mask.png")
	if err != nil {
		log.Fatal(err)
	}
	defer Mask.Close()
	sourceFile, _ := png.Decode(fSrc)
	maskFile, _ := png.Decode(Mask)
	dst := image.NewRGBA(image.Rect(0, 0, 2480, 3508))
	//green := image.NewUniform(color.RGBA{0x00, 0x1f, 0x00, 0xff})
	draw.DrawMask(dst, dst.Bounds(), sourceFile, image.ZP, maskFile, image.ZP, draw.Src)

	file, err := os.Create(output)
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	png.Encode(file, dst)
}
