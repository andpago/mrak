package worldgen

import (
	"github.com/andpago/mrak/gui"
	"math"
)

func GeneratePerlinLatitudeTemperature(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	const polarTemperature = 200
	const equatorialTemperature = 350
	height := float64(w.Height)
	semiHeight := height / 2

	chunkNum := int(math.Ceil(float64(w.Width) / 1000))


	for y := 0.0; y < height; y++ {
		for x := 0; x < w.Width; x++ {
			latitude := math.Abs(y - semiHeight) / semiHeight // from 0 to 1
			w.TemperatureMap[int(y)][x] =
				float32(equatorialTemperature + (polarTemperature - equatorialTemperature) * latitude)
		}
	}
	Visualize(w, buf, vis)

	AddNoise(w.TemperatureMap, 2, 2, 3, 5, 20, chunkNum, func(){Visualize(w, buf, vis)})

}
