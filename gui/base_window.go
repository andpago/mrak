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

type TextAlign int

const (
	ALIGN_CENTER TextAlign = iota
	ALIGN_TOP    TextAlign = iota
	ALIGN_BOTTOM TextAlign = iota
	ALIGN_LEFT   TextAlign = iota
	ALIGN_RIGHT  TextAlign = iota
)

func DrawText(x, y int, txt string, verticalAlign TextAlign, horizontalAlign TextAlign) *text.Text {
	lineHeight := int(atlas.LineHeight())

	switch verticalAlign {
	case ALIGN_CENTER:
		y -= lineHeight / 2
	case ALIGN_BOTTOM:
		y -= lineHeight
	case ALIGN_TOP:

	}

	basicTxt := text.New(pV(x, y), atlas)
	switch horizontalAlign {
	case ALIGN_LEFT:
	case ALIGN_CENTER:
		basicTxt.Dot.X -= math.Floor(basicTxt.BoundsOf(txt).W() / 2)
	case ALIGN_RIGHT:
		basicTxt.Dot.X -= math.Floor(basicTxt.BoundsOf(txt).W())

	}

	fmt.Fprint(basicTxt, txt)

	return basicTxt
}