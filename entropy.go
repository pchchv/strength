package strength

import "math"

func logX(base, n float64) float64 {
	if base == 0 {
		return 0
	}
	// change of base formulae
	return math.Log2(n) / math.Log2(base)
}
