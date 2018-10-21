package worldgen

import (
	"fmt"
)

type Continent struct {
	Name string
	Points []Point
}

func NewContinent() Continent {
	res := Continent{
		"New Continent",
		[]Point{},
	}

	return res
}

func (w *World) FindContinents() {
	w.Continets = []Continent{}
	fmt.Println("finding continents")
	defer func(){fmt.Println(len(w.Continets), "continents found")}()

	colors := make([][]int, w.Height, w.Height)
	for y := 0; y < w.Height; y++ {
		colors[y] = make([]int, w.Width, w.Width)
		for x := 0; x < w.Width; x++ {
			colors[y][x] = -1
		}
	}

	step := w.Width / 10

	for y := 0; y < w.Height; y += step{
		for x := 0; x < w.Width; x += step{
			if colors[y][x] != -1 || w.IsSea[y][x] {
				continue
			}
			w.Continets = append(w.Continets, NewContinent())

			q := NewQueue(w.Width * w.Height)
			q.Push(QueueValueType{Point{x, y}, 0})
			for !q.IsEmpty() {
				p := q.Pop()

				X, Y := p.X, p.Y

				for _, nb := range []Point{{X-1,Y}, {X+1, Y}, {X, Y+1}, {X, Y-1}} {
					if nb.X == -1 || nb.X == w.Width || nb.Y == -1 || nb.Y == w.Height {
						continue
					}

					if colors[nb.Y][nb.X] != -1 || w.IsSea[nb.Y][nb.X]  {
						continue
					}

					colors[nb.Y][nb.X] = len(w.Continets)
					w.Continets[len(w.Continets) - 1].Points = append(w.Continets[len(w.Continets) - 1].Points, nb)
					q.Push(QueueValueType{nb, p.Distance + 1})
				}
			}
		}
	}
}