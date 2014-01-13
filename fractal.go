package main

import (
       "image"
       "image/color"
       "image/png"
       "math"
       "math/cmplx"
       "log"
       "os"
)

var (
  factor int = 1024
  numSteps float64 = float64(4 * factor)
)

func RenderFractal() image.Image {
  minX, maxX, minY, maxY := -2, 2, -2, 2
  image := image.NewNRGBA(image.Rect(factor * minX, factor * minY,
  factor * maxX, factor * maxY))

  xStepSize := float64(maxX - minX) / numSteps
  yStepSize := float64(maxY - minY) / numSteps

  for currX := float64(-2); currX < float64(2); currX += xStepSize {
    for currY := float64(-2); currY < float64(2); currY += yStepSize {
      currColor := GetColor(currX, currY)
      image.SetNRGBA(int(float64(factor) * currX), int(float64(factor) * currY), currColor)
    }
  }

  return image
}

func GetColor(x, y float64) color.NRGBA {
  z := complex(x, y)
  c := z
  steps := 0
  for ; steps < 255 && cmplx.Abs(z) < 2; steps += 1 {
    z = z * z + c
  }
  colorStep := 128 - int(math.Log(float64(x)))

 return color.NRGBA{uint8(colorStep), 128 - uint8(steps % 128), 127, 128}
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