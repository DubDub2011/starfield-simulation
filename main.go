package main

import (
	"starfield-simulation/starfield"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)

	game := &starfield.Game{
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}

	game.Setup(1000)

	ebiten.RunGame(game)
}
