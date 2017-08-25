package yuyvimport

import (
	"image"
	"image/color"
	"math"
)

func getRGB(Y byte, U byte, V byte) (c color.RGBA) {
	yy := float64(Y)
	uu := float64(U)
	vv := float64(V)

	rr := 1.164*(yy-16.0) + 1.596*(vv-128.0)
	gg := 1.164*(yy-16.0) - 0.813*(vv-128.0) - 0.391*(uu-128.0)
	bb := 1.164*(yy-16.0) + 2.018*(uu-128.0)

	if rr < 0.0 {
		rr = 0.0
	}

	if gg < 0.0 {
		gg = 0.0
	}

	if bb < 0.0 {
		bb = 0.0
	}

	c.R = byte(math.Floor(math.Abs(rr)))
	c.G = byte(math.Floor(math.Abs(gg)))
	c.B = byte(math.Floor(math.Abs(bb)))

	return
}

func loadYUYV(data []byte) (rgb []color.RGBA) {

	idx := 0
	for {
		Y0 := data[idx]
		idx++
		U0 := data[idx]
		idx++
		Y1 := data[idx]
		idx++
		V0 := data[idx]
		idx++

		rgb = append(rgb, getRGB(Y0, U0, V0))
		rgb = append(rgb, getRGB(Y1, U0, V0))

		if idx+4 > len(data) {
			break
		}
	}
	return
}

// Import reads yuyv encoded data from yuyvData and creates an RGBA image with size w x h
func Import(w int, h int, yuyvData []byte) image.Image {
	rgba := loadYUYV(yuyvData)

	var size image.Rectangle
	size.Min.X = 0
	size.Min.Y = 0
	size.Max.X = w
	size.Max.Y = h
	img := image.NewRGBA(size)

	x := 0
	y := 0
	idx := 0
	for {
		x = 0
		for {
			img.Set(x, y, rgba[idx])
			idx++

			x++
			if x >= size.Max.X {
				break
			}
		}
		y++
		if y >= size.Max.Y {
			break
		}
	}

	return img.SubImage(img.Bounds())
}
