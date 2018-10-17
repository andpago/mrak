package worldgen

type World struct {
	Name string
	Width, Height int // kilometers ?
	ElevationMap [][]float32 // [y][x]
	HumidityMap [][]float32
	TemperatureMap [][]float32
}

func NewEmptyWorld(Width int, Height int) World {
	res := World {
		"Mrak",
		Width,
		Height,
		make([][]float32, Height, Height),
		make([][]float32, Height, Height),
		make([][]float32, Height, Height),
	}

	for y := 0; y < Height; y++ {
		res.ElevationMap[y] = make([]float32, Width, Width)
		res.HumidityMap[y] = make([]float32, Width, Width)
		res.TemperatureMap[y] = make([]float32, Width, Width)
	}

	return res
}

func (w *World) Generate() {
	const (
		alpha = 2
		beta = 2
		n = 3
		seed = 100
	)
	p := NewPerlin(alpha, beta, n, seed)

	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			w.ElevationMap[y][x] = float32(p.Noise2D(float64(x), float64(y)))
		}
	}
}