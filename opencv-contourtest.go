package main

import (
	"gocv.io/x/gocv"
)

func opencv() {
	var output gocv.Mat

	img := gocv.IMRead("invoice.jpg", gocv.IMReadColor)
	gocv.Canny(img, &output, 100, 200)
}
