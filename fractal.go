package main

import (
       "image"
       "image/color"
       "image/png"
       "os"
       "png"
)

var (
  numSteps int = 1024
)

func RenderFractoal() image.Image {
  minX, maxX, minY, maxY = -512, 512, -512, 512
  image := image.NewNRGBA(Rect(minX, minY, maxX, maxY))

  xStepSize := float64(maxX - minX) / numSteps
  yStepSize := float64(maxY - minY) / numSteps

  for currX := minX; currX < maxX; currX += xStepSize {
    for currY := minY; currY < maxY; currY += yStepSize {
      currColor := color.NRGBA(int(currX) % 128, int(currY) % 128, 0, 0)
      image.SetNRGBA(int(currX), int(currY), currColor)
    }
  }

  return image
}

func SaveImage(image image.Image) {
  // Try to open the file, die if it doesn't work.
  file, err = os.Create("fractal.png")
  if err != nil {
    log.Fatal(err)
  }
  // Make sure that the file closes eventually.
  defer file.Close()

  png.Encode(file, image)
}

func main() {
     image := RenderFractal()
     SaveImage(image)
}