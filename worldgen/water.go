package worldgen

import (
	"github.com/andpago/mrak/gui"
	"math"
)

func GenerateSimpleWaterlevel(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	const waterlevel = 0
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			w.IsWater[y][x] = w.ElevationMap[y][x] < waterlevel
		}
	}
	Visualize(w, buf, vis)
}

func CalculateWaterAdjacency(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	const radius = 20
	chunkNum := int(math.Ceil(float64(w.Width) / 300))
	chunkSizeX := w.Width / chunkNum
	chunkSizeY := w.Height / chunkNum


	var (
		mult, wad, X, Y int
	)

	for yChunk := 0; yChunk < chunkNum; yChunk++ {
		for xChunk := 0; xChunk < chunkNum; xChunk++ {
			for y := yChunk * chunkSizeY; y < (yChunk + 1) * chunkSizeY; y++ {
				for x := xChunk * chunkSizeX; x < (xChunk + 1) * chunkSizeX; x++ {
					wad = 0
					for dx := -radius; dx <= radius; dx++ {
						for dy := -radius; dy <= radius; dy++ {
							X = x + dx
							Y = y + dy

							if X < 0 || X >= w.Width || Y < 0 || Y >= w.Height {
								continue
							}

							mult = int(math.Sqrt(radius - math.Sqrt(float64(dx * dx + dy * dy))))
							if mult < 0 {
								continue
							}

							if w.IsWater[Y][X] {
								wad += mult
							}
						}
					}
					w.WaterAdjacency[y][x] = wad
					if wad > w.MaxWad {
						w.MaxWad = wad
					}
				}
			}
			Visualize(w, buf, vis)
		}
	}

}