package unit_tests

import (
	"fmt"
	"math"
)

const (
	eps = 1e-4
)

func Solve(a, b, c float64) ([]float64, error) {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return nil, fmt.Errorf("NaN")
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return nil, fmt.Errorf("inf")
	}
	if math.Abs(a) <= eps {
		return nil, fmt.Errorf("коэффициент a равен 0")
	}
	d := b*b - 4*a*c
	// нет корней
	if d < -eps {
		return nil, nil
	}
	// 1 корень
	if math.Abs(d) <= eps {
		return []float64{-b / (2 * a), -b / (2 * a)}, nil
	}
	// 2 корня d > eps
	sqrtD := math.Sqrt(d)
	return []float64{(-b + sqrtD) / (2 * a), (-b - sqrtD) / (2 * a)}, nil
}
