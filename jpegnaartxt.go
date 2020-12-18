package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/otiai10/gosseract/v2"
)

func main() {

	files, err := ioutil.ReadDir("dataout")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {

		client := gosseract.NewClient()
		defer client.Close()

		client.SetImage(fmt.Sprintf("dataout/%s", file.Name()))
		text, _ := client.Text()

		output, _ := os.Create(fmt.Sprintf(fmt.Sprintf("dataout/%s.txt", file.Name())))
		output.WriteString(text)
		output.Close()

	} //print alle mappen die in gesorteert staan.

}
