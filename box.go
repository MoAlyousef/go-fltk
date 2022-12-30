package fltk

/*
#include "cfltk/cfl_box.h"
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
	return b
}
