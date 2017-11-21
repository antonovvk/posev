package posev

import (
	"github.com/mitsuse/matrix-go"
	"github.com/mitsuse/matrix-go/dense"
	"math"
	"math/rand"
)

// matrix.Matrix.Scalar() mutates the matrix, it is not always handy
func ScalarMult(m matrix.Matrix, s float64) matrix.Matrix {
	r := dense.Zeros(m.Shape())
	c := m.All()

	for c.HasNext() {
		v, i, j := c.Get()
		r.Update(i, j, v*s)
	}
	return r
}

func VecNorm(m matrix.Matrix) float64 {
	n := float64(0)
	c := m.All()
	for c.HasNext() {
		v, _, _ := c.Get()
		n += v * v
	}
	return math.Sqrt(n)
}

func VecNormalize(m matrix.Matrix) matrix.Matrix {
	n := VecNorm(m)
	return ScalarMult(m, float64(1)/n)
}

func randUnitVec(n int) matrix.Matrix {
	r := dense.Zeros(n, 1)
	for i := 0; i < n; i++ {
		r.Update(i, 0, rand.NormFloat64())
	}
	return VecNormalize(r)
}

func PowerTopEigen(a matrix.Matrix, maxIters int, eps float64) (v matrix.Matrix, e float64) {
	v = randUnitVec(a.Columns())

	o := float64(0)
	for i := 0; i < maxIters; i++ {
		z := a.Multiply(v)
		v = VecNormalize(z)
		e = v.Transpose().Multiply(z).Get(0, 0)
		if math.Abs((e-o)/e) < eps {
			return
		}
		o = e
	}
	return nil, 0
}

func DeltaEigen(e float64, v, a matrix.Matrix) float64 {
	dim := v.Rows()
	id := dense.Zeros(dim, dim)
	for i := 0; i < dim; i++ {
		id.Update(i, i, e)
	}
	z := id.Subtract(a).Multiply(v)
	return VecNorm(z)
}

func PowerTopSingular(a matrix.Matrix, maxIters int, eps float64) (u, v matrix.Matrix, s float64) {
	n := a.Columns()
	r := randUnitVec(n)

	b := a.Transpose()
	for i := 0; i < maxIters; i++ {
		u = VecNormalize(a.Multiply(r))
		v = b.Multiply(u)
		s = VecNorm(v)
		v.Scalar(float64(1) / s)

		d := float64(0)
		for j := 0; j < n; j++ {
			c := math.Abs(r.Get(j, 0) - v.Get(j, 0))
			if d < c {
				d = c
			}
		}
		if d < eps {
			return
		}
		r = v
	}
	return nil, nil, 0
}

func DeltaSingular(s float64, u, v, a matrix.Matrix) float64 {
	z := a.Multiply(v).Subtract(ScalarMult(u, s))
	return VecNorm(z)
}

func subSingular(vec *[]float64, sv matrix.Matrix) {
	if sv == nil || len(*vec) == 0 {
		return
	}
	for i := range *vec {
		w := float64(0)
		for j, v := range *vec {
			w += sv.Get(j, 0) * v
		}
		(*vec)[i] -= sv.Get(i, 0) * w
	}
}
