package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	leptonica "gopkg.in/GeertJohan/go.leptonica.v1"
	tesseract "gopkg.in/GeertJohan/go.tesseract.v1"
)

func main() {
	err := filepath.Walk("../outputtemp",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.Contains(path, ".jpg") {
				fmt.Println(path, info.Size())
				masker := fmt.Sprintf("%s mask.png", path)

				rectImage := image.NewRGBA(image.Rect(0, 0, 2480, 3508))
				zwart := color.RGBA{0, 0, 0, 255}
				letterruimte := color.RGBA{255, 255, 255, 255}
				draw.Draw(rectImage, rectImage.Bounds(), &image.Uniform{zwart}, image.ZP, draw.Src)

				tessdata_prefix := os.Getenv("TESSDATA_PREFIX")
				if tessdata_prefix == "" {
					tessdata_prefix = "/usr/share/tesseract-ocr/4.00/"
				}
				t, err := tesseract.NewTess(filepath.Join(tessdata_prefix, "tessdata"), "eng")
				if err != nil {
					log.Fatalf("Error while initializing Tess: %s\n", err)
				}
				defer t.Close()

				// open a new Pix from file with leptonica
				pix, err := leptonica.NewPixFromFile(path)
				if err != nil {
					log.Fatalf("Error while getting pix from file: %s\n", err)
				}
				defer pix.Close() // remember to cleanup

				// set the page seg mode to autodetect
				t.SetPageSegMode(tesseract.PSM_AUTO_OSD)

				// set the image to the tesseract instance
				t.SetImagePix(pix)

				// retrieve text from the tesseract instance
				fmt.Println(t.Text())

				// // retrieve text from the tesseract instance
				// fmt.Println(t.HOCRText(0))

				// retrieve text from the tesseract instance
				boxText, _ := t.BoxText(0)

				for _, char := range boxText.Characters {
					//letterImage := image.NewRGBA(image.Rect(int(char.StartX), int(char.StartY), int(char.EndX), int(char.EndY)))

					for j := char.EndY; j > char.StartY; j++ {
						rectImage.Set(int(char.StartX), int(j), letterruimte)
						for k := char.StartX; k < char.EndX; k++ {
							rectImage.Set(int(k), int(j), letterruimte)

						}
					}

					s := fmt.Sprintf("%c", char.Character)

					fmt.Printf("%s ", s)
				}
				fmt.Printf("\n")
				file, err := os.Create(masker)
				if err != nil {
					log.Fatalf("failed create file: %s", err)
				}
				png.Encode(file, rectImage)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	// files, err := ioutil.ReadDir("../outputtemp")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, f := range files {
	// 	fmt.Println(f.Name())
	// }

}

//func (tess *Tess) BoxText(pagenumber int) (*BoxText, error)
//tesseract.BoxCharacter
