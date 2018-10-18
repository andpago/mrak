package gui_menus

import (
	"github.com/andpago/mrak/gui"
	"github.com/andpago/mrak/worldgen"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var world = worldgen.NewEmptyWorld(4000, 4000)

func CreateWorldGenMenu(config *pixelgl.WindowConfig, comp *gui.Compositor, switchWindowChannel chan interface{}) (mwin *gui.RichWindow, wid gui.WindowID) {
	mwin = &gui.RichWindow{
		BaseGuiWindow: gui.BaseGuiWindow{
			W:           int(config.Bounds.W()) - 1,
			H:           int(config.Bounds.H()),
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

	h := int(config.Bounds.H()) - int(mwin.GetTitleRectangle().Y2 - mwin.GetTitleRectangle().Y1)
	canvas := gui.NewCanvas(mwin, h, h)
	canvas.X = int(config.Bounds.W()) - canvas.W

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

	mwin.Children = append(mwin.Children, &gui.Button{
		0, mwin.H - 50, 100, 30,
		"Elevation map",
		colornames.Gray,
		2,
		mwin,
		func(w chan interface{}) {
			worldgen.Visualize(&world, canvas.Colors, worldgen.VisualizeElevationGrayscale)
		},
		switchWindowChannel,
	})

	mwin.Children = append(mwin.Children, &gui.Button{
		0, mwin.H - 85, 100, 30,
		"Temperature map",
		colornames.Gray,
		2,
		mwin,
		func(w chan interface{}) {
			worldgen.Visualize(&world, canvas.Colors, worldgen.VisualizeTemerature)
		},
		switchWindowChannel,
	})

	mwin.Children = append(mwin.Children, &gui.Button{
		0, mwin.H - 120, 100, 30,
		"Final map",
		colornames.Gray,
		2,
		mwin,
		func(w chan interface{}) {
			worldgen.Visualize(&world, canvas.Colors, worldgen.VisualizeAll)
		},
		switchWindowChannel,
	})

	mwin.Children = append(mwin.Children, canvas)

	return mwin, comp.AddWindow(mwin)
}
