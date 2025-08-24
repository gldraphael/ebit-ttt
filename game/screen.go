package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type screen struct {
	// Define screen properties here
	board []tile
}

func (s *screen) Update() error {
	// Implement screen update logic here
	return nil
}

func (s *screen) Draw(screen *ebiten.Image) {
	// Implement screen drawing logic here
	for _, tile := range s.board {
		tile.Draw(screen)
	}
}

func (s *screen) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 300, 300
}

func newScreen() *screen {
	g := &screen{}
	for i := range 3 {
		for j := range 3 {
			g.board = append(g.board, tile{
				X:     float64(i * 50),
				Y:     float64(j * 50),
				Image: ebiten.NewImage(50, 50),
			})
		}
	}
	return g
}
