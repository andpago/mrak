package worldgen

import (
	"github.com/andpago/mrak/gui"
	"math"
)

func GeneratePerlinHumidity(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	chunkNum := int(math.Ceil(float64(w.Width) / 2000))
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			w.HumidityMap[y][x] = 50
		}
	}
	AddNoise(w.HumidityMap, 2, 2, 3, 2, 50, chunkNum, func(){Visualize(w, buf, vis)})
	//AddNoise(w.HumidityMap, 2, 2, 3, 10, 10, 5, func(){Visualize(w, buf, vis)})
}
