package math

import "math"

func Step(x float64) float64 {
	if x > 0 {
		return 1
	}

	return 0
}

func ReLU(x float64) float64 {
	return Max(0, x)
}

func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}
