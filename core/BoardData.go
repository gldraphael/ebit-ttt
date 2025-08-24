package core

type BoardData struct {
	tiles [3][3]TileValue
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
