package fltk

/*
#include "cfltk/cfl_box.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Box struct {
	widget
}

func NewBox(boxType BoxType, x, y, w, h int, text ...string) *Box {
	b := &Box{}
	initWidget(b, unsafe.Pointer(C.Fl_Box_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	b.SetBox(boxType)
	b.setDeletionCallback(b.onDelete)
	return b
}

func (w *Box) setDeletionCallback(handler func()) {
	w.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Box_set_deletion_callback((*C.Fl_Box)(w.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(w.deletionHandlerId))
}

func (w *Box) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Box_handle((*C.Fl_Box)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *Box) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Box_resize_callback((*C.Fl_Box)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *Box) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Box_draw((*C.Fl_Box)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}
