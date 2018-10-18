package worldgen

import (
	"github.com/andpago/mrak/gui"
	"time"
)

func GeneratePerlinElevation(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	const chunkNum = 5
	chunkSizeX := w.Width / chunkNum
	chunkSizeY := w.Height / chunkNum

	const (
		alpha = 2
		beta = 2
		n = 3
		scale = 5
	)

	seed := int64(time.Now().Nanosecond())

	p := NewPerlin(alpha, beta, n, seed)

	for yChunk := 0; yChunk < chunkNum; yChunk++ {
		for xChunk := 0; xChunk < chunkNum; xChunk++ {
			for y := yChunk * chunkSizeY; y < (yChunk + 1) * chunkSizeY; y++ {
				for x := xChunk * chunkSizeX; x < (xChunk + 1) * chunkSizeX; x++ {
					w.ElevationMap[y][x] = 50 + 100 * float32(p.Noise2D(float64(x * scale) /
						float64(w.Width), float64(y * scale) / float64(w.Height)))
				}
			}
			Visualize(w, buf, vis)
		}
	}
}
