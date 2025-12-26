package shape

import (
	"image/color"
	"time"
	uiColor "ui-test/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Shape struct {
	*ebiten.Image
	X, Y     float64
	color    color.RGBA
	Size     float64
	blinking bool
}

func NewSquare(size int) *Shape {
	square := ebiten.NewImage(size, size)
	square.Fill(uiColor.White)
	return &Shape{Image: square, X: 0, Y: 0, color: uiColor.White, Size: float64(size)}
}

func NewSquareAt(size int, x, y float64) *Shape {
	square := ebiten.NewImage(size, size)
	square.Fill(uiColor.White)
	return &Shape{Image: square, X: x, Y: y, color: uiColor.White, Size: float64(size)}
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

func (s *Shape) Bling(color color.RGBA) {
	go func() {
		for i := 0; i < 3; i++ {
			s.Image.Fill(color)
			time.Sleep(500 * time.Millisecond)
			s.Image.Fill(uiColor.White)
		}
	}()

}
