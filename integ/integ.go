// Package integ wraps the basic integration routines
package integ

/* 
#cgo pkg-config: gsl

#include <gsl/gsl_integration.h>

extern double integCB(double x, void *params);

static gsl_function mkintegCB(void *data) {
	gsl_function gf;
	gf.function = integCB;
	gf.params = data;
	return gf;
}

*/
import "C"

import (
	"errors"
	"github.com/npadmana/gslgo"
	"unsafe"
)

//export integCB
func integCB(x C.double, data unsafe.Pointer) C.double {
	ff := (*gslgo.GSLFuncWrapper)(data)
	return C.double(ff.Gofunc(float64(x)))
}

type WorkSpace struct {
	w *C.gsl_integration_workspace
}

func NewWork(n int) *WorkSpace {
	ret := new(WorkSpace)
	ret.w = C.gsl_integration_workspace_alloc(C.int(n))
	return ret
}

func (w *WorkSpace) Free() {
	C.gsl_integration_workspace_free(w.w)
}
