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

func NewTile() *tile {
	return &tile{}
}

func (t *tile) Draw(screen *ebiten.Image) {
	// Implement drawing logic here
}

func (t *tile) Update() {
	// Implement update logic here
}
