package fltk

/*
#include "cfltk/cfl_box.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Box struct {
	widget
}

func NewBox(boxType BoxType, x, y, w, h int, text ...string) *Box {
	b := &Box{}
	initWidget(b, unsafe.Pointer(C.Fl_Box_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	b.SetBox(boxType)
	b.setDeletionCallback(b.onDelete)
	return b
}

func (w *Box) setDeletionCallback(handler func()) {
	w.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Box_set_deletion_callback((*C.Fl_Box)(w.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(w.deletionHandlerId))
}
