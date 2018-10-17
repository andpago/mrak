package gui_menus

import (
	"github.com/andpago/mrak/gui"
	"github.com/andpago/mrak/worldgen"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var world = worldgen.NewEmptyWorld(3000, 3000)

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
	canvas := gui.NewCanvas(mwin, int(config.Bounds.W()), int(config.Bounds.H()) - 20)

	mwin.Children = append(mwin.Children, &gui.Button{
		110, 0, 100, 30,
		"Generate",
		colornames.Gray,
		2,
		mwin,
		func(w chan interface{}) {
			worldgen.GenerateInteractive(&world, canvas.Colors, worldgen.GenerateFractalWorld)
		},
		switchWindowChannel,
	})

	mwin.Children = append(mwin.Children, canvas)

	return mwin, comp.AddWindow(mwin)
}
