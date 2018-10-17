package gui_menus

import (
	"fmt"
	"github.com/andpago/mrak/gui"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"os"
)

func CreateMainMenu(config *pixelgl.WindowConfig, comp *gui.Compositor) (mwin *gui.RichWindow, wid gui.WindowID) {
	mwin = &gui.RichWindow{
		BaseGuiWindow: gui.BaseGuiWindow{
			W: config.Bounds.W() - 1,
			H: config.Bounds.H(),
			X: 1,
			Y: 0,
			Bgcolor: colornames.Gray,
			Bordercolor: colornames.Black,
			Zindex: 1,
		},
		FixedPosition: true,
		Title: "Main Menu",
		Children: []gui.Clickable{},
	}
	wid = comp.AddWindow(mwin)

	mwin.Children = append(mwin.Children, &gui.Button{
		350, 250, 100, 30,
		"Continue",
		colornames.Black,
		10,
		mwin,
		func() {
			fmt.Println("Not implemented")
		},
	})

	mwin.Children = append(mwin.Children, &gui.Button{
		350, 200, 100, 30,
		"New game",
		colornames.Black,
		10,
		mwin,
		func() {
			fmt.Println("Not implemented")
		},
	})

	mwin.Children = append(mwin.Children, &gui.Button{
		350, 150, 100, 30,
		"Exit",
		colornames.Black,
		10,
		mwin,
		func() {
			os.Exit(0)
		},
	})

	return mwin, wid
}
