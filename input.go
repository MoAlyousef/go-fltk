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
