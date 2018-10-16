package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"time"
)


var windowConfig = pixelgl.WindowConfig{
	Title:  "Mrak",
	Bounds: pixel.R(0, 0, 800, 500),
	VSync:  true,
}

var atlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)

func run() {

	win, err := pixelgl.NewWindow(windowConfig)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Skyblue)

	comp := NewCompositor(win)

	comp.AddWindow(RichWindow{
		BaseGuiWindow: BaseGuiWindow{
			W: windowConfig.Bounds.W() - 11,
			H: windowConfig.Bounds.H() - 100,
			X: 1,
			Y: 50,
			Bgcolor: color.RGBA{128, 128, 128, 255},
			Bordercolor: color.RGBA{255, 0, 0, 255},
			Zindex: 1,
		},
		Title: "Hello World",
	})



	for !win.Closed() {
		comp.DrawAllWindows()
		win.Update()
		time.Sleep(40 * time.Millisecond)
	}
}

func main() {
	pixelgl.Run(run)
}