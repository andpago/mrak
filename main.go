package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
)


var windowConfig = pixelgl.WindowConfig{
	Title:  "Mrak",
	Bounds: pixel.R(0, 0, 800, 500),
	VSync:  true,
}

var w = MainWindow{
	BaseGuiWindow{
		W: windowConfig.Bounds.W(),
		H: windowConfig.Bounds.H(),
		X: 0,
		Y: 0,
		Bgcolor: color.RGBA{128, 128, 128, 255},
		Bordercolor: color.RGBA{0, 0, 0, 255},
	},
}

func run() {
	win, err := pixelgl.NewWindow(windowConfig)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Skyblue)
	w.Draw(win)

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}