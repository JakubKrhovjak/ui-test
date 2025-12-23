package shape

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Shape struct {
	*ebiten.Image
}

func NewSquare(size int) *Shape {
	square := ebiten.NewImage(size, size)
	square.Fill(color.White)
	return &Shape{Image: square}
}

func (s *Shape) Draw(x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	s.DrawImage(s.Image, op)
}
