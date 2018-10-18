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