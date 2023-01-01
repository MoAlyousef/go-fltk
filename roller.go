package fltk

/*
#include "cfltk/cfl_valuator.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Roller struct {
	valuator
}

func NewRoller(x, y, w, h int, text ...string) *Roller {
	r := &Roller{}
	initWidget(r, unsafe.Pointer(C.Fl_Roller_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	r.setDeletionCallback(r.onDelete)
	return r
}

func (b *Roller) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Roller_set_deletion_callback((*C.Fl_Roller)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}