package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	width   = 800
	height  = 600
	maxIter = 1000
)

func must(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func mandelbrot(c complex128) uint8 {
	z := complex(0, 0)
	for i := 0; i < maxIter; i++ {
		z = z*z + c
		if cmplx.Abs(z) > 2 {
			return uint8(255 * i / maxIter)
		}
	}
	return 0
}

func julia(z complex128) uint8 {
	c := complex(-0.7, 0.27015)
	for i := 0; i < maxIter; i++ {
		z = z*z + c
		if cmplx.Abs(z) > 2 {
			return uint8(255 * i / maxIter)
		}
	}
	return 0
}

func drawFractal(name string) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	must(img != nil, "image allocation failed")

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			re := (float64(x) - width/2) * 4.0 / width
			im := (float64(y) - height/2) * 4.0 / width
			z := complex(re, im)

			var v uint8
			switch name {
			case "mandelbrot":
				v = mandelbrot(z)
			case "julia":
				v = julia(z)
			default:
				panic("unknown fractal type")
			}

			img.Set(x, y, color.RGBA{v, 0, 255 - v, 255})
		}
	}

	file, err := os.Create(name + ".png")
	must(err == nil, "failed to create file")
	defer file.Close()

	err = png.Encode(file, img)
	must(err == nil, "failed to encode png")
}

func main() {
	must(width > 0 && height > 0, "invalid image size")
	must(maxIter > 0, "invalid iteration count")

	drawFractal("mandelbrot")
	drawFractal("julia")
}
