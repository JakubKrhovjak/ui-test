package main

import (
	"log"
	"ui-test/game"
	"ui-test/mover"
	"ui-test/shape"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	square := shape.NewSquare(50)
	square2 := shape.NewSquare(30)
	mover := mover.NewMover()

	game := game.NewGame(mover, []*shape.Shape{square, square2})
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
