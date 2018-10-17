package main

import (
	"fmt"
	"github.com/andpago/mrak/gui"
	"github.com/andpago/mrak/gui_menus"
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

	_, wid := gui_menus.CreateMainMenu(&windowConfig, &comp)

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