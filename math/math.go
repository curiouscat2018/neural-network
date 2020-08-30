package math

func DotProduct(a []float64, b []float64) float64 {
	if len(a) != len(b) {
		panic("DotProduct: length must be equal")
	}

	sum := 0.0
	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}

	return sum
}

func Max(a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}