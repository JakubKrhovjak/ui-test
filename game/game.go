package game

import (
	"time"
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
	if g.isHit() {
		g.target.Bling(uiColor.Red)
		g.destroy()
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
		if g.isHit() {
			return
		}

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

func (g *Game) isHit() bool {
	return g.inX() && g.inY()
}

func (g *Game) inX() bool {
	return g.cube.X < g.target.X+g.target.Size && g.cube.X+g.cube.Size > g.target.X
}

func (g *Game) inY() bool {
	return g.cube.Y < g.target.Y+g.target.Size && g.cube.Y+g.cube.Size > g.target.Y
}

func (g *Game) destroy() {
	go func() {
		time.Sleep(time.Duration(500) * time.Millisecond)
		g.target = shape.NewSquareAt(50, 0, 20)
		g.shapes[0] = g.target
	}()
}
