package posev

import (
	"github.com/mitsuse/matrix-go/dense"

	//~ "log"
	"math"
	"testing"
)

func TestPowerTopEigen(t *testing.T) {
	eps := float64(0.00001)
	a := dense.New(3, 3)(
		-261, 209, -49,
		-530, 422, -98,
		-800, 631, -144,
	)
	v, e := PowerTopEigen(a, 40, eps)

	eps = float64(0.001)
	if math.Abs(e-float64(10)) > eps {
		t.Errorf("Top eigenvector failed: %f", e)
	}
	if d := DeltaEigen(e, v, a); d > eps {
		t.Errorf("Eigenvector check failed: %f", d)
	}
}

func TestPowerTopSingular(t *testing.T) {
	eps := float64(0.00001)
	b := dense.New(4, 5)(
		1, 0, 0, 0, 2,
		0, 0, 3, 0, 0,
		0, 0, 0, 0, 0,
		0, 4, 0, 0, 0,
	)
	c := []float64{4, 3, math.Sqrt(5), 0}

	u, v, s := PowerTopKSingular(b, 4, 40, eps)

	m, n := b.Shape()
	eps = float64(0.001)
	for i := 0; i < 4; i++ {
		if math.Abs(s[i]-c[i]) > eps {
			t.Errorf("Top %d singular value failed: %f", i, s[i])
		}
		ui := u.View(0, i, m, 1)
		vi := v.View(0, i, n, 1)
		if d := DeltaSingular(s[i], ui, vi, b); d > eps {
			t.Errorf("Left %d singular vector check failed: %f", i, d)
		}
		if d := DeltaSingular(s[i], vi, ui, b.Transpose()); d > eps {
			t.Errorf("Right %d singular vector check failed: %f", i, d)
		}
	}
}
