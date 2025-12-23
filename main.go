package main

import (
	"log"
	"ui-test/game"
	"ui-test/mover"
	"ui-test/shape"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	square := shape.NewSquareAt(50, 20, 20)
	square2 := shape.NewSquareAt(30, 100, 100)
	mover := mover.NewMover()

	game := game.NewGame(mover, []*shape.Shape{square, square2})
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
