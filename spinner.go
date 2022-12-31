package fltk

/*
#include "cfltk/cfl_enums.h"
#include "cfltk/cfl_input.h"
#include "cfltk/cfl_misc.h"
*/
import "C"
import "unsafe"

type Spinner struct {
	widget
}

func NewSpinner(x, y, w, h int, text ...string) *Spinner {
	s := &Spinner{}
	initWidget(s, unsafe.Pointer(C.Fl_Spinner_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return s
}

type SpinnerInputType uint8

var SPINNER_INT_INPUT = SpinnerInputType(0)
var SPINNER_FLOAT_INPUT = SpinnerInputType(1)

func (s *Spinner) SetType(inputType SpinnerInputType) {
	C.Fl_Spinner_set_type((*C.Fl_Spinner)(s.ptr()), C.int(inputType))
}

func (s *Spinner) SetMaximum(max float64) {
	C.Fl_Spinner_set_maximum((*C.Fl_Spinner)(s.ptr()), C.double(max))
}

func (s *Spinner) SetMinimum(min float64) {
	C.Fl_Spinner_set_minimum((*C.Fl_Spinner)(s.ptr()), C.double(min))
}

func (s *Spinner) SetStep(step float64) {
	C.Fl_Spinner_set_step((*C.Fl_Spinner)(s.ptr()), C.double(step))
}

func (s *Spinner) SetValue(val float64) {
	C.Fl_Spinner_set_value((*C.Fl_Spinner)(s.ptr()), C.double(val))
}

func (s *Spinner) Value() float64 {
	return (float64)(C.Fl_Spinner_value((*C.Fl_Spinner)(s.ptr())))
}
