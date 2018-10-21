package worldgen

import (
	"fmt"
	"github.com/andpago/mrak/gui"
	"golang.org/x/image/colornames"
	"image/color"
)

type Visualizer func(w *World, buf *gui.ProtectedColorBuffer)

func Visualize(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	buf.Mu.Lock()
	vis(w, buf)
	buf.Mu.Unlock()
}

func VisualizeElevationGrayscaleWithAlpha(w *World, buf *gui.ProtectedColorBuffer, alpha uint8) {
	var (
		maxElevation = float32(0)
		minElevation = float32(0)
	)

	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			el := w.ElevationMap[y][x]
			if el > maxElevation {
				maxElevation = el
			}
			if el < minElevation {
				minElevation = el
			}
		}
	}

	for y := 0; y < len(buf.Colors); y++ {
		for x := 0; x < len(buf.Colors[0]); x++ {
			gray := uint8((w.ElevationMap[y * w.Height / len(buf.Colors)][x * w.Width / len(buf.Colors[0])] -
				minElevation ) * 255 / (maxElevation - minElevation))
			buf.Colors[y][x] = color.RGBA{gray, gray, gray, alpha}
		}
	}
}

func VisualizeElevationGrayscale(w *World, buf *gui.ProtectedColorBuffer) {
	VisualizeElevationGrayscaleWithAlpha(w, buf, 255)
}

func VisualizeWaterLevel(w *World, buf *gui.ProtectedColorBuffer) {
	for y := 0; y < len(buf.Colors); y++ {
		for x := 0; x < len(buf.Colors[0]); x++ {
			X := x * w.Width / len(buf.Colors[0])
			Y := y * w.Height / len(buf.Colors)
			if w.IsWater[Y][X] {
				b := uint8(128 + w.ElevationMap[Y][X])
				buf.Colors[y][x] = color.RGBA{0, 0, b, 255}
			}
		}
	}
}

func VisualizeTemerature(w *World, buf *gui.ProtectedColorBuffer) {
	const (
		greenStart = 270
		orangeStart = 320
		redStart = 340
		maxTemp = 500
	)


	for y := 0; y < len(buf.Colors); y++ {
		for x := 0; x < len(buf.Colors[0]); x++ {
			X := x * w.Width / len(buf.Colors[0])
			Y := y * w.Height / len(buf.Colors)

			t := w.TemperatureMap[Y][X]

			if t < greenStart {
				// blue
				relT := (t - 0) / (greenStart - 0)
				buf.Colors[y][x] = color.RGBA{0, uint8(64 * relT), uint8(255 * (1 - relT)), 255}
			} else if t < orangeStart {
				// green
				relT := (t - greenStart) / (orangeStart - greenStart)
				buf.Colors[y][x] = color.RGBA{uint8(relT * 127), uint8(relT * 63) + 64, 0, 255}
			} else if t < redStart {
				// orange
				relT := (t - orangeStart) / (redStart - orangeStart)
				buf.Colors[y][x] = color.RGBA{uint8(relT * 63) + 127, uint8(relT * 63) + 128, 0, 255}
			} else {
				// red
				relT := (t - redStart) / (maxTemp - redStart)
				buf.Colors[y][x] = color.RGBA{
					uint8(relT * 63) + 127 + 63,
					192 - uint8(relT * 192),
					0,
					255}
			}
		}
	}
}

func BlendRGBColorBuffers(buffers []*gui.ProtectedColorBuffer, weights []uint32) *gui.ProtectedColorBuffer {
	if len(buffers) == 0 {
		panic("cannot blend colors: unable to infer dimensions from no buffers")
	}

	wsum := uint32(0)
	for _, wt := range weights {
		wsum += wt
	}

	buffers[0].Mu.Lock()
	H := len(buffers[0].Colors)
	W := len(buffers[0].Colors[0])
	R, G, B := make([][]uint32, H, H), make([][]uint32, H, H), make([][]uint32, H, H)
	for y := 0; y < H; y++ {
		R[y] = make([]uint32, W, W)
		G[y] = make([]uint32, W, W)
		B[y] = make([]uint32, W, W)
	}
	buffers[0].Mu.Unlock()

	for i, buf := range buffers {
		buf.Mu.Lock()
		for y := 0; y < H; y++ {
			for x := 0; x < W; x++ {
				r, g, b, a := buf.Colors[y][x].RGBA()
				const maxA = 0xffff
				R[y][x] += (r * a)/ maxA * weights[i] / wsum
				G[y][x] += (g * a)/ maxA * weights[i] / wsum
				B[y][x] += (b * a)/ maxA * weights[i] / wsum
			}
		}
		buf.Mu.Unlock()
	}

	res := gui.NewProtectedColorBuffer(W, H, colornames.White)
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			res.Colors[y][x] = color.RGBA{
				uint8(R[y][x] * 255 / 0xffff),
				uint8(G[y][x] * 255 / 0xffff),
				uint8(B[y][x] * 255 / 0xffff),
				0xff}
		}
	}


	return res
}

func NoVisualizer(w *World, buf *gui.ProtectedColorBuffer) {

}

func VisualizeBlendAll(w *World, buf *gui.ProtectedColorBuffer) {
	bufs := make([]*gui.ProtectedColorBuffer, 3, 3)

	for i := range bufs {
		bufs[i] = gui.NewProtectedColorBuffer(len(buf.Colors[0]), len(buf.Colors), color.Transparent)
	}

	VisualizeElevationGrayscale(w, bufs[0])
	VisualizeTemerature(w, bufs[1])
	VisualizeWaterLevel(w, bufs[2])
	res := BlendRGBColorBuffers(bufs, []uint32{1, 5, 5})

	buf.Colors = res.Colors
}


func VisualizeAll(w *World, buf *gui.ProtectedColorBuffer) {
	VisualizeTemerature(w, buf)
	VisualizeWaterLevel(w, buf)
}

func VisualizeHumidity(w *World, buf *gui.ProtectedColorBuffer) {
	const (
		redStart = 50
		maxHumidity = 100
	)


	for y := 0; y < len(buf.Colors); y++ {
		for x := 0; x < len(buf.Colors[0]); x++ {
			X := x * w.Width / len(buf.Colors[0])
			Y := y * w.Height / len(buf.Colors)

			hum := w.HumidityMap[Y][X]

			if hum < redStart {
				// blue
				relHum := (hum - 0) / (redStart - 0)
				buf.Colors[y][x] = color.RGBA{uint8(127.0 * relHum), 0, uint8(127.0 * (1 - relHum)) , 255}
			} else {
				// red
				relHum := (hum - redStart) / (maxHumidity - redStart)
				buf.Colors[y][x] = color.RGBA{uint8(127.0 + 127.0 * relHum), 0, 0, 255}
			}
		}
	}
}

func VisualizeClimate(w *World, buf *gui.ProtectedColorBuffer) {
	for y := 0; y < len(buf.Colors); y++ {
		for x := 0; x < len(buf.Colors[0]); x++ {
			X := x * w.Width / len(buf.Colors[0])
			Y := y * w.Height / len(buf.Colors)

			if w.IsWater[Y][X] {
				if w.TemperatureMap[Y][X] < 250 {
					buf.Colors[y][x] = colornames.Aqua // ice
				} else {
					buf.Colors[y][x] = colornames.Blue // water
				}
				continue
			}

			if w.TemperatureMap[Y][X] < 250 {
				buf.Colors[y][x] = colornames.White // polar cap
				continue
			}

			if w.ElevationMap[Y][X] < 20 {
				buf.Colors[y][x] = colornames.Yellow // coast
				continue
			}

			if w.HumidityMap[Y][X] < 25 {
				if w.TemperatureMap[Y][X] > 320 {
					buf.Colors[y][x] = colornames.Yellow // desert
				} else {
					buf.Colors[y][x] = colornames.Gray // badlands
				}
				continue
			}

			if w.ElevationMap[Y][X] > 1000 {
				buf.Colors[y][x] = colornames.Lightgreen // alpine forest
				continue
			}

			if w.HumidityMap[Y][X] > 70 && w.ElevationMap[Y][X] < 100 {
				buf.Colors[y][x] = colornames.Darkgreen // bog
				continue
			}

			if w.ElevationMap[Y][X] > 1200 {
				buf.Colors[y][x] = colornames.White // mountain
				continue
			}

			buf.Colors[y][x] = colornames.Green // temperate forest
		}
	}

	const alpha = 100
	alphaBuffer := gui.NewProtectedColorBuffer(len(buf.Colors[0]), len(buf.Colors), color.Transparent)
	VisualizeElevationGrayscaleWithAlpha(w, alphaBuffer, alpha)

	for y := 0; y < len(buf.Colors); y++ {
		for x := 0; x < len(buf.Colors[0]); x++ {

			r, g, b, _ := buf.Colors[y][x].RGBA()
			r /= 255
			g /= 255
			b /= 255

			A, _, _, _ := alphaBuffer.Colors[y][x].RGBA()
			A /= 255

			buf.Colors[y][x] = color.RGBA{
				uint8((r * (255 - alpha) + A * alpha) / 255),
				uint8((g * (255 - alpha) + A * alpha) / 255),
				uint8((b * (255 - alpha) + A * alpha) / 255),
				255,
			}
		}
	}
}

func VisualizeWaterAdjacency(w *World, buf *gui.ProtectedColorBuffer) {
	maxWad := w.MaxWad

	for y := 0; y < len(buf.Colors); y++ {
		for x := 0; x < len(buf.Colors[0]); x++ {
			X := x * w.Width / len(buf.Colors[0])
			Y := y * w.Height / len(buf.Colors)

			buf.Colors[y][x] = color.Gray{uint8(float64(w.WaterAdjacency[Y][X]) * 255.0 / float64(maxWad))}
		}
	}
}

func VisualizeDistanceToWater(w *World, buf *gui.ProtectedColorBuffer) {
	for y := 0; y < len(buf.Colors); y++ {
		for x := 0; x < len(buf.Colors[0]); x++ {
			X := x * w.Width / len(buf.Colors[0])
			Y := y * w.Height / len(buf.Colors)

			buf.Colors[y][x] = color.Gray{uint8(float64(w.DistanceToWater[Y][X]))}
		}
	}
}

func VisualizeContinents(w *World, buf *gui.ProtectedColorBuffer) {
	VisualizeClimate(w, buf)

	colors := []color.Color{
		colornames.Red,
		colornames.Green,
		colornames.Blue,
		colornames.Black,
		colornames.White,
		colornames.Yellow,
		colornames.Violet,
	}

	for i, c := range w.Continets {
		for _, p := range c.Points {
			X := p.X * len(buf.Colors[0]) / w.Width
			Y := p.Y * len(buf.Colors) / w.Height

			//fmt.Println(X, Y)

			buf.Colors[Y][X] = colors[i % len(colors)]
		}
	}

	fmt.Println("continents drawn")
}