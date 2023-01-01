package fltk

/*
#include "cfltk/cfl_misc.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Progress struct {
	widget
}

func NewProgress(x, y, w, h int, text ...string) *Progress {
	p := &Progress{}
	initWidget(p, unsafe.Pointer(C.Fl_Progress_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	p.setDeletionCallback(p.onDelete)
	return p
}

func (b *Progress) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Progress_set_deletion_callback((*C.Fl_Progress)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (p *Progress) SetMaximum(max float64) {
	C.Fl_Progress_set_maximum((*C.Fl_Progress)(p.ptr()), C.double(max))
}

func (p *Progress) Maximum() float64 {
	return float64(C.Fl_Progress_maximum((*C.Fl_Progress)(p.ptr())))
}

func (p *Progress) SetMinimum(max float64) {
	C.Fl_Progress_set_minimum((*C.Fl_Progress)(p.ptr()), C.double(max))
}

func (p *Progress) Minimum() float64 {
	return float64(C.Fl_Progress_minimum((*C.Fl_Progress)(p.ptr())))
}

func (p *Progress) SetValue(value float64) {
	C.Fl_Progress_set_value((*C.Fl_Progress)(p.ptr()), C.double(value))
}

func (p *Progress) Value() float64 {
	return float64(C.Fl_Progress_value((*C.Fl_Progress)(p.ptr())))
}
