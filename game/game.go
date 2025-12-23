package game

import (
	"ui-test/mover"
	"ui-test/shape"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	squareSize   = 50
	moveSpeed    = 3
)

type Game struct {
	x, y   float64
	mover  *mover.Mover
	shapes []*shape.Shape
}

func NewGame(mover *mover.Mover, shapes []*shape.Shape) *Game {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebiten - Pohyblivý čtverec")
	return &Game{
		x:      screenWidth/2 - squareSize/2,
		y:      screenHeight/2 - squareSize/2,
		mover:  mover,
		shapes: shapes,
	}
}

func (g *Game) Update() error {
	for _, shape := range g.shapes {
		g.mover.Move(shape)
	}
	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, shape := range g.shapes {
		shape.DrawAtPosition(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
