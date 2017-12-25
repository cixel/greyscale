package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func readFile(name string) (image.Image, error) {
	reader, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := reader.Close()
		if err != nil {
			log.Println("problem closing reader")
		}
	}()

	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
		log.Fatal(err)
	}

	return img, nil
}

func greyscale(img image.Image) image.Image {
	b := img.Bounds()
	gray := image.NewGray(b)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			px := color.GrayModel.Convert(img.At(x, y))
			gray.Set(x, y, px)
		}
	}
	return gray
}

func writeFile(name string, img image.Image) error {
	writer, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		err := writer.Close()
		if err != nil {
			log.Println("problem closing writer")
		}
	}()

	err = jpeg.Encode(writer, img, &jpeg.Options{100})
	return err
}

func main() {
	img, err := readFile("testdata/rincewind.jpg")
	if err != nil {
		log.Panic(err)
	}

	gray := greyscale(img)

	err = writeFile("testdata/rincewind_gray.jpg", gray)
	if err != nil {
		log.Panic(err)
	}
}
