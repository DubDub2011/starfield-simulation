package starfield

import (
	"image/color"
	"starfield-simulation/utils"

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
	for _, star := range g.starfield {
		star.z--
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, star := range g.starfield {
		sx := utils.Map(star.x/star.z, 0, 1, 0, float64(g.ScreenWidth)) + float64(g.ScreenWidth/2)
		sy := utils.Map(star.y/star.z, 0, 1, 0, float64(g.ScreenHeight)) + float64(g.ScreenHeight/2)

		ebitenutil.DrawRect(screen, sx, sy, 2, 2, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
