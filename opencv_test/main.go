package main

import (
	"fmt"
	opencv "gocv.io/x/gocv"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func main() {
	f, err := os.Open("opencv_test/test.jpeg")
	if err != nil {
		fmt.Println("open err:", err)
		return
	}

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("Decode err:", err)
		return
	}

	rectangle := image.Rect(250, 750, 350, 850)

	var col = color.RGBA{
		R: 144,
		G: 238,
		B: 144,
		A: 10,
	}

	mat, err := opencv.ImageToMatRGBA(img)
	if err != nil {
		fmt.Println("ImageToMatRGBA err:", err)
		return
	}

	size := mat.Size()
	fmt.Println("size:", size)

	opencv.Rectangle(&mat, rectangle, col, 3)

	img2, err := mat.ToImage()
	if err != nil {
		fmt.Println("ToImage err:", err)
		return
	}

	w, err := os.Create("opencv_test/after.jpeg")
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}

	err = jpeg.Encode(w, img2, nil)
	if err != nil {
		fmt.Println("Encode err:", err)
		return
	}

}
