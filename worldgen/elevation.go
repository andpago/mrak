package worldgen

import (
	"github.com/andpago/mrak/gui"
)

func GeneratePerlinElevation(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			w.ElevationMap[y][x] = 0
		}
	}
	AddNoise(w.ElevationMap, 2, 2, 3, 2, 50, 5, func(){Visualize(w, buf, vis)})
	AddNoise(w.ElevationMap, 2, 2, 3, 10, 10, 5, func(){Visualize(w, buf, vis)})
	AddNoise(w.ElevationMap, 2, 2, 3, 50, 2, 5, func(){Visualize(w, buf, vis)})
}
