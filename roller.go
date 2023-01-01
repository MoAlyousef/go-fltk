package fltk

/*
#include "cfltk/cfl_valuator.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Roller struct {
	valuator
}

func NewRoller(x, y, w, h int, text ...string) *Roller {
	r := &Roller{}
	initWidget(r, unsafe.Pointer(C.Fl_Roller_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	r.setDeletionCallback(r.onDelete)
	return r
}

func (b *Roller) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Roller_set_deletion_callback((*C.Fl_Roller)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *Roller) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Roller_handle((*C.Fl_Roller)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *Roller) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Roller_resize_callback((*C.Fl_Roller)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *Roller) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Roller_draw((*C.Fl_Roller)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}
