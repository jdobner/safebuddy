package main

import (
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/rostislaved/screenshot"
)

func main() {

	//screenshoter := screenshot.New()

	img, err := screenshot.CaptureScreen()
	if err != nil {
		log.Fatal(err)
	}
	filename := filepath.Join(os.TempDir(), "screenshot.png")
	log.Println("writing file as ", filename)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}

}
