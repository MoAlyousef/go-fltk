package fltk

/*
#include <stdlib.h>
#include "include/cfltk/cfl_misc.h"
*/
import "C"
import "unsafe"

type HelpView struct {
	widget
}

func NewHelpView(x, y, w, h int, text ...string) *HelpView {
	t := &HelpView{}
	initWidget(t, unsafe.Pointer(C.Fl_Help_View_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return t
}

func (h *HelpView) Directory() string {
	return C.GoString(C.Fl_Help_View_directory((*C.Fl_Help_View)(h.ptr())))
}

func (h *HelpView) Filename() string {
	return C.GoString(C.Fl_Help_View_filename((*C.Fl_Help_View)(h.ptr())))
}

func (h *HelpView) Find(str string, i ...int) int {
	if len(i) < 1 {
		i = append(i, 0)
	}

	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	return int(C.Fl_Help_View_find((*C.Fl_Help_View)(h.ptr()), cStr, C.int(i[0])))
}

func (h *HelpView) Load(f string) {
	fStr := C.CString(f)
	defer C.free(unsafe.Pointer(fStr))
	C.Fl_Help_View_load((*C.Fl_Help_View)(h.ptr()), fStr)
}

func (h *HelpView) LeftLine() int {
	return int(C.Fl_Help_View_leftline((*C.Fl_Help_View)(h.ptr())))
}

func (h *HelpView) SetLeftLine(i int) {
	C.Fl_Help_View_set_leftline((*C.Fl_Help_View)(h.ptr()), C.int(i))
}

func (h *HelpView) TopLine() int {
	return int(C.Fl_Help_View_topline((*C.Fl_Help_View)(h.ptr())))
}

func (h *HelpView) SetTopLine(i int) {
	C.Fl_Help_View_set_topline2((*C.Fl_Help_View)(h.ptr()), C.int(i))
}

func (h *HelpView) SetTopLineString(str string) {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C.Fl_Help_View_set_topline((*C.Fl_Help_View)(h.ptr()), cStr)
}

func (h *HelpView) Value() string {
	return C.GoString(C.Fl_Help_View_value((*C.Fl_Help_View)(h.ptr())))
}

func (h *HelpView) SetValue(str string) {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C.Fl_Help_View_set_value((*C.Fl_Help_View)(h.ptr()), cStr)
}
