package worldgen

import (
	"github.com/andpago/mrak/gui"
	"image/color"
)

type Visualizer func(w *World, buf *gui.ProtectedColorBuffer)

func Visualize(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	buf.Mu.Lock()
	vis(w, buf)
	buf.Mu.Unlock()
}

func VisualizeElevationGrayscale(w *World, buf *gui.ProtectedColorBuffer) {
	for y := 0; y < len(buf.Colors); y++ {
		for x := 0; x < len(buf.Colors[0]); x++ {
			gray := uint8(w.ElevationMap[y * w.Height / len(buf.Colors)][x * w.Width / len(buf.Colors[0])])
			buf.Colors[y][x] = color.Gray{gray}
		}
	}
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
		orangeStart = 330
		redStart = 370
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
				buf.Colors[y][x] = color.RGBA{0, 0, uint8(relT * 255), 255}
			} else if t < orangeStart {
				// green
				relT := (t - greenStart) / (orangeStart - greenStart)
				buf.Colors[y][x] = color.RGBA{0, uint8(relT * 255), 0, 255}
			} else if t < redStart {
				// orange
				relT := (t - orangeStart) / (redStart - orangeStart)
				buf.Colors[y][x] = color.RGBA{uint8(relT * 255), uint8(relT * 127), 0, 255}
			} else {
				// red
				relT := (t - redStart) / (maxTemp - redStart)
				buf.Colors[y][x] = color.RGBA{uint8(relT * 255), 0, 0, 255}
			}
		}
	}
}