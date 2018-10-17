package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
)

type Canvas struct {
	X, Y float64
	W, H float64
	Colors [][]color.Color
	Zindex int
	Parent *RichWindow
}

func NewCanvas(parent *RichWindow) *Canvas {
	res := &Canvas {
		0, 0, 300, 300,
		make([][]color.Color, 300, 300),
		1,
		parent,
	}
	for y := 0; y < 300; y++ {
		res.Colors[y] = make([]color.Color, 300, 300)
		for x := 0; x < 300; x++ {
			res.Colors[y][x] = colornames.Blue
		}
	}

	return res
}

func (c *Canvas) GetZindex() int {
	return c.Zindex
}

func (c *Canvas) Draw(w *pixelgl.Window) {
	m := image.NewRGBA(image.Rect(0, 0, int(c.W), int(c.H)))

	dx := c.Parent.X + c.X + c.W / 2
	dy := c.Parent.Y + c.Y + c.H / 2

	for y := 0; y < int(c.H); y++ {
		for x := 0; x < int(c.W); x++ {
			m.Set(x, y, c.Colors[y][x])
		}
	}

	p := pixel.PictureDataFromImage(m)
	pixel.NewSprite(p, p.Bounds()).Draw(w, pixel.IM.Moved(pixel.Vec{dx, dy}))

}

func (c *Canvas) Move(dx float64, dy float64) {
	c.X += dx
	c.Y += dy
}

func (c *Canvas) GetBoundaries() Rectangle {
	return Rectangle{c.X, c.X + c.W, c.Y, c.Y + c.H}
}

func (c *Canvas) Click(float64, float64) {

}