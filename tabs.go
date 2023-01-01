package fltk

/*
#include "cfltk/cfl_group.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Tabs struct {
	Group
}

func NewTabs(x, y, w, h int, text ...string) *Tabs {
	i := &Tabs{}
	initWidget(i, unsafe.Pointer(C.Fl_Tabs_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	i.setDeletionCallback(i.onDelete)
	return i
}

func (b *Tabs) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Tabs_set_deletion_callback((*C.Fl_Tabs)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *Tabs) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Tabs_handle((*C.Fl_Tabs)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *Tabs) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Tabs_resize_callback((*C.Fl_Tabs)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *Tabs) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Tabs_draw((*C.Fl_Tabs)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

func (t *Tabs) Value() int {
	ptr := C.Fl_Tabs_value((*C.Fl_Tabs)(t.ptr()))
	val := C.Fl_Tabs_find((*C.Fl_Tabs)(t.ptr()), unsafe.Pointer(ptr))
	return int(val)
}

func (t *Tabs) SetValue(value int) {
	child := C.Fl_Tabs_child((*C.Fl_Tabs)(t.ptr()), C.int(value))
	C.Fl_Tabs_set_value((*C.Fl_Tabs)(t.ptr()), child)
}
