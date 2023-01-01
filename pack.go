package fltk

/*
#include "cfltk/cfl_group.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Pack struct {
	Group
}

func NewPack(x, y, w, h int, text ...string) *Pack {
	p := &Pack{}
	initWidget(p, unsafe.Pointer(C.Fl_Pack_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	p.setDeletionCallback(p.onDelete)
	return p
}

func (b *Pack) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Pack_set_deletion_callback((*C.Fl_Pack)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
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
