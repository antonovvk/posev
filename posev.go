package posev

import (
	"github.com/mitsuse/matrix-go"
	"github.com/mitsuse/matrix-go/dense"
	//~ "log"
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

func PowerTopSingular(a, r matrix.Matrix, maxIters int, eps float64) (u, v, w matrix.Matrix, s float64) {
	n := a.Columns()
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
		if d > 10*eps {
			w = r.Subtract(v)
		}
		if d < eps {
			return
		}
		r = v
	}
	return nil, nil, nil, 0
}

func HotellingDeflation(a, l, r matrix.Matrix, s float64) (b matrix.Matrix) {
	b = dense.Zeros(a.Shape())
	for i := 0; i < l.Rows(); i++ {
		for j := 0; j < r.Rows(); j++ {
			b.Update(i, j, a.Get(i, j)-s*l.Get(i, 0)*r.Get(j, 0))
		}
	}
	return
}

func PowerTopKSingular(a matrix.Matrix, k, maxIters int, eps float64) (u, v matrix.Matrix, s []float64) {
	tol := 0.000000001

	m := a.Rows()
	n := a.Columns()
	u = dense.Zeros(m, k)
	v = dense.Zeros(n, k)
	s = make([]float64, k)

	r := randUnitVec(n)
	for i := 0; i < k; i++ {
		ui, vi, w, si := PowerTopSingular(a, r, maxIters, eps)
		if si/(s[0]+tol) < tol {
			// Singluar values are too small
			return
		}
		s[i] = si
		for j := 0; j < m; j++ {
			u.Update(j, i, ui.Get(j, 0))
		}
		for j := 0; j < n; j++ {
			v.Update(j, i, vi.Get(j, 0))
		}
		a = HotellingDeflation(a, ui, vi, si)
		r = VecNormalize(w)
	}
	return
}

func DeltaSingular(s float64, u, v, a matrix.Matrix) float64 {
	z := a.Multiply(v).Subtract(ScalarMult(u, s))
	return VecNorm(z)
}
