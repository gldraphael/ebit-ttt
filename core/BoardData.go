package core

type BoardData struct {
	tiles [3][3]TileValue
}

func newBoardData() BoardData {
	bd := BoardData{}

	for i := range 3 {
		for j := range 3 {
			bd.tiles[i][j] = Tile_Empty
		}
	}

	return bd
}

func (bd BoardData) IsGameOver() bool {
	return false
}

func (bd BoardData) GetWinner() TileValue {
	if !bd.IsGameOver() {
		return Tile_Empty
	}
	return Tile_X // TODO: fix this
}
