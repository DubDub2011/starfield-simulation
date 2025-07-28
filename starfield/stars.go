package starfield

import (
	"math/rand"
)

type Star struct {
	x, y, z int
}

func NewStar(widthLimit, heightLimit int) *Star {
	rx := rand.Intn(widthLimit)
	ry := rand.Intn(heightLimit)
	rz := rand.Intn(MaxSpeed)

	return &Star{
		x: rx,
		y: ry,
		z: rz,
	}
}
