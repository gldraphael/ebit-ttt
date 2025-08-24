package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("Hello World Test")
	if err := ebiten.RunGame(NewScreen()); err != nil {
		log.Fatal(err)
	}
}
