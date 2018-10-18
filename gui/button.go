package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
)

type ResultChan chan interface{}
type Action func(w chan interface{})

type Button struct {
	X, Y, W, H int
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

func pV(x, y int) pixel.Vec {
	return pixel.V(float64(x), float64(y))
}

func (b *Button) Draw(w *pixelgl.Window) {
	imd := imdraw.New(nil)
	
	//draw background
	imd.Color = b.BgColor
	dx := b.Parent.X
	dy := b.Parent.Y

	imd.Push(pV(b.X + dx, b.Y + dy))
	imd.Push(pV(b.X + dx + b.W, b.Y + dy))
	imd.Push(pV(b.X + dx + b.W, b.Y + dy + b.H))
	imd.Push(pV(b.X + dx, b.Y + dy + b.H))
	imd.Polygon(0)

	// draw title
	basicTxt := DrawText(b.X + dx + b.W / 2, b.Y + dy + b.H / 2, b.Text, ALIGN_CENTER)

	imd.Draw(w)
	basicTxt.Draw(w, pixel.IM)
}

func (b *Button) Move(dx int, dy int) {
	b.X += dx
	b.Y += dy
}