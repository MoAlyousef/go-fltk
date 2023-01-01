package fltk

/*
#include "cfltk/cfl_group.h"
#include <stdint.h>
#include "fltk.h"
*/
import "C"
import (
	"unsafe"
)

type Group struct {
	widget
	// Child widgets in an unspecified order
	deletionHandlerId uintptr
}

func NewGroup(x, y, w, h int, text ...string) *Group {
	g := &Group{}
	initWidget(g, unsafe.Pointer(C.Fl_Group_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	g.setDeletionCallback(g.onDelete)
	return g
}

func (b *Group) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Group_set_deletion_callback((*C.Fl_Group)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *Group) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Group_handle((*C.Fl_Group)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *Group) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Group_resize_callback((*C.Fl_Group)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *Group) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Group_draw((*C.Fl_Group)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

func (g *Group) getGroup() *Group {
	return g
}

func (g *Group) Begin() {
	C.Fl_Group_begin((*C.Fl_Group)(g.ptr()))
}

func (g *Group) End() {
	C.Fl_Group_end((*C.Fl_Group)(g.ptr()))
}

func (g *Group) Add(w Widget) {
	C.Fl_Group_add((*C.Fl_Group)(g.getWidget().ptr()), (unsafe.Pointer)(w.getWidget().ptr()))
}

func (g *Group) Remove(w Widget) {
	C.Fl_Group_remove((*C.Fl_Group)(g.getWidget().ptr()), (unsafe.Pointer)(w.getWidget().ptr()))
}

func (g *Group) Clear() {
	C.Fl_Group_clear((*C.Fl_Group)(g.getWidget().ptr()))
}

func (g *Group) Resizable(w Widget) {
	C.Fl_Group_resizable((*C.Fl_Group)(g.ptr()), unsafe.Pointer(w.getWidget().ptr()))
}

func (g *Group) DrawChildren() {
	C.Fl_Group_draw_children((*C.Fl_Group)(g.ptr()))
}

func (g *Group) Destroy() {
	g.Clear()
	C.Fl_Group_delete((*C.Fl_Group)(g.ptr()))
}
