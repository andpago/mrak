package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)


type DragNDrop struct {
	LastCoords pixel.Vec
	Initiated bool
	Window *RichWindow
}

func (dnd *DragNDrop) Check(win *pixelgl.Window, comp *Compositor) {
	pos := win.MousePosition()
	if win.Pressed(pixelgl.MouseButtonLeft) && dnd.Initiated {
		dnd.Window.Move(int(pos.X - dnd.LastCoords.X), int(pos.Y - dnd.LastCoords.Y))
		dnd.LastCoords = pos
	}

	if win.JustPressed(pixelgl.MouseButtonLeft) {
		window := comp.GetWindowTitleAt(pos)
		if window != nil && !window.FixedPosition {
			dnd.Initiated = true
			dnd.Window = window
			dnd.LastCoords = win.MousePosition()
		}
	}

	if win.JustReleased(pixelgl.MouseButtonLeft) {
		dnd.Initiated = false
	}
}
