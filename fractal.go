package main

import (
       "image"
       "image/color"
       "image/png"
       "math/cmplx"
       "log"
       "os"
)

var (
  numSteps int = 1024
)

func RenderFractal() image.Image {
  minX, maxX, minY, maxY := -512, 512, -512, 512
  image := image.NewNRGBA(image.Rect(minX, minY, maxX, maxY))

  xStepSize := (maxX - minX) / numSteps
  yStepSize := (maxY - minY) / numSteps

  for currX := minX; currX < maxX; currX += xStepSize {
    for currY := minY; currY < maxY; currY += yStepSize {
      currColor := GetColor(currX, currY)
      image.SetNRGBA(int(currX), int(currY), currColor)
    }
  }

  return image
}

func GetColor(x, y int) color.NRGBA {
  z := complex(float64(x) / 256.0, float64(y) / 256)
  c := z
  steps := 0
  for ; steps < 127 && cmplx.Abs(z) < 2; steps += 1 {
    z = z * z + c
  }

 return color.NRGBA{uint8(steps), 127, 127, 127}
}

func SaveImage(image image.Image) {
  // Try to open the file, die if it doesn't work.
  file, err := os.Create("fractal.png")
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