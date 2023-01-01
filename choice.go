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

func (w *Choice) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Choice_handle((*C.Fl_Choice)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *Choice) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Choice_resize_callback((*C.Fl_Choice)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *Choice) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Choice_draw((*C.Fl_Choice)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}
