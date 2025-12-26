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
	target *shape.Shape
	cube   *shape.Shape
}

func NewGame(mover *mover.Mover, shapes map[int]*shape.Shape) *Game {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebiten - Pohyblivý čtverec")
	return &Game{
		x:      screenWidth/2 - squareSize/2,
		mover:  mover,
		shapes: shapes,
		target: shapes[0],
		cube:   shapes[1],
	}
}

func (g *Game) Update() error {
	g.handleTarget(g.target)
	g.handleCube()

	// AABB collision: check if cube overlaps target by at least 1 pixel
	if g.cube.X < g.target.X+g.target.Size && g.cube.X+g.cube.Size > g.target.X &&
		g.cube.Y < g.target.Y+g.target.Size && g.cube.Y+g.cube.Size > g.target.Y {
		if g.target.Blinking == false {
			g.target.Bling(uiColor.Red)
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

func (g *Game) handleCube() {
	for index, shape := range g.shapes {
		key := ebiten.Key(int(ebiten.Key1) + index)
		if ebiten.IsKeyPressed(key) && index > 0 {
			shape.SetColor(uiColor.Blue)
			g.mover.Move(shape)
		}
		//if !ebiten.IsKeyPressed(key) {
		//	shape.SetColor(uiColor.White)
		//}
	}

}
