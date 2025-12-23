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
	// Pohyb pomocí šipek
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.x -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.x += moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.y -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.y += moveSpeed
	}

	// Omezení na hranice obrazovky
	if g.x < 0 {
		g.x = 0
	}
	if g.x > screenWidth-squareSize {
		g.x = screenWidth - squareSize
	}
	if g.y < 0 {
		g.y = 0
	}
	if g.y > screenHeight-squareSize {
		g.y = screenHeight - squareSize
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, shape := range g.shapes {
		shape.Draw(screen, g.x, g.y)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
