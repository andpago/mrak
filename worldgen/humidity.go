package worldgen

import "github.com/andpago/mrak/gui"

func GeneratePerlinHumidity(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			w.HumidityMap[y][x] = 50
		}
	}
	AddNoise(w.HumidityMap, 2, 2, 3, 2, 30, 5, func(){Visualize(w, buf, vis)})
	//AddNoise(w.HumidityMap, 2, 2, 3, 10, 10, 5, func(){Visualize(w, buf, vis)})
}
