package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"sync"
)

type Label struct {
	X, Y int
	Zindex int
	Parent *RichWindow
	text string
	mu *sync.Mutex
	vertAlign, horAlign TextAlign
}

func NewLabel(X, Y int, Text string, Parent *RichWindow, vertAlign, horAlign TextAlign) *Label {
	return &Label {
		X, Y,
		3,
		Parent,
		Text,
		&sync.Mutex{},
		vertAlign,
		horAlign,
	}
}

func (l *Label) Click(float64, float64) {

}

func (l *Label) SetText(text string) {
	l.mu.Lock()
	l.text = text
	l.mu.Unlock()
}

func (l *Label) GetZindex() int {
	return l.Zindex
}

func (l *Label) Draw(w *pixelgl.Window) {
	l.mu.Lock()
	defer l.mu.Unlock()

	const padding = 5

	rect := l.getBoundaries()
	x, y := rect.X1, rect.Y1
	text := DrawText(x, y, l.text, l.vertAlign, l.horAlign)
	x1, x2, y1, y2 := rect.X1, rect.X2, rect.Y1, rect.Y2
	W := int(text.Bounds().W())
	H := int(atlas.LineHeight())

	switch l.horAlign {
	case ALIGN_LEFT:
	case ALIGN_RIGHT:
		x1 -= W
		x2 -= W
	case ALIGN_CENTER:
		x1 -= W / 2
		x2 -= W / 2
	}

	switch l.vertAlign {
	case ALIGN_BOTTOM:
	case ALIGN_TOP:
		y1 -= H
		y2 -= H
	case ALIGN_CENTER:
		y1 -= H / 2
		y2 -= H / 2
	}

	x1 -= padding
	x2 += padding
	y1 -= padding
	y2 += padding

	imd := imdraw.New(nil)

	//draw background
	imd.Color = color.Black

	imd.Push(pV(x1, y1))
	imd.Push(pV(x1, y2))
	imd.Push(pV(x2, y2))
	imd.Push(pV(x2, y1))
	imd.Polygon(0)

	imd.Draw(w)
	text.Draw(w, pixel.IM)
}

func (l *Label) Move(dx int, dy int) {
	l.X += dx
	l.Y += dy
}

func (l *Label) getBoundaries() Rectangle {
	x := l.X + int(l.Parent.X)
	y := l.Y + int(l.Parent.Y)

	text := DrawText(x, y, l.text, ALIGN_TOP, ALIGN_RIGHT)

	return Rectangle{x, x + int(text.Bounds().W()), y, y + int(atlas.LineHeight())}
}

func (l *Label) GetBoundaries() Rectangle {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.getBoundaries()
}


