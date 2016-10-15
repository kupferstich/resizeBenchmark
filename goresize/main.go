package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

var filters = map[string]imaging.ResampleFilter{
	"NearestNeighbor":   imaging.NearestNeighbor,
	"Box":               imaging.Box,
	"Linear":            imaging.Linear,
	"MitchellNetravali": imaging.MitchellNetravali,
	"CatmullRom":        imaging.CatmullRom,
	"Gaussian":          imaging.Gaussian,
	"Lanczos":           imaging.Lanczos,
}

var pics = []string{
	"test1.jpg",
	"test2.jpg",
	"test3.jpg",
}

func main() {
	for _, p := range pics {
		img := openPic(p)
		for name, filter := range filters {
			base := filepath.Base(p)
			outname := filepath.Join(
				"out",
				base,
				fmt.Sprintf("%s_%s", name, p),
			)
			os.MkdirAll(filepath.Dir(outname), 0777)
			resizePic(img, filter, outname)
		}

	}

}

func openPic(picturePath string) image.Image {

	file, err := os.Open(filepath.Join("src", picturePath))
	defer file.Close()
	if err != nil {
		log.Println(err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println(err)
	}
	return img
}

func resizePic(img image.Image, filter imaging.ResampleFilter, out string) {
	log.Println(out)
	//dstImage128 := imaging.Resize(img, 128, 128, imaging.Lanczos)
	thumb := imaging.Thumbnail(img, 700, 700, filter)
	err := imaging.Save(thumb, out)
	if err != nil {
		log.Println(err)
	}
}
