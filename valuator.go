package fltk

/*
#include "include/cfltk/cfl_valuator.h"
*/
import "C"

type valuator struct {
	widget
}

func (v *valuator) SetMinimum(value float64) {
	C.Fl_Slider_set_minimum((*C.Fl_Slider)(v.ptr()), C.double(value))
}

func (v *valuator) SetMaximum(value float64) {
	C.Fl_Slider_set_maximum((*C.Fl_Slider)(v.ptr()), C.double(value))
}

func (v *valuator) SetStep(value float64) {
	C.Fl_Slider_set_step((*C.Fl_Slider)(v.ptr()), C.double(value), C.int(1))
}

func (v *valuator) Value() float64 {
	return float64(C.Fl_Slider_value((*C.Fl_Slider)(v.ptr())))
}

func (v *valuator) SetValue(value float64) {
	C.Fl_Slider_set_value((*C.Fl_Slider)(v.ptr()), C.double(value))
}
