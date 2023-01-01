package fltk

/*
#include "cfltk/cfl_misc.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Progress struct {
	widget
}

func NewProgress(x, y, w, h int, text ...string) *Progress {
	p := &Progress{}
	initWidget(p, unsafe.Pointer(C.Fl_Progress_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	p.setDeletionCallback(p.onDelete)
	return p
}

func (b *Progress) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Progress_set_deletion_callback((*C.Fl_Progress)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *Progress) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Progress_handle((*C.Fl_Progress)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *Progress) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Progress_resize_callback((*C.Fl_Progress)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *Progress) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Progress_draw((*C.Fl_Progress)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

func (p *Progress) SetMaximum(max float64) {
	C.Fl_Progress_set_maximum((*C.Fl_Progress)(p.ptr()), C.double(max))
}

func (p *Progress) Maximum() float64 {
	return float64(C.Fl_Progress_maximum((*C.Fl_Progress)(p.ptr())))
}

func (p *Progress) SetMinimum(max float64) {
	C.Fl_Progress_set_minimum((*C.Fl_Progress)(p.ptr()), C.double(max))
}

func (p *Progress) Minimum() float64 {
	return float64(C.Fl_Progress_minimum((*C.Fl_Progress)(p.ptr())))
}

func (p *Progress) SetValue(value float64) {
	C.Fl_Progress_set_value((*C.Fl_Progress)(p.ptr()), C.double(value))
}

func (p *Progress) Value() float64 {
	return float64(C.Fl_Progress_value((*C.Fl_Progress)(p.ptr())))
}
