package worldgen

import "github.com/andpago/mrak/gui"

func GenerateSimpleWaterlevel(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	const waterlevel = 205
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			w.IsWater[y][x] = w.ElevationMap[y][x] < waterlevel
		}
	}
	Visualize(w, buf, vis)
}
