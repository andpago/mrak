package gui

import (
	"github.com/faiface/pixel/pixelgl"
	"image/color"
)

type Drawable interface {
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
