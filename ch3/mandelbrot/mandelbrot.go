// Package mandelbrot 生成一个PNG格式的 Mandelbrot 分形图
package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

func Mandelbrot(writer io.Writer) {
	const (
		xmin, ymin, xmax, ymax float64 = -2, -2, 2, 2
		width, height                  = 1024, 1024
	)

	//e3.6 superSampling
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py += 2 {
		y1 := float64(py)/height*(ymax-ymin) + ymin
		y2 := float64(py+1)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px += 2 {
			x1 := float64(px)/width*(xmax-xmin) + xmin
			x2 := float64(px+1)/width*(xmax-xmin) + xmin

			z1 := complex(x1, y1)
			z2 := complex(x1, y2)
			z3 := complex(x2, y1)
			z4 := complex(x2, y2)

			ssMap := make([]color.Color, 0)
			ssMap = append(ssMap, mandelbrot(z1), mandelbrot(z2),
				mandelbrot(z3), mandelbrot(z4))

			img.Set(px, py, superSampling(ssMap))
		}
	}
	png.Encode(writer, img)
}

func superSampling(ssMap []color.Color) color.Color {
	var r, g, b, a uint32
	n := len(ssMap)
	for _, c := range ssMap {
		_r, _g, _b, _a := c.RGBA()
		r += _r / uint32(n)
		g += _g / uint32(n)
		b += _b / uint32(n)
		a += _a / uint32(n)
	}
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//e3.5 change color
			switch {
			case n < 10:
				return color.RGBA{B: 120, A: 255}
			case n > 10 && n < 100:
				return color.RGBA{R: 190, G: 80, B: 120}
			case n > 100 && n < 150:
				return color.RGBA{R: 90, G: 80, B: 120}
			default:
				return color.Gray{Y: 255 - contrast*n}
			}
			//return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.RGBA{R: 100, A: 255}
	//return color.Black
}

// E3 e.3.7 newton algo
func E3(writer io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(writer, img) // NOTE: ignoring errors
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{Y: 255 - contrast*i}
		}
	}
	return color.Black
}
