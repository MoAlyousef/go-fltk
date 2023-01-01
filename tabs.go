package fltk

/*
#include "cfltk/cfl_group.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Tabs struct {
	Group
}

func NewTabs(x, y, w, h int, text ...string) *Tabs {
	i := &Tabs{}
	initWidget(i, unsafe.Pointer(C.Fl_Tabs_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	i.setDeletionCallback(i.onDelete)
	return i
}

func (b *Tabs) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Tabs_set_deletion_callback((*C.Fl_Tabs)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (t *Tabs) Value() int {
	ptr := C.Fl_Tabs_value((*C.Fl_Tabs)(t.ptr()))
	val := C.Fl_Tabs_find((*C.Fl_Tabs)(t.ptr()), unsafe.Pointer(ptr))
	return int(val)
}

func (t *Tabs) SetValue(value int) {
	child := C.Fl_Tabs_child((*C.Fl_Tabs)(t.ptr()), C.int(value))
	C.Fl_Tabs_set_value((*C.Fl_Tabs)(t.ptr()), child)
}
