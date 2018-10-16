package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"image/color"
	"math"
	"sort"
)

type GuiWindow interface {
	GetZindex() int
	Draw(w *pixelgl.Window)
}

type BaseGuiWindow struct {
	X, Y, W, H float64
	Bgcolor color.Color
	Bordercolor color.Color
	Zindex int
}

func (b BaseGuiWindow) GetZindex() int {
	return b.Zindex
}

type RichWindow struct {
	BaseGuiWindow
	Title string
}

func (m RichWindow) Draw(w *pixelgl.Window) {
	imd := imdraw.New(nil)

	//draw background
	imd.Color = m.Bgcolor
	imd.Push(pixel.V(m.X, m.Y))
	imd.Push(pixel.V(m.X + m.W, m.Y))
	imd.Push(pixel.V(m.X + m.W, m.Y + m.H))
	imd.Push(pixel.V(m.X, m.Y + m.H))
	imd.Polygon(0)

	// draw border
	imd.Color = m.Bordercolor
	imd.Push(pixel.V(m.X, m.Y))
	imd.Push(pixel.V(m.X + m.W, m.Y))
	imd.Push(pixel.V(m.X + m.W, m.Y + m.H))
	imd.Push(pixel.V(m.X, m.Y + m.H))
	imd.Polygon(1)

	// draw title
	lineHeight := atlas.LineHeight()
	basicTxt := text.New(pixel.V(math.Floor(m.X + m.W / 2), m.Y + m.H - lineHeight), atlas)
	basicTxt.Dot.X -= basicTxt.BoundsOf(m.Title).W() / 2
	fmt.Fprint(basicTxt, m.Title)

	// title underline
	imd.Color = m.Bordercolor
	imd.Push(pixel.V(m.X, m.Y + m.H - lineHeight - 5))
	imd.Push(pixel.V(m.X + m.W, m.Y + m.H - lineHeight - 5))
	imd.Line(1)


	imd.Draw(w)
	basicTxt.Draw(w, pixel.IM)
}

type Compositor struct {
	Zbuffer map[int][]GuiWindow
	win *pixelgl.Window
}

func (c *Compositor) AddWindow(w GuiWindow) {
	z := w.GetZindex()

	if _, ok := c.Zbuffer[z]; ok {
		c.Zbuffer[z] = append(c.Zbuffer[z], w)
	} else {
		c.Zbuffer[z] = []GuiWindow{w}
	}
}

func (c *Compositor) DrawAllWindows() {
	zindices := make([]int, len(c.Zbuffer), len(c.Zbuffer))

	i := 0
	for idx := range c.Zbuffer {
		zindices[i] = idx
		i++
	}

	sort.Ints(zindices)

	for _, idx := range zindices {
		for _, window := range c.Zbuffer[idx] {
			window.Draw(c.win)
		}
	}
}

func NewCompositor(win *pixelgl.Window) Compositor {
	return Compositor {
		Zbuffer: map[int][]GuiWindow{},
		win: win,
	}
}