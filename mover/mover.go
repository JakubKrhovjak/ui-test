package mover

import (
	"ui-test/shape"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	moveSpeed    = 3
	screenWidth  = 640
	screenHeight = 480
)

type Mover struct {
}

func NewMover() *Mover {
	return &Mover{}
}

func (m *Mover) Move(object *shape.Shape) {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		object.X -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		object.X += moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		object.Y -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		object.Y += moveSpeed
	}

	checkBounds(object)
}

func (m *Mover) MoveAt(object *shape.Shape, y float64) {
	object.Y += y
	checkBounds(object)
}

func checkBounds(object *shape.Shape) {
	bounds := object.Bounds()
	width := float64(bounds.Dx())
	height := float64(bounds.Dy())

	if object.X < 0 {
		object.X = 0
	}
	if object.X > screenWidth-width {
		object.X = screenWidth - width
	}
	if object.Y < 0 {
		object.Y = 0
	}
	if object.Y > screenHeight-height {
		object.Y = screenHeight - height
	}
}
