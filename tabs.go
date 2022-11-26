package fltk

/*
#include "include/cfltk/cfl_group.h"
*/
import "C"
import "unsafe"

type Tabs struct {
	Group
}

func NewTabs(x, y, w, h int, text ...string) *Tabs {
	i := &Tabs{}
	initGroup(i, unsafe.Pointer(C.Fl_Tabs_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}

func (t *Tabs) Value() int {

	// val := C.Fl_Tabs_value((*C.Fl_Tabs)(t.ptr()))
	return int(0)
}

func (t *Tabs) SetValue(value int) {
	// C.Fl_Tabs_set_value((*C.Fl_Tabs)(t.ptr()), (C.int)(value))
}
