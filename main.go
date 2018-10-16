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

func run() {

	win, err := pixelgl.NewWindow(windowConfig)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Skyblue)

	comp := NewCompositor(win)

	comp.AddWindow(MainWindow{
		BaseGuiWindow{
			W: windowConfig.Bounds.W(),
			H: windowConfig.Bounds.H() - 100,
			X: 0,
			Y: 50,
			Bgcolor: color.RGBA{128, 128, 128, 255},
			Bordercolor: color.RGBA{0, 0, 0, 255},
			Zindex: 1,
		},
	})

	comp.AddWindow(MainWindow{
		BaseGuiWindow{
			W: windowConfig.Bounds.W() - 300,
			H: windowConfig.Bounds.H(),
			X: 200,
			Y: 0,
			Bgcolor: color.RGBA{255, 255, 255, 255},
			Bordercolor: color.RGBA{0, 0, 0, 255},
			Zindex: -3,
		},
	})

	comp.AddWindow(MainWindow{
		BaseGuiWindow{
			W: windowConfig.Bounds.W() - 500,
			H: windowConfig.Bounds.H() / 10,
			X: 500,
			Y: 0,
			Bgcolor: color.RGBA{255, 0, 0, 128},
			Bordercolor: color.RGBA{0, 255, 0, 255},
			Zindex: 0,
		},
	})


	comp.DrawAllWindows()

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}