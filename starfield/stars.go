package starfield

import (
	"math/rand"
)

type Star struct {
	x, y, z float64
}

func NewStar(widthLimit, heightLimit float64) *Star {
	rx := rand.Float64()*widthLimit*2 - widthLimit
	ry := rand.Float64()*heightLimit*2 - heightLimit
	rz := rand.Float64() * widthLimit

	return &Star{
		x: rx,
		y: ry,
		z: rz,
	}
}
