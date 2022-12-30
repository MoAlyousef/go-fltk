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

type groupInterface interface {
	Widget
	getGroup() *Group
}

func initGroup(g groupInterface, p unsafe.Pointer) {
	initWidget(g, p)
	// group := g.getGroup()
	// group.setDeletionHandler(group.onDelete)
}

func NewGroup(x, y, w, h int, text ...string) *Group {
	g := &Group{}
	initGroup(g, unsafe.Pointer(C.Fl_Group_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return g
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

func (g *Group) onDelete() {
	// if g.deletionHandlerId > 0 {
	// 	globalCallbackMap.unregister(g.deletionHandlerId)
	// }
	// g.deletionHandlerId = 0
	// for i := range g.children {
	// 	g.children[i].getWidget().parent = nil
	// 	g.children[i] = nil
	// }
	// g.children = g.children[:0]
}

func (w *Group) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Group_draw((*C.Fl_Group)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))

}
