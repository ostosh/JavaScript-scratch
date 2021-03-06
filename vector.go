package feature

import (
	"errors"
	"math"
)

//returns mean of given vector
func Mean(x []float64) float64 {
	s := Sum(x)

	n := float64(len(x))

	return s / n
}

//returns variance of given vector
func Variance(x []float64) float64 {
	n := float64(len(x))
	if n == 1 {
		return 0
	} else if n < 2 {
		n = 2
	}

	m := Mean(x)

	ss := 0.0
	for _, v := range x {
		ss += math.Pow(v - m, 2.0)
	}

	return ss / (n - 1)
}


//returns standard deviation of given vector
func StandardDeviation(x []float64) float64 {
	return math.Sqrt(Variance(x))
}

//returns sum of given vector
func Sum(x []float64) float64 {
	s := 0.0
	for _, v := range x {
		s += v
	}
	return s
}

//returns product of given vectors
func Product(x, y []float64) ([]float64, error) {
	if len(x) != len(y) {
		return nil, errors.New("vectors must be same length")
	}

	p := make([]float64, len(x))
	for i, _ := range x {
		p[i] = x[i] * y[i]
	}
	return p, nil
}

//returns dot product of given vectors
func DotProduct(x, y []float64) (float64, error) {
	p, err := Product(x, y)
	if err != nil {
		return -1, err
	}
	return Sum(p), nil
}

//returns norm of given vector 
func Norm(x []float64, pow float64) float64 {
	s := 0.0

	for _, xval := range x {
		s += math.Pow(xval, pow)
	}

	return math.Pow(s, 1 / pow)
}

//returns cosine of given vectors
func Cosine(x, y []float64) (float64, error) {
	d, err := DotProduct(x, y)
	if err != nil {
		return -1, err
	}

	xnorm := Norm(x, 2.0)
	ynorm := Norm(y, 2.0)

	return d / (xnorm * ynorm), nil
}

//returns correlation of given vectors
func Correlation(x, y []float64) (float64, error) {
	n := float64(len(x))
	xy, err := Product(x, y)
	if err != nil {
		return -1, err
	}

	sx := StandardDeviation(x)
	sy := StandardDeviation(y)

	mx := Mean(x)
	my := Mean(y)

	r := (Sum(xy) - n * mx * my) / ((n - 1) * sx * sy)
	return r, nil
}
