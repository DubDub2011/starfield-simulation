package starfield

import (
	"math/rand"
)

type Star struct {
	x, y, z float64
}

func NewStar(widthLimit, heightLimit int) *Star {
	rx := rand.Float64()*float64(widthLimit*2) - float64(widthLimit)
	ry := rand.Float64()*float64(heightLimit*2) - float64(heightLimit)
	rz := float64(widthLimit)

	return &Star{
		x: rx,
		y: ry,
		z: rz,
	}
}
