package worldgen

import (
	"fmt"
	"github.com/andpago/mrak/gui"
	"math"
)

func GeneratePerlinElevation(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	chunkNum := int(math.Ceil(float64(w.Width) / 1000))
	minLandFraction := 0.25
	maxLandFraction := 0.45

	landFraction := 2.0


	for landFraction < minLandFraction || landFraction > maxLandFraction {
		land := 0

		for y := 0; y < w.Height; y++ {
			for x := 0; x < w.Width; x++ {
				w.ElevationMap[y][x] = 0
			}
		}
		AddNoise(w.ElevationMap, 2, 2, 3, 1, 1000, chunkNum, func() { Visualize(w, buf, vis) })
		AddNoise(w.ElevationMap, 2, 2, 3, 5, 1000, chunkNum, func() { Visualize(w, buf, vis) })
		AddNoise(w.ElevationMap, 2, 2, 3, 10, 200, chunkNum, func() { Visualize(w, buf, vis) })
		AddNoise(w.ElevationMap, 2, 2, 3, 50, 100, chunkNum, func() { Visualize(w, buf, vis) })

		for y := 0; y < w.Height; y++ {
			for x := 0; x < w.Width; x++ {
				if w.ElevationMap[y][x] > 0 {
					land++
				}
			}
		}

		landFraction = float64(land) / float64(w.Width * w.Height)
		if landFraction < minLandFraction || landFraction > maxLandFraction {
			fmt.Printf("rejected world with landFraction = %.2v\n", landFraction)
		}
	}
}
