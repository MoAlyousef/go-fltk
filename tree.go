package fltk

/*
#include <stdlib.h>
#include "include/cfltk/cfl_tree.h"
#include "include/cfltk/cfl_enums.h"
*/
import "C"
import "unsafe"

type Tree struct {
	Group
}

func NewTree(x, y, w, h int, text ...string) *Tree {
	t := &Tree{}
	initGroup(t, unsafe.Pointer(C.Fl_Tree_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return t
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
