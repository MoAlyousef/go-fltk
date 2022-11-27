package fltk

/*
#include "include/cfltk/cfl_group.h"
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
	children          []Widget
	deletionHandlerId uintptr
}

type groupInterface interface {
	Widget
	getGroup() *Group
}

var toplevelGroup *Group = &Group{}
var currentGroup groupInterface = toplevelGroup

func initGroup(g groupInterface, p unsafe.Pointer) {
	initWidget(g, p)
	group := g.getGroup()
	group.deletionHandlerId = group.addDeletionHandler(group.onDelete)
	currentGroup = g
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
	currentGroup = g
}
func (g *Group) End() {
	C.Fl_Group_end((*C.Fl_Group)(g.ptr()))
	currentGroup = g.parent
}

func (g *Group) removeChild(child Widget) {
	childWidget := child.getWidget()
	childWidget.parent = nil
	childPtr := childWidget.ptr()
	for i, c := range g.children {
		if c.getWidget().ptr() == childPtr {
			childrenCount := len(g.children)
			g.children[i] = g.children[childrenCount-1]
			g.children[childrenCount-1] = nil
			g.children = g.children[:childrenCount-1]
			return
		}
	}
}

func (g *Group) Add(w Widget) {
	C.Fl_Group_add((*C.Fl_Group)(g.ptr()), unsafe.Pointer(w.getWidget().ptr()))
	if ww := w.getWidget(); ww.parent != nil {
		ww.parent.getGroup().removeChild(w)
	}
	w.getWidget().parent = g
	g.children = append(g.children, w)
}
func (g *Group) Remove(w Widget) {
	C.Fl_Group_remove((*C.Fl_Group)(g.ptr()), unsafe.Pointer(w.getWidget().ptr()))
	g.removeChild(w)
	w.getWidget().parent = toplevelGroup
}

func (g *Group) Resizable(w Widget) {
	C.Fl_Group_resizable((*C.Fl_Group)(g.ptr()), unsafe.Pointer(w.getWidget().ptr()))
}
func (g *Group) DrawChildren() {
	C.Fl_Group_draw_children((*C.Fl_Group)(g.ptr()))
}

func (g *Group) onDelete() {
	if g.deletionHandlerId > 0 {
		globalCallbackMap.unregister(g.deletionHandlerId)
	}
	g.deletionHandlerId = 0
	for i := range g.children {
		g.children[i].getWidget().parent = nil
		g.children[i] = nil
	}
	g.children = g.children[:0]
	if currentGroup == g {
		currentGroup = nil
	}
}

func (w *Group) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Group_draw((*C.Fl_Group)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}
