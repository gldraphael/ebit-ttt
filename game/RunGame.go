package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func RunGame() {
	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("Hello World Test")
	if err := ebiten.RunGame(newScreen()); err != nil {
		log.Fatal(err)
	}
}
