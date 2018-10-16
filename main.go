package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
	"time"
	"github.com/andpago/mrak/gui"
)


var windowConfig = pixelgl.WindowConfig{
	Title:  "Mrak",
	Bounds: pixel.R(0, 0, 800, 500),
	VSync:  true,
}

var dnd = gui.DragNDrop {
	Initiated: false,
}

func run() {

	win, err := pixelgl.NewWindow(windowConfig)
	if err != nil {
		panic(err)
	}


	comp := gui.NewCompositor(win)

	wid := comp.AddWindow(&gui.RichWindow{
		BaseGuiWindow: gui.BaseGuiWindow{
			W: windowConfig.Bounds.W() - 1,
			H: windowConfig.Bounds.H(),
			X: 1,
			Y: 0,
			Bgcolor: color.RGBA{128, 128, 128, 255},
			Bordercolor: color.RGBA{0, 0, 0, 255},
			Zindex: 1,
		},
		FixedPosition: true,
		Title: "Hello World",
	})

	fmt.Println("created window with wid =", wid)

	for !win.Closed() {
		win.Clear(colornames.Skyblue)

		dnd.Check(win, &comp)

		comp.DrawAllWindows()
		win.Update()
		time.Sleep(40 * time.Millisecond)
	}
}

func main() {
	pixelgl.Run(run)
}