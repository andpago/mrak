package gui_menus

import (
	"github.com/andpago/mrak/gui"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func CreateWorldGenMenu(config *pixelgl.WindowConfig, comp *gui.Compositor, switchWindowChannel chan interface{}) (mwin *gui.RichWindow, wid gui.WindowID) {
	mwin = &gui.RichWindow{
		BaseGuiWindow: gui.BaseGuiWindow{
			W:           config.Bounds.W() - 1,
			H:           config.Bounds.H(),
			X:           1,
			Y:           0,
			Bgcolor:     colornames.White,
			Bordercolor: colornames.Black,
			Zindex:      2,
		},
		FixedPosition: false,
		Title:         "World Generation",
		Children:      []gui.Clickable{},
	}

	mwin.Children = append(mwin.Children, &gui.Button{
		0, 0, 100, 30,
		"back",
		colornames.Gray,
		2,
		mwin,
		func(w chan interface{}) {
			w <- 0
		},
		switchWindowChannel,
	})

	mwin.Children = append(mwin.Children, gui.NewCanvas(mwin))

	return mwin, comp.AddWindow(mwin)
}
