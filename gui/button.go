package gui

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"image/color"
	"math"
)

type ResultChan chan interface{}
type Action func(w chan interface{})

type Button struct {
	X, Y, W, H float64
	Text string
	BgColor color.Color
	Zindex int
	Parent *RichWindow
	Callback Action
	ResultChan ResultChan
}

func (b *Button) Click(X float64, Y float64) {
	b.Callback(b.ResultChan)
}

func (b *Button) GetBoundaries() Rectangle {
	return Rectangle{
		b.X, b.X + b.W, b.Y, b.Y + b.W,
	}
}

func (b *Button) GetZindex() int {
	return b.Zindex
}

func (b *Button) Draw(w *pixelgl.Window) {
	imd := imdraw.New(nil)
	
	//draw background
	imd.Color = b.BgColor
	dx := b.Parent.X
	dy := b.Parent.Y

	imd.Push(pixel.V(b.X + dx, b.Y + dy))
	imd.Push(pixel.V(b.X + dx + b.W, b.Y + dy))
	imd.Push(pixel.V(b.X + dx + b.W, b.Y + dy + b.H))
	imd.Push(pixel.V(b.X + dx, b.Y + dy + b.H))
	imd.Polygon(0)

	// draw title
	lineHeight := atlas.LineHeight()
	basicTxt := text.New(pixel.V(math.Floor(b.X + dx + b.W / 2), math.Floor(b.Y + dy + b.H / 2 - lineHeight / 2)), atlas)
	basicTxt.Dot.X -= basicTxt.BoundsOf(b.Text).W() / 2
	fmt.Fprint(basicTxt, b.Text)

	imd.Draw(w)
	basicTxt.Draw(w, pixel.IM)
}

func (b *Button) Move(dx float64, dy float64) {
	b.X += dx
	b.Y += dy
}