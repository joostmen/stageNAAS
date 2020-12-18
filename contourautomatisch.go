package main

import (
	"fmt"

	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
)

func main() {

	naam := "invoice"
	extend := ""
	img, err := imgio.Open(fmt.Sprintf("%s.jpg", naam))
	if err != nil {
		fmt.Println(err)
		return
	}
	result := effect.EdgeDetection(img, 1.0)
	for i := 0; i < 3; i++ {

		switch i {
		case 0:
			result = effect.Invert(img)
			extend = "Invert"
		case 1:
			result = effect.Sobel(img)
			extend = "Sobel"
		case 2:
			result = effect.EdgeDetection(img, 5.0)
			extend = "EdgeDetection5.0"
		}

		if err := imgio.Save(fmt.Sprintf("%s-%s.png", naam, extend), result, imgio.PNGEncoder()); err != nil {
			fmt.Println(err)
			return
		}
	}
}
