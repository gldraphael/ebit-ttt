package game

import (
	"fmt"
	"image/color"

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

func (s *screen) Draw(sc *ebiten.Image) {
	// Implement screen drawing logic here
	for _, t := range s.board {
		t.Draw(sc)
	}
}

func (s *screen) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 300, 300
}

func newScreen() *screen {
	fmt.Println("Screen created")
	g := &screen{}
	for i := range 3 {
		for j := range 3 {
			img := ebiten.NewImage(50, 50)
			img.Fill(color.RGBA{255, uint8(i * 10 / 255), uint8(j * 10 / 255), 255})
			g.board = append(g.board, tile{
				X:     float64(i * 50),
				Y:     float64(j * 50),
				Image: img,
			})
		}
	}
	return g
}
