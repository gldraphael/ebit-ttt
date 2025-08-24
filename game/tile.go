package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type tile struct {
	// Define tile properties here
	X     float64
	Y     float64
	Image *ebiten.Image
}

func (t *tile) Draw(sc *ebiten.Image) {
	// Implement drawing logic here
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.X, t.Y)
	sc.DrawImage(t.Image, op)
}
