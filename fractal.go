/*
Compute the mandelbrot and write it out to disk.
Nothing special.
*/

package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
)

var (
	// How dense do we want our pixels? n^2 slow down here.
	factor *int = flag.Int("pixel_density", 1024, "how many pixels wide to make the image")
	// 4 since we are running from -2 to 2.
	numSteps float64 = float64(4 * *factor)
)

func RenderFractal() image.Image {
	// Logical bounds are [-2, 2], but we want lots of pixels.
	minX, maxX, minY, maxY := -2, 2, -2, 2
	image := image.NewNRGBA(image.Rect(*factor*minX, *factor*minY,
		*factor*maxX, *factor*maxY))

	// 1 step per pixel.
	xStepSize := float64(maxX-minX) / numSteps
	yStepSize := float64(maxY-minY) / numSteps

	// Walk through pixel space, compute each value.
	for currX := float64(-2); currX < float64(2); currX += xStepSize {
		for currY := float64(-2); currY < float64(2); currY += yStepSize {
			currColor := GetColor(currX, currY)
			image.SetNRGBA(int(float64(*factor)*currX),
				int(float64(*factor)*currY), currColor)
		}
	}

	return image
}

// Compute the color to show for the pixel at (x,y)
func GetColor(x, y float64) color.NRGBA {
	z := complex(x, y)
	c := z
	// Compute the number of steps util the pixel "escapes".
	// The boundary is what creates all of the coolness.
	steps := 0
	for ; steps < 1024 && cmplx.Abs(z) < 2; steps += 1 {
		z = z*z + c
	}
	colorStep := 128 - int(math.Log(float64(x)))

	return color.NRGBA{uint8(colorStep), 128 - uint8(steps%128), 127, 128}
}

func SaveImage(image image.Image) {
	// Try to open the file, die if it doesn't work.
	file, err := os.Create("fractal.png")
	if err != nil {
		log.Fatal(err)
	}
	// Make sure that the file closes eventually.
	defer file.Close()
	// Write out the image data.
	png.Encode(file, image)
}

func main() {
	flag.Parse()
	// Hard work, compute the image.
	image := RenderFractal()
	// Save it out to disk.
	SaveImage(image)
}
