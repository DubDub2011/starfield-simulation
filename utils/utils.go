// love a good utils package
package utils

func Map(value, inMin, inMax, outMin, outMax float64) float64 {
	return (value-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}
