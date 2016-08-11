package main

import (
	"fmt"
	"image/jpeg"
	"os"
)

func main() {
	fmt.Printf("Starting ASCII converter\n")
	filename := "sample.jpg"
	infile, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer infile.Close()

	img, err := jpeg.Decode(infile)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%#v\n", img.Bounds())
}
