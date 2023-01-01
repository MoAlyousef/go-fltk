package fltk

/*
#include "cfltk/cfl_enums.h"
#include "cfltk/cfl_input.h"
#include "cfltk/cfl_misc.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Spinner struct {
	widget
}

func NewSpinner(x, y, w, h int, text ...string) *Spinner {
	s := &Spinner{}
	initWidget(s, unsafe.Pointer(C.Fl_Spinner_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	s.setDeletionCallback(s.onDelete)
	return s
}

func (b *Spinner) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Spinner_set_deletion_callback((*C.Fl_Spinner)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *Spinner) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Spinner_handle((*C.Fl_Spinner)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *Spinner) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Spinner_resize_callback((*C.Fl_Spinner)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *Spinner) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Spinner_draw((*C.Fl_Spinner)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

type SpinnerInputType uint8

var SPINNER_INT_INPUT = SpinnerInputType(0)
var SPINNER_FLOAT_INPUT = SpinnerInputType(1)

func (s *Spinner) SetType(inputType SpinnerInputType) {
	C.Fl_Spinner_set_type((*C.Fl_Spinner)(s.ptr()), C.int(inputType))
}

func (s *Spinner) SetMaximum(max float64) {
	C.Fl_Spinner_set_maximum((*C.Fl_Spinner)(s.ptr()), C.double(max))
}

func (s *Spinner) SetMinimum(min float64) {
	C.Fl_Spinner_set_minimum((*C.Fl_Spinner)(s.ptr()), C.double(min))
}

func (s *Spinner) SetStep(step float64) {
	C.Fl_Spinner_set_step((*C.Fl_Spinner)(s.ptr()), C.double(step))
}

func (s *Spinner) SetValue(val float64) {
	C.Fl_Spinner_set_value((*C.Fl_Spinner)(s.ptr()), C.double(val))
}

func (s *Spinner) Value() float64 {
	return (float64)(C.Fl_Spinner_value((*C.Fl_Spinner)(s.ptr())))
}
