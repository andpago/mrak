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
	Move(dx float64, dy float64)
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

func (m *RichWindow) Draw(w *pixelgl.Window) {
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

func (w *RichWindow) GetTitleRectange() Rectangle {
	return Rectangle{
		w.X,
		w.X + w.W,
		w.Y + w.H - atlas.LineHeight() - 5,
		w.Y + w.H,
	}
}

type Compositor struct {
	Zbuffer map[int][]WindowID
	Windows map[WindowID]GuiWindow
	win *pixelgl.Window
}

type WindowID int
var currWid WindowID = 0

func (c *Compositor) AddWindow(w GuiWindow) WindowID {
	z := w.GetZindex()
	wid := currWid

	if _, ok := c.Zbuffer[z]; ok {
		c.Zbuffer[z] = append(c.Zbuffer[z], wid)
	} else {
		c.Zbuffer[z] = []WindowID{wid}
	}

	c.Windows[wid] = w

	currWid++
	return wid
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
		for _, wid := range c.Zbuffer[idx] {
			c.Windows[wid].Draw(c.win)
		}
	}
}

type Rectangle struct {
	X1, X2, Y1, Y2 float64
}

func (r Rectangle) Contains(vec pixel.Vec) bool {
	return (vec.X <= r.X2) && (vec.X >= r.X1) && (vec.Y <= r.Y2) && (vec.Y >= r.Y1)
}

const NULL WindowID = -1

func (c *Compositor) GetWindowTitleAt(vec pixel.Vec) WindowID {
	res := map[int]WindowID{}

	for wid, window := range c.Windows {
		if rw, ok := window.(*RichWindow); ok {
			if rw.GetTitleRectange().Contains(vec) {
				res[rw.Zindex] = wid
			}
		}
	}

	keys := make([]int, len(res), len(res))
	i := 0
	for zindex := range res {
		keys[i] = zindex
		i++
	}
	sort.Ints(keys)

	if len(keys) != 0 {
		return res[keys[len(keys) - 1]]
	}

	return NULL
}

func (r* RichWindow) Move(dx float64, dy float64) {
	r.X += dx
	r.Y += dy
}

func (c *Compositor) MoveWindow(wid WindowID, dx float64, dy float64) {
	c.Windows[wid].Move(dx, dy)
}

func NewCompositor(win *pixelgl.Window) Compositor {
	return Compositor {
		Zbuffer: map[int][]WindowID{},
		Windows: map[WindowID]GuiWindow{},
		win: win,
	}
}

type DragNDrop struct {
	LastCoords pixel.Vec
	Initiated bool
	Window WindowID
}

func (dnd *DragNDrop) Check(win *pixelgl.Window, comp *Compositor) {
	pos := win.MousePosition()
	if win.Pressed(pixelgl.MouseButtonLeft) && dnd.Initiated {
		comp.MoveWindow(dnd.Window, pos.X - dnd.LastCoords.X, pos.Y - dnd.LastCoords.Y)
		dnd.LastCoords = pos
	}

	if win.JustPressed(pixelgl.MouseButtonLeft) {
		window := comp.GetWindowTitleAt(pos)
		if window != NULL {
			dnd.Initiated = true
			dnd.Window = window
			dnd.LastCoords = win.MousePosition()
		}
	}

	if win.JustReleased(pixelgl.MouseButtonLeft) {
		dnd.Initiated = false
	}
}