package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"sync"
)

type ProtectedColorBuffer struct {
	Colors [][]color.Color
	Mu     *sync.Mutex
}


func NewProtectedColorBuffer(W int, H int, defaultColor color.Color) *ProtectedColorBuffer {
	res := &ProtectedColorBuffer{
		Colors: make([][]color.Color, H, H),
		Mu:     &sync.Mutex{},
	}

	for y := 0; y < H; y++ {
		res.Colors[y] = make([]color.Color, W, W)
		for x := 0; x < W; x++ {
			res.Colors[y][x] = defaultColor
		}
	}

	return res
}

type Canvas struct {
	X, Y float64
	W, H float64
	Colors *ProtectedColorBuffer
	Zindex int
	Parent *RichWindow
}

func NewCanvas(parent *RichWindow, W, H int) *Canvas {
	res := &Canvas {
		X:0, Y: 0, W: float64(W), H: float64(H),
		Zindex: 1,
		Parent: parent,
	}

	res.Colors = NewProtectedColorBuffer(int(res.W), int(res.H), colornames.White)

	return res
}

func (c *Canvas) GetZindex() int {
	return c.Zindex
}

func (c *Canvas) Draw(w *pixelgl.Window) {
	m := image.NewRGBA(image.Rect(0, 0, int(c.W), int(c.H)))

	dx := c.Parent.X + c.X + c.W / 2
	dy := c.Parent.Y + c.Y + c.H / 2

	c.Colors.Mu.Lock()
	for y := 0; y < int(c.H); y++ {
		for x := 0; x < int(c.W); x++ {
			m.Set(x, y, c.Colors.Colors[y][x])
		}
	}
	c.Colors.Mu.Unlock()

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