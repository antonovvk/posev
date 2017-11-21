package posev

import (
    "github.com/mitsuse/matrix-go/dense"

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
    if math.Abs(e - float64(10)) > eps {
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

	u, v, s := PowerTopSingular(b, 40, eps)

    eps = float64(0.001)
    if math.Abs(s - float64(4)) > eps {
        t.Errorf("Top singular vector failed: %f", s)
    }
	if d := DeltaSingular(s, u, v, b); d > eps {
        t.Errorf("Left singular vector check failed: %f", d)
    }
	if d := DeltaSingular(s, v, u, b.Transpose()); d > eps {
        t.Errorf("Right singular vector check failed: %f", d)
    }
}
