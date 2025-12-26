package game

import (
	uiColor "ui-test/color"
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

var (
	direction = "left"
)

type Game struct {
	x, y   float64
	mover  *mover.Mover
	shapes map[int]*shape.Shape
}

func NewGame(mover *mover.Mover, shapes map[int]*shape.Shape) *Game {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebiten - Pohyblivý čtverec")
	return &Game{
		x:      screenWidth/2 - squareSize/2,
		mover:  mover,
		shapes: shapes,
	}
}

func (g *Game) Update() error {
	handleTarget(g.shapes[0])

	for index, shape := range g.shapes {
		key := ebiten.Key(int(ebiten.Key1) + index)
		if ebiten.IsKeyPressed(key) && index > 0 {
			shape.SetColor(uiColor.Red)
			g.mover.Move(shape)
		}
		if !ebiten.IsKeyPressed(key) {
			shape.SetColor(uiColor.White)
		}
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

func (g *Game) handleTarget(target *shape.Shape) {
	go func() {
		if target.X >= screenWidth-target.Size {
			direction = "left"
		}

		if target.X == 0 {
			direction = "right"
		}

		if direction == "left" {
			g.mover.MoveAt(target, -3)
		} else if direction == "right" {
			g.mover.MoveAt(target, 3)
		}

	}()
}
