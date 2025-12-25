package shape

import (
	"image/color"
	uiColor "ui-test/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Shape struct {
	*ebiten.Image
	X, Y  float64
	color color.RGBA
}

func NewSquare(size int) *Shape {
	square := ebiten.NewImage(size, size)
	square.Fill(uiColor.White)
	return &Shape{Image: square, X: 0, Y: 0, color: uiColor.White}
}

func NewSquareAt(size int, x, y float64) *Shape {
	square := ebiten.NewImage(size, size)
	square.Fill(uiColor.White)
	return &Shape{Image: square, X: x, Y: y, color: uiColor.White}
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

func (s *Shape) SetColor(color color.RGBA) {
	s.Image.Fill(color)
}
