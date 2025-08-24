package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Screen struct {
	// Define screen properties here
	board []Tile
}

func (s *Screen) Update() error {
	// Implement screen update logic here
	return nil
}

func (s *Screen) Draw(screen *ebiten.Image) {
	// Implement screen drawing logic here
	for _, tile := range s.board {
		tile.Draw(screen)
	}
}

func (s *Screen) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 300, 300
}

func NewScreen() *Screen {
	g := &Screen{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			g.board = append(g.board, Tile{
				X:   float64(i * 50),
				Y:   float64(j * 50),
				box: ebiten.NewImage(50, 50),
			})
		}
	}
	return g
}
