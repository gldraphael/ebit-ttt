package core

import (
	"errors"
	"fmt"
)

/*
 * The game state
 */
type GameState struct {
	board         *BoardData
	currentPlayer TileValue
}

func NewGameState() *GameState {
	gs := GameState{}

	gs.board = newBoardData()
	gs.currentPlayer = Tile_X

	return &gs
}

/*
 * Returns the current player
 */
func (gs GameState) CurrentPlayer() TileValue {
	return gs.currentPlayer
}

/*
 * Returns the tile at the specified coordinate
 */
func (gs GameState) Tile(x int, y int) (TileValue, error) {
	if err := validateCoords(x, y); err != nil {
		return Tile_Empty, err
	}
	return gs.board.tiles[x][y], nil
}

/*
 * Make a move
 */
func (gs GameState) MakeMove(x int, y int) error {

	if err := validateCoords(x, y); err != nil {
		return err
	}

	if gs.board.tiles[x][y] != Tile_Empty {
		return errors.New(fmt.Sprintf("The tile (%d, %d) is not empty.", x, y))
	}

	gs.board.tiles[x][y] = gs.currentPlayer
	gs.nextPlayer()

	return nil
}

func validateCoords(x int, y int) error {
	if x < 0 || x >= 3 || y < 0 || y >= 3 {
		return errors.New(fmt.Sprintf("The tile (%d, %d) is invalid because it does not exist.", x, y))
	}
	return nil
}

func (gs GameState) nextPlayer() {
	if gs.currentPlayer == Tile_X {
		gs.currentPlayer = Tile_Y
	} else {
		gs.currentPlayer = Tile_X
	}
}
