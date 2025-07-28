package starfield

import (
	"image/color"
	"starfield-simulation/utils"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const maxSpeed = 25

var speed = 0.0

type Game struct {
	ScreenWidth  float64
	ScreenHeight float64

	starfield []*Star
}

func (g *Game) Setup(starNum int) {
	for i := 0; i < starNum; i++ {
		star := NewStar(g.ScreenWidth, g.ScreenHeight)
		g.starfield = append(g.starfield, star)
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	cx, _ := ebiten.CursorPosition()
	if cx < 0 || cx > int(g.ScreenWidth) { // not sure why get a very weird value at the start that screws everything
		return nil
	}

	speed = utils.Map(float64(cx), 0, float64(g.ScreenWidth), 1, maxSpeed)

	for idx, star := range g.starfield {
		star.z -= speed
		if star.z <= 0 {
			g.starfield[idx] = NewStar(g.ScreenWidth, g.ScreenHeight)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, star := range g.starfield {
		// Translate the star's position based on its z value
		sx := utils.Map(star.x/star.z, 0, 1, 0, float64(g.ScreenWidth)) + float64(g.ScreenWidth/2)
		sy := utils.Map(star.y/star.z, 0, 1, 0, float64(g.ScreenHeight)) + float64(g.ScreenHeight/2)

		// Get the previous position based on the speed
		psx := utils.Map(star.x/(star.z+speed), 0, 1, 0, float64(g.ScreenWidth)) + float64(g.ScreenWidth/2)
		psy := utils.Map(star.y/(star.z+speed), 0, 1, 0, float64(g.ScreenHeight)) + float64(g.ScreenHeight/2)

		// Calculate the radius based on the z value
		r := utils.Map(star.z, 0, float64(g.ScreenWidth), 2, 0)

		// Draw the star and its trace
		ebitenutil.DrawRect(screen, sx, sy, r, r, color.White)
		ebitenutil.DrawLine(screen, sx, sy, psx, psy, color.RGBA{240, 240, 240, 255})
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(g.ScreenWidth), int(g.ScreenHeight)
}
