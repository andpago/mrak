package worldgen

import (
	"github.com/andpago/mrak/gui"
	"math"
)

func GenerateSimpleWaterlevel(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	const waterlevel = 0
	hasWater := false

	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			hasWater = w.ElevationMap[y][x] < waterlevel
			w.IsWater[y][x] = hasWater
			w.IsSea[y][x] = hasWater
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
							if w.IsSea[Y][X] {
								mult /= 5
							}

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

func CalculateDistanceToWater(w *World, buf *gui.ProtectedColorBuffer, vis Visualizer) {
	visited := map[Point]bool{}
	nbrs := NewQueue(w.Width * w.Height)

	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			if w.IsWater[y][x] {
				w.DistanceToWater[y][x] = 0
				visited[Point{x, y}] = true
			} else {
				w.DistanceToWater[y][x] = math.MaxInt32
			}
		}
	}

	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			if w.IsWater[y][x] {
				continue
			}

			for _, p := range []Point{{x-1,y}, {x+1,y}, {x,y+1}, {x,y-1}} {
				if p.X == -1 || p.X == w.Width || p.Y == -1 || p.Y == w.Height {
					continue
				}

				if w.IsWater[p.Y][p.X] {
					visited[p] = true
					nbrs.Push(QueueValueType{p, 1})
					break
				}
			}
		}
	}

	for !nbrs.IsEmpty() {
		p := nbrs.Pop()
		for _, nb := range []Point{{p.X-1,p.Y}, {p.X+1,p.Y}, {p.X,p.Y+1}, {p.X,p.Y-1}} {
			if nb.X == -1 || nb.X == w.Width || nb.Y == -1 || nb.Y == w.Height {
				continue
			}

			if _, vis := visited[nb]; vis {
				continue
			}

			visited[nb] = true
			w.DistanceToWater[nb.Y][nb.X] = p.Distance + 1
			nbrs.Push(QueueValueType{nb, p.Distance + 1})
		}
	}

	Visualize(w, buf, vis)
}