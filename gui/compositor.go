package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"sort"
)

type Compositor struct {
	Zbuffer map[int][]WindowID
	Windows map[WindowID]Drawable
	win *pixelgl.Window
}

type WindowID int
var currWid WindowID = 0

func (c *Compositor) AddWindow(w Drawable) WindowID {
	z := w.GetZindex()
	wid := currWid

	if _, ok := c.Zbuffer[z]; ok {
		c.Zbuffer[z] = append(c.Zbuffer[z], wid)
	} else {
		c.Zbuffer[z] = []WindowID{wid}
	}

	c.Windows[wid] = w

	currWid++
	return wid
}

func (c *Compositor) DrawAllWindows() {
	zindices := make([]int, len(c.Zbuffer), len(c.Zbuffer))

	i := 0
	for idx := range c.Zbuffer {
		zindices[i] = idx
		i++
	}

	sort.Ints(zindices)

	for _, idx := range zindices {
		for _, wid := range c.Zbuffer[idx] {
			c.Windows[wid].Draw(c.win)
		}
	}
}

func (c *Compositor) GetWindowTitleAt(vec pixel.Vec) *RichWindow {
	res := map[int]*RichWindow{}

	for _, window := range c.Windows {
		if rw, ok := window.(*RichWindow); ok {
			if rw.GetTitleRectange().Contains(vec) {
				res[rw.Zindex] = rw
			}
		}
	}

	keys := make([]int, len(res), len(res))
	i := 0
	for zindex := range res {
		keys[i] = zindex
		i++
	}
	sort.Ints(keys)

	if len(keys) != 0 {
		return res[keys[len(keys) - 1]]
	}

	return nil
}


func NewCompositor(win *pixelgl.Window) Compositor {
	return Compositor {
		Zbuffer: map[int][]WindowID{},
		Windows: map[WindowID]Drawable{},
		win: win,
	}
}