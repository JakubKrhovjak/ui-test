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
	shapes[1] = shape.NewSquareAt(50, 20, 20)
	shapes[2] = shape.NewSquareAt(30, 100, 100)
	mover := mover.NewMover()

	game := game.NewGame(mover, shapes)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
