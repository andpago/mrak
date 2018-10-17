package worldgen

import (
	"fmt"
	"github.com/andpago/mrak/gui"
	"image/color"
	"time"
)

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

type Generator func(w *World, buf *gui.ProtectedColorBuffer)

func GenerateFractalWorld(w *World, buf *gui.ProtectedColorBuffer) {
	const chunkNum = 5
	chunkSizeX := w.Width / chunkNum
	chunkSizeY := w.Height / chunkNum

	const (
		alpha = 2
		beta = 2
		n = 3
	)

	seed := int64(time.Now().Nanosecond())

	p := NewPerlin(alpha, beta, n, seed)

	for yChunk := 0; yChunk < chunkNum; yChunk++ {
		for xChunk := 0; xChunk < chunkNum; xChunk++ {
			for y := yChunk * chunkSizeY; y < (yChunk + 1) * chunkSizeY; y++ {
				for x := xChunk * chunkSizeX; x < (xChunk + 1) * chunkSizeX; x++ {
					w.ElevationMap[y][x] = 50 + 100 * float32(p.Noise2D(float64(x * 5) / float64(w.Width), float64(y * 5) / float64(w.Height)))
				}
			}
			buf.Mu.Lock()
			for y := 0; y < len(buf.Colors); y++ {
				for x := 0; x < len(buf.Colors[0]); x++ {
					gray := uint8(w.ElevationMap[y * w.Height / len(buf.Colors)][x * w.Width / len(buf.Colors[0])])
					buf.Colors[y][x] = color.Gray{gray}
				}
			}
			buf.Mu.Unlock()
			//fmt.Printf("chunk %d done\n", yChunk * chunkNum + xChunk)
		}
	}
}

func GenerateInteractive(w *World, buf *gui.ProtectedColorBuffer, generator Generator) {
	go func(){
		fmt.Println("generating world")
		defer fmt.Println("world generated")
		generator(w, buf)
	}()
}