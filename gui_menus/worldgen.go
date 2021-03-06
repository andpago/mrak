package gui_menus

import (
	"github.com/andpago/mrak/gui"
	"github.com/andpago/mrak/worldgen"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func exclusively(c *gui.Compositor, f func(chan interface{})) func(chan interface{}) {
	return func(ch chan interface{}) {
		c.LockAllButtons()
		f(ch)
		c.UnlockAllButtons()
	}
}

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

	label := gui.NewLabel(mwin.W - 1, mwin.H - 20, "Hello label", mwin, gui.ALIGN_CENTER, gui.ALIGN_RIGHT)

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
		exclusively(comp, func(w chan interface{}) {
			worldgen.GenerateInteractive(&worldgen.TheWorld, canvas.Colors, func(msg string){label.SetText(msg)})
		}),
		switchWindowChannel,
	})

	mwin.Children = append(mwin.Children, &gui.Button{
		0, mwin.H - 50, 120, 30,
		"Elevation map",
		colornames.Gray,
		2,
		mwin,
		func(w chan interface{}) {
			worldgen.Visualize(&worldgen.TheWorld, canvas.Colors, worldgen.VisualizeElevationGrayscale)
		},
		switchWindowChannel,
	})

	mwin.Children = append(mwin.Children, &gui.Button{
		0, mwin.H - 85, 120, 30,
		"Temperature map",
		colornames.Gray,
		2,
		mwin,
		func(w chan interface{}) {
			worldgen.Visualize(&worldgen.TheWorld, canvas.Colors, worldgen.VisualizeTemerature)
		},
		switchWindowChannel,
	})


	mwin.Children = append(mwin.Children, &gui.Button{
		0, mwin.H - 120, 120, 30,
		"Humidity map",
		colornames.Gray,
		2,
		mwin,
		func(w chan interface{}) {
			worldgen.Visualize(&worldgen.TheWorld, canvas.Colors, worldgen.VisualizeHumidity)
		},
		switchWindowChannel,
	})

	mwin.Children = append(mwin.Children, &gui.Button{
		0, mwin.H - 155, 120, 30,
		"Climate + continents",
		colornames.Gray,
		2,
		mwin,
		func(w chan interface{}) {
			worldgen.Visualize(&worldgen.TheWorld, canvas.Colors, worldgen.VisualizeContinents)
		},
		switchWindowChannel,
	})

	mwin.Children = append(mwin.Children, &gui.Button{
		0, mwin.H - 190, 120, 30,
		"Final map",
		colornames.Gray,
		2,
		mwin,
		func(w chan interface{}) {
			worldgen.Visualize(&worldgen.TheWorld, canvas.Colors, worldgen.VisualizeAll)
		},
		switchWindowChannel,
	})


	mwin.Children = append(mwin.Children, &gui.Button{
		0, mwin.H - 225, 120, 30,
		"Save world",
		colornames.Gray,
		2,
		mwin,
		exclusively(comp, func(w chan interface{}) {
			worldgen.TheWorld.Save("world.bin")
		}),
		switchWindowChannel,
	})

	mwin.Children = append(mwin.Children, &gui.Button{
		0, mwin.H - 260, 120, 30,
		"Load world",
		colornames.Gray,
		2,
		mwin,
		exclusively(comp, func(w chan interface{}) {
			worldgen.TheWorld.Load("world.bin")
			worldgen.Visualize(&worldgen.TheWorld, canvas.Colors, worldgen.VisualizeClimate)
		}),
		switchWindowChannel,
	})




	mwin.Children = append(mwin.Children, label)


	mwin.Children = append(mwin.Children, canvas)

	return mwin, comp.AddWindow(mwin)
}
