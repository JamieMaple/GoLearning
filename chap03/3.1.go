package main

import "math"

// ....

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	if r != 0 {
		return math.Sin(r) / r
	} else {
		return 0
	}
}
