package game

import (
	"github.com/gldraphael/ebit-ttt/core"
	"github.com/hajimehoshi/ebiten/v2"
)

type screen struct {
	board [3][3]tile
}

func (s *screen) Update() error {
	return nil
}

func (s *screen) Draw(sc *ebiten.Image) {
	for i, row := range s.board {
		for j, t := range row {
			t.Draw(sc, core.Tile_Empty, i, j)
		}
	}
}

func (s *screen) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 300, 300
}

func newScreen() *screen {
	g := &screen{}
	return g
}
