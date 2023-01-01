package fltk

/*
#include "cfltk/cfl_menu.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Choice struct {
	menu
}

func NewChoice(x, y, w, h int, text ...string) *Choice {
	c := &Choice{}
	initWidget(c, unsafe.Pointer(C.Fl_Choice_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	c.setDeletionCallback(c.onDelete)
	return c
}

func (b *Choice) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Choice_set_deletion_callback((*C.Fl_Choice)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}