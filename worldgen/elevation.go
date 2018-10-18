package worldgen

import (
	"github.com/andpago/mrak/gui"
	"time"
)

func AddNoise(dst [][]float32, alpha float64, beta float64, n int, scale int, wt float32, chunkNum int, callback func()) {
	height := len(dst)
	width := len(dst[0])
	chunkSizeX := width / chunkNum
	chunkSizeY := height / chunkNum

	seed := int64(time.Now().Nanosecond())
	p := NewPerlin(alpha, beta, n, seed)

	for yChunk := 0; yChunk < chunkNum; yChunk++ {
		for xChunk := 0; xChunk < chunkNum; xChunk++ {
			for y := yChunk * chunkSizeY; y < (yChunk + 1) * chunkSizeY; y++ {
				for x := xChunk * chunkSizeX; x < (xChunk + 1) * chunkSizeX; x++ {
					dst[y][x] += wt * float32(p.Noise2D(float64(x * scale) /
						float64(width), float64(y * scale) / float64(height)))
				}
			}
			callback()
		}
	}
}

func GeneratePerlinElevation(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			w.ElevationMap[y][x] = 200
		}
	}
	AddNoise(w.ElevationMap, 2, 2, 3, 2, 50, 5, func(){Visualize(w, buf, vis)})
	AddNoise(w.ElevationMap, 2, 2, 3, 10, 10, 5, func(){Visualize(w, buf, vis)})
	AddNoise(w.ElevationMap, 2, 2, 3, 50, 2, 5, func(){Visualize(w, buf, vis)})
}
