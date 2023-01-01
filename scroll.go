package fltk

/*
#include "cfltk/cfl_group.h"
#include "cfltk/cfl_enums.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Scroll struct {
	Group
}

func NewScroll(x, y, w, h int, text ...string) *Scroll {
	s := &Scroll{}
	initWidget(s, unsafe.Pointer(C.Fl_Scroll_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	s.setDeletionCallback(s.onDelete)
	return s
}

func (b *Scroll) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Scroll_set_deletion_callback((*C.Fl_Scroll)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *Scroll) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Scroll_handle((*C.Fl_Scroll)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *Scroll) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Scroll_resize_callback((*C.Fl_Scroll)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *Scroll) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Scroll_draw((*C.Fl_Scroll)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
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
