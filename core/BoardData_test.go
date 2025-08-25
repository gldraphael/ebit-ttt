package core

import (
	"testing"
)

func TestNewBoardData_InitializesWithEmptyTiles(t *testing.T) {
	board := newBoardData()

	if len(board.tiles) != 3 {
		t.Error("Expected 3 rows in board")
	}

	for i := range 3 {
		if len(board.tiles[i]) != 3 {
			t.Errorf("Expected 3 columns at row %d", i)
		}
		for j := range 3 {
			if board.tiles[i][j] != Tile_Empty {
				t.Errorf("Expected tile at (%d, %d) to be Tile_Empty, got %v", i, j, board.tiles[i][j])
			}
		}
	}
}
