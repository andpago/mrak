package worldgen

import (
	"fmt"
	"github.com/andpago/mrak/gui"
)

type World struct {
	Name string
	Width, Height int // kilometers ?
	ElevationMap [][]float32 // [y][x] meters over the water
	HumidityMap [][]float32
	TemperatureMap [][]float32
	IsWater [][]bool
	WaterAdjacency [][]int
	MaxWad int
	DistanceToWater [][]int
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
		make([][]int, Height, Height),
		0,
		make([][]int, Height, Height),
	}

	for y := 0; y < Height; y++ {
		res.ElevationMap[y] = make([]float32, Width, Width)
		res.HumidityMap[y] = make([]float32, Width, Width)
		res.TemperatureMap[y] = make([]float32, Width, Width)
		res.IsWater[y] = make([]bool, Width, Width)
		res.WaterAdjacency[y] = make([]int, Width, Width)
		res.DistanceToWater[y] = make([]int, Width, Width)
	}

	return res
}

type Generator func(w *World, buf *gui.ProtectedColorBuffer)

func GenerateFractalWorld(w *World, buf *gui.ProtectedColorBuffer, upd func(msg string)) {
	upd("generating perlin elevation")
	GeneratePerlinElevation(w, buf, VisualizeElevationGrayscale)
	upd("generating water lever")
	GenerateSimpleWaterlevel(w, buf, VisualizeWaterLevel)
	upd("generating temperature")
	GeneratePerlinLatitudeTemperature(w, buf, VisualizeTemerature)
	upd("running rivers")
	GenerateRivers(w, buf, VisualizeClimate)
	upd("calculating water adjacency")
	CalculateWaterAdjacency(w, buf, VisualizeWaterAdjacency)
	upd("generating distance to warer")
	CalculateDistanceToWater(w, buf, VisualizeDistanceToWater)
	upd("generating humidity")
	GeneratePerlinWadHumidity(w, buf, VisualizeHumidity)
	Visualize(w, buf, VisualizeClimate)
	upd("done!")
}

func GenerateInteractive(w *World, buf *gui.ProtectedColorBuffer, upd func(msg string)) {
	go func(){
		fmt.Println("generating world")
		defer fmt.Println("world generated")
		GenerateFractalWorld(w, buf, upd)
	}()
}