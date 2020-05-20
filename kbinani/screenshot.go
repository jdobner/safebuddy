package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/kbinani/screenshot"
)

// save *image.RGBA to filePath with PNG format.
func save(img *image.RGBA, fileName string) {
	fileName = filepath.Join(os.TempDir(), fileName)
	log.Println("Saving to ", fileName)
	file, err := os.Create(fileName)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func main() {
	capture(4)
}

func capture(delay int32) {
	log.Println("slleping for ", delay)
	time.Sleep(4 * time.Second)
	// Capture each displays.
	n := screenshot.NumActiveDisplays()
	log.Println("Number of screens: ", n)
	if n <= 0 {
		log.Panic("Active display not found")
	}

	var all image.Rectangle = image.Rect(0, 0, 0, 0)

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		all = bounds.Union(all)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())

		save(img, fileName)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}

	// Capture all desktop region into an image.
	fmt.Printf("%v\n", all)
	img, err := screenshot.Capture(all.Min.X, all.Min.Y, all.Dx(), all.Dy())
	if err != nil {
		panic(err)
	}
	save(img, "all.png")
}
