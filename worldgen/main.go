package worldgen

import (
	"fmt"
	"github.com/andpago/mrak/gui"
)

type World struct {
	Name string
	Width, Height int // kilometers ?
	ElevationMap [][]float32 // [y][x]
	HumidityMap [][]float32
	TemperatureMap [][]float32
	IsWater [][]bool
}

func NewEmptyWorld(Width int, Height int) World {
	res := World {
		"Mrak",
		Width,
		Height,
		make([][]float32, Height, Height),
		make([][]float32, Height, Height),
		make([][]float32, Height, Height),
		make([][]bool, Height, Height),
	}

	for y := 0; y < Height; y++ {
		res.ElevationMap[y] = make([]float32, Width, Width)
		res.HumidityMap[y] = make([]float32, Width, Width)
		res.TemperatureMap[y] = make([]float32, Width, Width)
		res.IsWater[y] = make([]bool, Width, Width)
	}

	return res
}

type Generator func(w *World, buf *gui.ProtectedColorBuffer)

func GenerateFractalWorld(w *World, buf *gui.ProtectedColorBuffer) {
	GeneratePerlinElevation(w, buf, VisualizeElevationGrayscale)
	GenerateSimpleWaterlevel(w, buf, VisualizeWaterLevel)
	GeneratePerlinLatitudeTemperature(w, buf, VisualizeTemerature)
	Visualize(w, buf, VisalizeAll)
}

func GenerateInteractive(w *World, buf *gui.ProtectedColorBuffer, generator Generator) {
	go func(){
		fmt.Println("generating world")
		defer fmt.Println("world generated")
		generator(w, buf)
	}()
}