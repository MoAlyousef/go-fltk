package fltk

/*
#include <stdlib.h>
#include "cfltk/cfl_window.h"
#include "cfltk/cfl_image.h"
#include "cfltk/cfl_enums.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Window struct {
	Group
	icons []*RgbImage
}

func NewWindow(w, h int, text ...string) *Window {
	win := &Window{}
	ptr := C.Fl_Double_Window_new(C.int(0), C.int(0), C.int(w), C.int(h), cStringOpt(text))
	C.Fl_Double_Window_free_position(ptr)
	initWidget(win, unsafe.Pointer(ptr))
	win.setDeletionCallback(win.onDelete)
	return win
}

func (w *Window) setDeletionCallback(handler func()) {
	w.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Double_Window_set_deletion_callback((*C.Fl_Double_Window)(w.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(w.deletionHandlerId))
}

func (w *Window) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Double_Window_handle((*C.Fl_Double_Window)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *Window) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Double_Window_resize_callback((*C.Fl_Double_Window)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *Window) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Double_Window_draw((*C.Fl_Double_Window)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

func (w *Window) IsShown() bool {
	return C.Fl_Double_Window_shown((*C.Fl_Double_Window)(w.ptr())) != 0
}

func (w *Window) XRoot() int {
	return int(C.Fl_Double_Window_x_root((*C.Fl_Double_Window)(w.ptr())))
}

func (w *Window) YRoot() int {
	return int(C.Fl_Double_Window_y_root((*C.Fl_Double_Window)(w.ptr())))
}

func (w *Window) SetLabel(label string) {
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	C.Fl_Double_Window_set_label((*C.Fl_Double_Window)(w.ptr()), labelStr)
}

func (w *Window) SetCursor(cursor Cursor) {
	C.Fl_Double_Window_set_cursor((*C.Fl_Double_Window)(w.ptr()), C.int(cursor))
}

func (w *Window) SetFullscreen(flag bool) {
	f := 0
	if flag {
		f = 1
	}
	C.Fl_Double_Window_fullscreen((*C.Fl_Double_Window)(w.ptr()), C.uint(f))
}

func (w *Window) FullscreenActive() bool {
	return C.Fl_Double_Window_fullscreen_active((*C.Fl_Double_Window)(w.ptr())) != 0
}

func (w *Window) SetModal() {
	C.Fl_Double_Window_make_modal((*C.Fl_Double_Window)(w.ptr()), C.uint(1))
}

func (w *Window) SetNonModal() {
	C.Fl_Double_Window_make_modal((*C.Fl_Double_Window)(w.ptr()), C.uint(0))
}

func (w *Window) SetIcons(icons []*RgbImage) {
	images := make([]unsafe.Pointer, 0, len(icons))
	for _, icon := range icons {
		images = append(images, unsafe.Pointer(icon.iPtr))
	}
	C.Fl_Double_Window_set_icons((*C.Fl_Double_Window)(w.ptr()), &images[0], C.int(len(images)))
	w.icons = icons
}

type Cursor int

var (
	CURSOR_DEFAULT = Cursor(C.Fl_Cursor_Default)
	CURSOR_ARROW   = Cursor(C.Fl_Cursor_Arrow)
	CURSOR_CROSS   = Cursor(C.Fl_Cursor_Cross)
	CURSOR_WAIT    = Cursor(C.Fl_Cursor_Wait)
	CURSOR_INSERT  = Cursor(C.Fl_Cursor_Insert)
	CURSOR_HAND    = Cursor(C.Fl_Cursor_Hand)
	CURSOR_HELP    = Cursor(C.Fl_Cursor_Help)
	CURSOR_MOVE    = Cursor(C.Fl_Cursor_Move)
	CURSOR_NS      = Cursor(C.Fl_Cursor_NS)
	CURSOR_WE      = Cursor(C.Fl_Cursor_WE)
	CURSOR_NWSE    = Cursor(C.Fl_Cursor_NWSE)
	CURSOR_NESW    = Cursor(C.Fl_Cursor_NESW)
	CURSOR_N       = Cursor(C.Fl_Cursor_N)
	CURSOR_NE      = Cursor(C.Fl_Cursor_NE)
	CURSOR_E       = Cursor(C.Fl_Cursor_E)
	CURSOR_SE      = Cursor(C.Fl_Cursor_SE)
	CURSOR_S       = Cursor(C.Fl_Cursor_S)
	CURSOR_SW      = Cursor(C.Fl_Cursor_SW)
	CURSOR_W       = Cursor(C.Fl_Cursor_W)
	CURSOR_NW      = Cursor(C.Fl_Cursor_NW)
)
