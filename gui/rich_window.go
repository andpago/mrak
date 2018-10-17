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
	Children []Clickable
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
	titleArea := m.GetTitleRectangle()
	imd.Push(pixel.V(titleArea.X1, titleArea.Y1))
	imd.Push(pixel.V(titleArea.X1, titleArea.Y2))
	imd.Push(pixel.V(titleArea.X2, titleArea.Y2))
	imd.Push(pixel.V(titleArea.X2, titleArea.Y1))
	imd.Polygon(0)


	imd.Draw(w)
	basicTxt.Draw(w, pixel.IM)

	//fmt.Println("window has", len(m.Children), "children")
	for _, child := range m.Children {
		child.Draw(w)
	}
}

type Clickable interface {
	Drawable
	Click(X float64, Y float64)
}

func (w *RichWindow) GetChildAt(vec pixel.Vec) Clickable {
	for _, child := range w.Children {
		if child.GetBoundaries().Contains(vec) {
			return child
		}
	}

	return nil
}

func (w *RichWindow) GetTitleRectangle() Rectangle {
	return Rectangle{
		w.X,
		w.X + w.W,
		w.Y + w.H - atlas.LineHeight() - 5,
		w.Y + w.H,
	}
}

func (w *RichWindow) GetBoundaries() Rectangle {
	return Rectangle{
		w.X,
		w.X + w.W,
		w.Y,
		w.Y + w.H,
	}
}

type Rectangle struct {
	X1, X2, Y1, Y2 float64
}

func (r Rectangle) Contains(vec pixel.Vec) bool {
	return (vec.X <= r.X2) && (vec.X >= r.X1) && (vec.Y <= r.Y2) && (vec.Y >= r.Y1)
}
