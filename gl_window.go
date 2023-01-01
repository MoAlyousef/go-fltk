package fltk

/*
#include "cfltk/cfl_window.h"
#include "fltk.h"
*/
import "C"
import (
	"unsafe"
)

type GlWindow struct {
	Window
	deletionHandlerId uintptr
	drawFunId         uintptr
}

func NewGlWindow(x, y, w, h int, drawFun func(), text ...string) *GlWindow {
	win := &GlWindow{}
	win.drawFunId = globalCallbackMap.register(drawFun)
	initWidget(win, unsafe.Pointer(C.Fl_Gl_Window_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	win.SetDrawHandler(drawFun)
	win.setDeletionCallback(win.onDelete)
	return win
}

func (b *GlWindow) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Gl_Window_set_deletion_callback((*C.Fl_Gl_Window)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *GlWindow) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Gl_Window_handle((*C.Fl_Gl_Window)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *GlWindow) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Gl_Window_resize_callback((*C.Fl_Gl_Window)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *GlWindow) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Gl_Window_draw((*C.Fl_Gl_Window)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

func (w *GlWindow) MakeCurrent() {
	C.Fl_Gl_Window_make_current((*C.Fl_Gl_Window)(w.ptr()))
}

func (w *GlWindow) Destroy() {
	if w.drawFunId > 0 {
		globalCallbackMap.unregister(w.drawFunId)
	}
	w.drawFunId = 0
	w.Window.Destroy()
}

func (w *GlWindow) ContextValid() bool {
	return C.Fl_Gl_Window_context_valid((*C.Fl_Gl_Window)(w.ptr())) != 0
}

func (w *GlWindow) Valid() bool {
	return C.Fl_Gl_Window_valid((*C.Fl_Gl_Window)(w.ptr())) != 0
}

func (w *GlWindow) CanDo() bool {
	return C.Fl_Gl_Window_can_do((*C.Fl_Gl_Window)(w.ptr())) != 0
}

func (w *GlWindow) SetMode(mode int) {
	C.Fl_Gl_Window_set_mode((*C.Fl_Gl_Window)(w.ptr()), C.int(mode))
}
