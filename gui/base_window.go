package gui

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"image/color"
	"math"
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

type TextVerticalAlign int

const (
	ALIGN_CENTER TextVerticalAlign = iota
	ALIGN_TOP TextVerticalAlign = iota
	ALIGN_BOTTOM TextVerticalAlign = iota
)

func DrawText(x, y int, txt string, align TextVerticalAlign) *text.Text {
	lineHeight := int(atlas.LineHeight())

	switch align {
	case ALIGN_CENTER:
		y -= lineHeight / 2
	case ALIGN_BOTTOM:
		y -= lineHeight
	case ALIGN_TOP:

	}

	basicTxt := text.New(pV(x, y), atlas)
	basicTxt.Dot.X -= math.Floor(basicTxt.BoundsOf(txt).W() / 2)
	fmt.Fprint(basicTxt, txt)

	return basicTxt
}