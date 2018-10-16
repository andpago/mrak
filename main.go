package main

import (
	"fmt"
	"github.com/andpago/mrak/gui"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
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


	comp := gui.NewCompositor(win)

	mwin := &gui.RichWindow{
		BaseGuiWindow: gui.BaseGuiWindow{
			W: windowConfig.Bounds.W() - 1,
			H: windowConfig.Bounds.H(),
			X: 1,
			Y: 0,
			Bgcolor: colornames.Gray,
			Bordercolor: colornames.Black,
			Zindex: 1,
		},
		FixedPosition: false,
		Title: "Hello World",
		Children: []gui.Clickable{},
	}
	wid := comp.AddWindow(mwin)

	mwin.Children = append(mwin.Children, &gui.Button{
		50, 50, 100, 30,
		"Hello btn",
		colornames.Black,
		10,
		mwin,
		func() {
			fmt.Println("I am clicked")
		},
	})


	fmt.Println("created window with wid =", wid)

	for !win.Closed() {
		win.Clear(colornames.Skyblue)

		comp.Dnd.Check(win, &comp)
		comp.CheckButtons()

		comp.DrawAllWindows()
		win.Update()
		time.Sleep(40 * time.Millisecond)
	}
}

func main() {
	pixelgl.Run(run)
}