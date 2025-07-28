package starfield

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	MaxSpeed = 5
)

type Game struct {
	ScreenWidth  int
	ScreenHeight int

	starfield []*Star
}

func (g *Game) Setup(starNum int) {
	for i := 0; i < starNum; i++ {
		star := NewStar(g.ScreenWidth, g.ScreenHeight)
		g.starfield = append(g.starfield, star)
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, star := range g.starfield {
		ebitenutil.DrawRect(screen, float64(star.x), float64(star.y), 2, 2, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
