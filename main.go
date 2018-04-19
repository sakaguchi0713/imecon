package main

import (
	"os"
	"image"
	"log"
	"image/jpeg"
	_ "image/png"
)


func changeImage(path string) image.Image {
	f, _ := os.Open(path)
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func main() {
	img := changeImage("images/png.png")

	c, err := os.Create("./images/huga.jpeg")
	if err != nil {
		panic(err)
	}

	jpeg.Encode(c, img, nil)
}