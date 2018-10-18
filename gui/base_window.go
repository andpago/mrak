package gui

import (
	"github.com/faiface/pixel/pixelgl"
	"image/color"
)

type Drawable interface {
	GetZindex() int
	Draw(w *pixelgl.Window)
	Move(dx int, dy int)
	GetBoundaries() Rectangle
}

type BaseGuiWindow struct {
	X, Y, W, H int
	Bgcolor color.Color
	Bordercolor color.Color
	Zindex int
}

func (b BaseGuiWindow) GetZindex() int {
	return b.Zindex
}
