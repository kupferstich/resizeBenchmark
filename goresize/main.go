package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"

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
		resizePic(img)
	}

}

func openPic(picturePath string) image.Image {
	file, err := os.Open(picturePath)
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

func resizePic(img image.Image) {
	//dstImage128 := imaging.Resize(img, 128, 128, imaging.Lanczos)
	thumb := imaging.Thumbnail(img, 700, 700, imaging.CatmullRom)
	err := imaging.Save(thumb, "out/dst.jpg")
	if err != nil {
		panic(err)
	}
}
