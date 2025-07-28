package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1280, 720
}

func main() {
	ebiten.SetWindowSize(1280, 720)

	game := &Game{}
	ebiten.RunGame(game)
}
