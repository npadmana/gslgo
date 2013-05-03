package deriv

import (
	"math"
	"testing"
)

type Quad struct {
	a, b, c float64
}

func (q *Quad) Eval(x float64) float64 {
	return q.a*x*x + q.b*x + q.c
}

func (q *Quad) Deriv(x float64) float64 {
	return 2*q.a*x + q.b
}

func TestDeriv1(t *testing.T) {
	var x, dd, dderr float64
	var err error
	q1 := Quad{1, 0, 0}
	ff := func(x float64) float64 { return q1.Eval(x) }
	x = 2.5

	// Central
	dd, dderr, err = Diff(Central, ff, x, 0.001)
	if err != nil {
		t.Error("Unexpected error")
	}
	if math.Abs(dd-q1.Deriv(x)) > 1.e-5 {
		t.Errorf("Derivative failed : x=%f, expected=%f, actual=%f, error=%f", x, q1.Deriv(x), dd, dderr)
	}

	// Forward 
	dd, dderr, err = Diff(Forward, ff, x, 0.001)
	if err != nil {
		t.Error("Unexpected error")
	}
	if math.Abs(dd-q1.Deriv(x)) > 1.e-5 {
		t.Errorf("Derivative failed : x=%f, expected=%f, actual=%f, error=%f", x, q1.Deriv(x), dd, dderr)
	}

	// Reverse
	dd, dderr, err = Diff(Backward, ff, x, 0.001)
	if err != nil {
		t.Error("Unexpected error")
	}
	if math.Abs(dd-q1.Deriv(x)) > 1.e-5 {
		t.Errorf("Derivative failed : x=%f, expected=%f, actual=%f, error=%f", x, q1.Deriv(x), dd, dderr)
	}

}

func TestDeriv2(t *testing.T) {
	x := 1.0
	h := 0.001
	y := math.Cos(x)
	dd, dderr, err := Diff(Central,
		func(x float64) float64 { return math.Sin(x) },
		x, h)

	if err != nil {
		t.Error("Unexpected error")
	}

	if math.Abs(y-dd) > 1.e-5 {
		t.Errorf("Derivative failed : x=%f, expected=%f, actual=%f, error=%f", x, y, dd, dderr)
	}
}
