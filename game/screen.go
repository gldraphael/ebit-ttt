package game

import (
	"log"

	"github.com/gldraphael/ebit-ttt/core"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type screen struct {
	board [3][3]tile
	state *core.GameState
}

func (s *screen) Update() error {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {

		cx, cy := ebiten.CursorPosition()
		log.Default().Printf("Cursor position: %d, %d\n", cx, cy)

		x, y := int(cx/100), int(cy/100)
		log.Default().Printf("Tile   position: %d, %d\n", x, y)

		if err := s.state.MakeMove(x, y); err != nil {
			log.Default().Println(err)
		}

	}
	return nil
}

func (s *screen) Draw(sc *ebiten.Image) {
	for i, row := range s.board {
		for j, t := range row {
			tile, err := s.state.Tile(j, i)
			if err != nil {
				log.Default().Println(err)
				continue
			}
			t.Draw(sc, tile, j, i)
		}
	}
}

func (s *screen) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 300, 300
}

func newScreen() *screen {
	screen := &screen{}
	screen.state = core.NewGameState()
	return screen
}
