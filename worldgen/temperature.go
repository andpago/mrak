package worldgen

import (
	"github.com/andpago/mrak/gui"
	"math"
	"time"
)

func GeneratePerlinLatitudeTemperature(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	const polarTemperature = 200
	const equatorialTemperature = 350
	height := float64(w.Height)
	width := float64(w.Width)
	semiHeight := height / 2

	const (
		alpha = 2
		beta = 2
		n = 3
		mult = 20
		scale = 5
	)
	seed := int64(time.Now().Nanosecond())
	p := NewPerlin(alpha, beta, n, seed)

	const chunkNum = 5
	chunkSizeX := w.Width / chunkNum
	chunkSizeY := w.Height / chunkNum

	for y := 0.0; y < height; y++ {
		for x := 0; x < w.Width; x++ {
			latitude := math.Abs(y - semiHeight) / semiHeight // from 0 to 1
			w.TemperatureMap[int(y)][x] =
				float32(equatorialTemperature + (polarTemperature - equatorialTemperature) * latitude)
		}
	}
	Visualize(w, buf, vis)

	for yChunk := 0; yChunk < chunkNum; yChunk++ {
		for xChunk := 0; xChunk < chunkNum; xChunk++ {
			for y := yChunk * chunkSizeY; y < (yChunk + 1) * chunkSizeY; y++ {
				for x := xChunk * chunkSizeX; x < (xChunk + 1) * chunkSizeX; x++ {
					w.TemperatureMap[int(y)][x] +=
							float32(p.Noise2D(float64(x) * scale / width, float64(y) * scale / height) * mult)
				}
			}
			Visualize(w, buf, vis)
		}
	}


}
