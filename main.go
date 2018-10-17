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

const (
	MODE_MAIN_MENU = iota
	MODE_WORLDGEN = iota
)

var modeChange = make(chan interface{}, 1)

func changeMode(comp *gui.Compositor, mode int) {
	fmt.Println("game mode changed to", mode)
	comp.DestroyAllWindows()
	switch mode {
	case MODE_MAIN_MENU:
		gui_menus.CreateMainMenu(&windowConfig, comp, modeChange)
	case MODE_WORLDGEN:
		gui_menus.CreateWorldGenMenu(&windowConfig, comp, modeChange)
	}
}

func run() {

	win, err := pixelgl.NewWindow(windowConfig)
	if err != nil {
		panic(err)
	}


	comp := gui.NewCompositor(win)
	changeMode(&comp, MODE_MAIN_MENU)

	for !win.Closed() {
		win.Clear(colornames.Skyblue)

		comp.Dnd.Check(win, &comp)
		comp.CheckButtons()

		comp.DrawAllWindows()
		win.Update()

		select {
			case modeRead := <-modeChange:
				if newMode, ok := modeRead.(int); ok {
					changeMode(&comp, newMode)
				} else {
					fmt.Println("fail")
				}
			default:

		}

		time.Sleep(40 * time.Millisecond)
	}
}

func main() {
	pixelgl.Run(run)
}