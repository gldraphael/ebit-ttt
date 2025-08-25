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
	// TODO: do not generate a new image on the fly
	//       instead, cache 3 tile images and reuse them
	img := ebiten.NewImage(TILE_SIZE, TILE_SIZE)

	// TODO: the tiles can surely look much better than these...
	switch tile {

	case core.Tile_X:
		img.Fill(color.RGBA{R: 255, G: 0, B: 0, A: 0})

	case core.Tile_Y:
		img.Fill(color.RGBA{R: 0, G: 255, B: 0, A: 0})

	default:
		img.Fill(color.RGBA{R: 0, G: 0, B: 255, A: 0})
	}

	return img
}
