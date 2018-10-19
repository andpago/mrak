package worldgen

import (
	"github.com/andpago/mrak/gui"
	"math"
)

func GeneratePerlinHumidity(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	chunkNum := int(math.Ceil(float64(w.Width) / 1000))
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			w.HumidityMap[y][x] = 50
		}
	}
	AddNoise(w.HumidityMap, 2, 2, 3, 2, 50, chunkNum, func(){Visualize(w, buf, vis)})
	AddNoise(w.HumidityMap, 2, 2, 3, 10, 10, chunkNum, func(){Visualize(w, buf, vis)})
}


func GeneratePerlinWadHumidity(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	GeneratePerlinHumidity(w, buf, vis)

	const maxDistanceEffect = 10

	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			w.HumidityMap[y][x] += 60.0 * (float32(w.WaterAdjacency[y][x]) / float32(w.MaxWad))


			distanceEffect := float32(w.DistanceToWater[y][x]) / 10
			if distanceEffect > maxDistanceEffect {
				distanceEffect = maxDistanceEffect
			}


			w.HumidityMap[y][x] -= distanceEffect
			if w.HumidityMap[y][x] > 100 {
				w.HumidityMap[y][x] = 100
			}
			if w.HumidityMap[y][x] < 0 {
				w.HumidityMap[y][x] = 0
			}
		}
	}

	Visualize(w, buf, vis)
}