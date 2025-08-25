package game

import (
	"image/color"

	"github.com/gldraphael/ebit-ttt/core"
	"github.com/hajimehoshi/ebiten/v2"
)

type tile struct{}

const TILE_SIZE = 100

func (t *tile) Draw(screen *ebiten.Image, tile core.TileValue, x int, y int) {
	image := getImageFor(tile)
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(TILE_SIZE*x), float64(TILE_SIZE*y))
	screen.DrawImage(image, options)
}

func getImageFor(tile core.TileValue) *ebiten.Image {
	img := ebiten.NewImage(TILE_SIZE, TILE_SIZE)
	// TODO: draw tiles for X, Y, and EMPTY
	img.Fill(color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 0,
	})
	return img
}
