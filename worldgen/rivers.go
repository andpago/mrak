package worldgen

import (
	"fmt"
	"github.com/andpago/mrak/gui"
	"math/rand"
	"time"
)

type Point struct {
	X, Y int
}

func Shuffle(vals []Point) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	// We start at the end of the slice, inserting our random
	// values one at a time.
	for n := len(vals); n > 0; n-- {
		randIndex := r.Intn(n)
		// We swap the value at index n-1 and the random index
		// to move our randomly chosen value to the end of the
		// slice, and to move the value that was at n-1 into our
		// unshuffled portion of the slice.
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
	}
}


func GenerateRivers(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	const maxRiverSurface = 1000
	const minRiverSurface = 20

	for n := 0; n < 50; n++ {
		X := rand.Intn(w.Width)
		Y := rand.Intn(w.Height)

		for w.IsWater[Y][X] {
			X = rand.Intn(w.Width)
			Y = rand.Intn(w.Height)
		}

		cancel := false

		initialWater := make([][]bool, w.Height, w.Height)
		for y := 0; y < w.Height; y++ {
			initialWater[y] = make([]bool, w.Width, w.Width)
			copy(initialWater[y], w.IsWater[y])
		}

		initialElevation := make([][]float32, w.Height, w.Height)
		for y := 0; y < w.Height; y++ {
			initialElevation[y] = make([]float32, w.Width, w.Width)
			copy(initialElevation[y], w.ElevationMap[y])
		}

		t := 0

		for {
			if w.TemperatureMap[Y][X] < 250 {
				cancel = true
				break
			}

			if initialWater[Y][X] {
				break
			}

			nb := []Point{
				{X - 1, Y},
				{X + 1, Y},
				{X, Y - 1},
				{X, Y + 1},
				{X - 1, Y - 1},
				{X + 1, Y + 1},
				{X + 1, Y - 1},
				{X - 1, Y + 1},
			}
			Shuffle(nb)
			minP := Point{X, Y}
			minElev := float32(-1)

			for _, p := range nb {
				if p.X == -1 || p.X == w.Width || p.Y == -1 || p.Y == w.Height {
					continue
				}

				if minElev == -1 {
					minElev = w.ElevationMap[p.Y][p.X]
				}

				if w.ElevationMap[minP.Y][minP.X] > w.ElevationMap[p.Y][p.X] {
					minP = p
					minElev = w.ElevationMap[p.Y][p.X]
				}
			}

			if minP.X == X && minP.Y == Y {
				if minElev == -1 {
					panic("error: minElev = -1")
				}
				w.ElevationMap[Y][X] = minElev + 1
			} else {
				w.IsWater[Y][X] = true
				t++
			}

			X, Y = minP.X, minP.Y
		}

		for y := 0; y < w.Height; y++ {
			copy(w.ElevationMap[y], initialElevation[y])
		}

		if cancel || t > maxRiverSurface || t < minRiverSurface {
			for y := 0; y < w.Height; y++ {
				copy(w.IsWater[y], initialWater[y])
			}
			n--
		} else {
			Visualize(w, buf, vis)
			fmt.Printf("river %d / %d\n", n, 50)
		}
	}
}
