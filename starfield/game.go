package starfield

import "github.com/hajimehoshi/ebiten"

type Game struct {
	ScreenWidth  int
	ScreenHeight int
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
