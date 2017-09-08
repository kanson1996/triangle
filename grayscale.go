package triangulator

import (
	"image"
	"image/color"
)

func Grayscale(src image.Image) image.Image {
	dx, dy := src.Bounds().Max.X, src.Bounds().Max.Y
	dst := image.NewNRGBA(src.Bounds())
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			r, g, b, _ := src.At(x, y).RGBA()
			lum := float32(r) * 0.299 + float32(g) * 0.587 + float32(b) * 0.114
			pixel := color.Gray{uint8(lum / 256)}
			dst.Set(x, y, pixel)
		}
	}
	return dst
}