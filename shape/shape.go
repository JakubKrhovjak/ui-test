package shape

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Shape struct {
	*ebiten.Image
	X, Y float64
}

func NewSquare(size int) *Shape {
	square := ebiten.NewImage(size, size)
	square.Fill(color.White)
	return &Shape{Image: square, X: 0, Y: 0}
}

func NewSquareAt(size int, x, y float64) *Shape {
	square := ebiten.NewImage(size, size)
	square.Fill(color.White)
	return &Shape{Image: square, X: x, Y: y}
}

func (s *Shape) Draw(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(s.Image, op)
}

func (s *Shape) DrawAtPosition(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.X, s.Y)
	screen.DrawImage(s.Image, op)
}
