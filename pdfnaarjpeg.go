package main

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"

	"github.com/gen2brain/go-fitz"
	"github.com/otiai10/gosseract/v2"
)

/*
stap 1: vind een manier om map volledig over te nemen, vervolgens die mappenstructuur te kopieren naar een output achtige plek
stap 2: probeer op deze manier alle data naar text om te zetten en vervolgens te OCRen.
stap 3: probeer met tfidf een standaard te maken per type document.

als alles fout gaat moet de data handmatig worden verzameld uit outputtemp
*/
func main() {

	files, err := ioutil.ReadDir("data")
	if err != nil {
		log.Fatal(err)
	}
	os.MkdirAll("dataout", 1)
	for _, file := range files {

		doc, err := fitz.New(fmt.Sprintf("data/%s", file.Name()))
		if err != nil {
			panic(err)
		}

		img, err := doc.Image(0)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(fmt.Sprintf("dataout/%s.jpg", file.Name()))
		if err != nil {
			panic(err)
		}
		defer f.Close()

		jpeg.Encode(f, img, &jpeg.Options{jpeg.DefaultQuality})

		client := gosseract.NewClient()
		defer client.Close()

		client.SetImage(fmt.Sprintf("dataout/%s.jpg", file.Name()))
		text, _ := client.Text()

		output, _ := os.Create(fmt.Sprintf(fmt.Sprintf("dataout/%s.txt", file.Name())))
		output.WriteString(text)
		output.Close()

		doc.Close()

	} //print alle mappen die in gesorteert staan.

}
