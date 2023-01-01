package fltk

/*
#include <stdlib.h>
#include "cfltk/cfl_input.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Input struct {
	widget
}

func NewInput(x, y, w, h int, text ...string) *Input {
	i := &Input{}
	initWidget(i, unsafe.Pointer(C.Fl_Input_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	i.setDeletionCallback(i.onDelete)
	return i
}

func (b *Input) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Input_set_deletion_callback((*C.Fl_Input)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *Input) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Input_handle((*C.Fl_Input)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *Input) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Input_resize_callback((*C.Fl_Input)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *Input) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Input_draw((*C.Fl_Input)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

func (i *Input) Value() string {
	return C.GoString(C.Fl_Input_value((*C.Fl_Input)(i.ptr())))
}

func (i *Input) SetValue(value string) bool {
	valueStr := C.CString(value)
	defer C.free(unsafe.Pointer(valueStr))
	return C.Fl_Input_set_value((*C.Fl_Input)(i.ptr()), valueStr) != 0
}

func (i *Input) Resize(x int, y int, w int, h int) {
	C.Fl_Input_resize((*C.Fl_Input)(i.ptr()), C.int(x), C.int(y), C.int(w), C.int(h))
}

type Output struct {
	Input
}

func NewOutput(x, y, w, h int, text ...string) *Output {
	i := &Output{}
	initWidget(i, unsafe.Pointer(C.Fl_Output_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	i.setDeletionCallback(i.onDelete)
	return i
}

func (b *Output) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Output_set_deletion_callback((*C.Fl_Output)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *Output) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Output_handle((*C.Fl_Output)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *Output) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Output_resize_callback((*C.Fl_Output)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *Output) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Output_draw((*C.Fl_Output)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

type FloatInput struct {
	Input
}

func NewFloatInput(x, y, w, h int, text ...string) *FloatInput {
	i := &FloatInput{}
	initWidget(i, unsafe.Pointer(C.Fl_Float_Input_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	i.setDeletionCallback(i.onDelete)
	return i
}

func (b *FloatInput) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Float_Input_set_deletion_callback((*C.Fl_Float_Input)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *FloatInput) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Float_Input_handle((*C.Fl_Float_Input)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *FloatInput) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Float_Input_resize_callback((*C.Fl_Float_Input)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *FloatInput) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Float_Input_draw((*C.Fl_Float_Input)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}
