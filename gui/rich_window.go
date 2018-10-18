package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"sort"
)

var atlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)

type RichWindow struct {
	BaseGuiWindow
	Title string
	FixedPosition bool
	Children []Clickable
}

func (w *RichWindow) Move(dx, dy int) {
	w.X += dx
	w.Y += dy
}

func (m *RichWindow) Draw(w *pixelgl.Window) {
	imd := imdraw.New(nil)

	//draw background
	imd.Color = m.Bgcolor
	imd.Push(pV(m.X, m.Y))
	imd.Push(pV(m.X + m.W, m.Y))
	imd.Push(pV(m.X + m.W, m.Y + m.H))
	imd.Push(pV(m.X, m.Y + m.H))
	imd.Polygon(0)

	// draw border
	imd.Color = m.Bordercolor
	imd.Push(pV(m.X, m.Y))
	imd.Push(pV(m.X + m.W, m.Y))
	imd.Push(pV(m.X + m.W, m.Y + m.H))
	imd.Push(pV(m.X, m.Y + m.H))
	imd.Polygon(1)

	// draw title
	basicTxt := DrawText(m.X + m.W / 2, m.Y + m.H, m.Title, ALIGN_BOTTOM)

	// title underline
	imd.Color = m.Bordercolor
	titleArea := m.GetTitleRectangle()
	imd.Push(pV(titleArea.X1, titleArea.Y1))
	imd.Push(pV(titleArea.X1, titleArea.Y2))
	imd.Push(pV(titleArea.X2, titleArea.Y2))
	imd.Push(pV(titleArea.X2, titleArea.Y1))
	imd.Polygon(0)


	imd.Draw(w)
	basicTxt.Draw(w, pixel.IM)

	// draw children layer by layer
	zToInd := map[int][]int{}
	for i, child := range m.Children {
		idx := child.GetZindex()

		if _, ok := zToInd[idx]; ok {
			zToInd[idx] = append(zToInd[idx], i)
		} else {
			zToInd[idx] = []int{i}
		}
	}
	zIndices := make([]int, len(zToInd), len(zToInd))
	i := 0
	for z := range zToInd {
		zIndices[i] = z
		i++
	}
	sort.Ints(zIndices)

	for _, z := range zIndices {
		for _, idx := range zToInd[z] {
			m.Children[idx].Draw(w)
		}
	}
}

type Clickable interface {
	Drawable
	Click(X float64, Y float64)
}

func (w *RichWindow) GetChildAt(vec pixel.Vec) Clickable {
	for _, child := range w.Children {
		if child.GetBoundaries().Contains(vec) {
			return child
		}
	}

	return nil
}

func (w *RichWindow) GetTitleRectangle() Rectangle {
	return Rectangle{
		w.X,
		w.X + w.W,
		w.Y + w.H - int(atlas.LineHeight()) - 5,
		w.Y + w.H,
	}
}

func (w *RichWindow) GetBoundaries() Rectangle {
	return Rectangle{
		w.X,
		w.X + w.W,
		w.Y,
		w.Y + w.H,
	}
}

type Rectangle struct {
	X1, X2, Y1, Y2 int
}

func (r Rectangle) Contains(vec pixel.Vec) bool {
	return (vec.X <= float64(r.X2)) && (vec.X >= float64(r.X1)) && (vec.Y <= float64(r.Y2)) && (vec.Y >= float64(r.Y1))
}
