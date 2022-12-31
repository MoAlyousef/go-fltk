package fltk

/*
#include "cfltk/cfl_group.h"
#include "cfltk/cfl_enums.h"
*/
import "C"
import "unsafe"

type Scroll struct {
	Group
}

func NewScroll(x, y, w, h int, text ...string) *Scroll {
	s := &Scroll{}
	initWidget(s, unsafe.Pointer(C.Fl_Scroll_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return s
}

func (s *Scroll) ScrollTo(x, y int) {
	C.Fl_Scroll_scroll_to((*C.Fl_Scroll)(s.ptr()), C.int(x), C.int(y))
}

func (s *Scroll) XPosition() int {
	return int(C.Fl_Scroll_xposition((*C.Fl_Scroll)(s.ptr())))
}

func (s *Scroll) YPosition() int {
	return int(C.Fl_Scroll_yposition((*C.Fl_Scroll)(s.ptr())))
}

type ScrollType uint8

var (
	SCROLL_HORIZONTAL        = ScrollType(C.Fl_ScrollType_Horizontal)
	SCROLL_VERTICAL          = ScrollType(C.Fl_ScrollType_Vertical)
	SCROLL_BOTH              = ScrollType(C.Fl_ScrollType_Both)
	SCROLL_HORIZONTAL_ALWAYS = ScrollType(C.Fl_ScrollType_HorizontalAlways)
	SCROLL_VERTICAL_ALWAYS   = ScrollType(C.Fl_ScrollType_VerticalAlways)
	SCROLL_BOTH_ALWAYS       = ScrollType(C.Fl_ScrollType_BothAlways)
)

func (s *Scroll) SetType(scrollType ScrollType) {
	s.widget.SetType(uint8(scrollType))
}
