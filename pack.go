package fltk

/*
#include "cfltk/cfl_group.h"
*/
import "C"
import "unsafe"

type Pack struct {
	Group
}

func NewPack(x, y, w, h int, text ...string) *Pack {
	p := &Pack{}
	initWidget(p, unsafe.Pointer(C.Fl_Pack_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return p
}

type PackType uint8

var (
	VERTICAL   = PackType(0)
	HORIZONTAL = PackType(1)
)

func (p *Pack) SetType(packType PackType) {
	p.widget.SetType(uint8(packType))
}

func (p *Pack) SetSpacing(spacing int) {
	C.Fl_Pack_set_spacing((*C.Fl_Pack)(p.ptr()), C.int(spacing))
}
