package fltk

/*
#include "cfltk/cfl_menu.h"
*/
import "C"
import "unsafe"

type Choice struct {
	menu
}

func NewChoice(x, y, w, h int, text ...string) *Choice {
	c := &Choice{}
	initWidget(c, unsafe.Pointer(C.Fl_Choice_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return c
}
