package gui

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"math"
)

var atlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)

type RichWindow struct {
	BaseGuiWindow
	Title string
	FixedPosition bool
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

type Rectangle struct {
	X1, X2, Y1, Y2 float64
}

func (r Rectangle) Contains(vec pixel.Vec) bool {
	return (vec.X <= r.X2) && (vec.X >= r.X1) && (vec.Y <= r.Y2) && (vec.Y >= r.Y1)
}
