package main

import (
	"log"
	"ui-test/game"
	"ui-test/mover"
	"ui-test/shape"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	shapes := make(map[int]*shape.Shape, 2)
	shapes[0] = shape.NewSquareAt(50, 200, 20)  // klávesa 1
	shapes[1] = shape.NewSquareAt(30, 200, 200) // klávesa 2
	mover := mover.NewMover()

	game := game.NewGame(mover, shapes)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
