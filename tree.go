package fltk

/*
#include <stdlib.h>
#include "cfltk/cfl_tree.h"
#include "cfltk/cfl_enums.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Tree struct {
	Group
}

func NewTree(x, y, w, h int, text ...string) *Tree {
	t := &Tree{}
	initWidget(t, unsafe.Pointer(C.Fl_Tree_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	t.setDeletionCallback(t.onDelete)
	return t
}

func (b *Tree) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Tree_set_deletion_callback((*C.Fl_Tree)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *Tree) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Tree_handle((*C.Fl_Tree)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *Tree) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Tree_resize_callback((*C.Fl_Tree)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *Tree) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Tree_draw((*C.Fl_Tree)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

func (t *Tree) SetShowRoot(show bool) {
	if show {
		C.Fl_Tree_set_showroot((*C.Fl_Tree)(t.ptr()), 1)
	} else {
		C.Fl_Tree_set_showroot((*C.Fl_Tree)(t.ptr()), 0)
	}
}

type TreeItem struct {
	ptr *C.Fl_Tree_Item
}

func (t *Tree) Add(path string) TreeItem {
	pathStr := C.CString(path)
	defer C.free(unsafe.Pointer(pathStr))
	itemPtr := C.Fl_Tree_add((*C.Fl_Tree)(t.ptr()), pathStr)
	return TreeItem{ptr: itemPtr}
}

func (t TreeItem) SetWidget(w Widget) {
	C.Fl_Tree_Item_set_widget(t.ptr, w.getWidget().ptr())
}

type TreeItemDrawMode uint

var (
	TreeItemDrawDefault        = TreeItemDrawMode(C.Fl_TreeItemDrawMode_Default)
	TreeItemDrawLabelAndWidget = TreeItemDrawMode(C.Fl_TreeItemDrawMode_LabelAndWidget)
	TreeItemHeightFromWidget   = TreeItemDrawMode(C.Fl_TreeItemDrawMode_HeightFromWidget)
)

func (t *Tree) SetItemDrawMode(drawMode TreeItemDrawMode) {
	C.Fl_Tree_set_item_draw_mode((*C.Fl_Tree)(t.ptr()), C.int(drawMode))
}

type TreeConnector int

var (
	TreeConnectorNone   = TreeConnector(C.Fl_TreeConnectorStyle_None)
	TreeConnectorDotted = TreeConnector(C.Fl_TreeConnectorStyle_Dotted)
	TreeConnectorSolid  = TreeConnector(C.Fl_TreeConnectorStyle_Solid)
)

func (t *Tree) SetConnectorStyle(style TreeConnector) {
	C.Fl_Tree_set_connectorstyle((*C.Fl_Tree)(t.ptr()), C.int(style))
}
