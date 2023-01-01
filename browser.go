package fltk

/*
#include <stdlib.h>
#include "cfltk/cfl_browser.h"
#include "fltk.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type Browser struct {
	Group
	icons        map[int]Image
	columnWidths []int
	dataMap      map[uintptr]interface{}
	lastDataID   uintptr
}

var (
	InvalidLine = errors.New("line doesn't exist")
)

func NewBrowser(x, y, w, h int, text ...string) *Browser {
	b := &Browser{}
	b.dataMap = make(map[uintptr]interface{})
	b.icons = make(map[int]Image)
	initWidget(b, unsafe.Pointer(C.Fl_Browser_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	b.setDeletionCallback(b.onDelete)
	return b
}

func (w *Browser) setDeletionCallback(handler func()) {
	w.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Browser_set_deletion_callback((*C.Fl_Browser)(w.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(w.deletionHandlerId))
}

func (b *Browser) Add(str string) {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))

	C.Fl_Browser_add((*C.Fl_Browser)(b.ptr()), cStr)
}

func (b *Browser) AddWithData(str string, data interface{}) {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))

	b.lastDataID++
	id := b.lastDataID
	b.dataMap[id] = data

	C.Fl_Browser_add((*C.Fl_Browser)(b.ptr()), cStr)
	sz := b.Size() - 1
	C.Fl_Browser_set_data((*C.Fl_Browser)(b.ptr()), C.int(sz), unsafe.Pointer(id))
}

// func (b *Browser) TopLine() int {
// 	return int(C.Fl_Browser_topline((*C.Fl_Browser)(b.ptr())))
// }

func (b *Browser) SetBottomLine(line int) error {
	if line < 1 || line > b.Size() {
		return InvalidLine
	}

	C.Fl_Browser_bottomline((*C.Fl_Browser)(b.ptr()), C.int(line))
	return nil
}

func (b *Browser) SetMiddleLine(line int) error {
	if line < 1 || line > b.Size() {
		return InvalidLine
	}

	C.Fl_Browser_middleline((*C.Fl_Browser)(b.ptr()), C.int(line))
	return nil
}

func (b *Browser) SetTopLine(line int) error {
	if line < 1 || line > b.Size() {
		return InvalidLine
	}

	C.Fl_Browser_topline((*C.Fl_Browser)(b.ptr()), C.int(line))
	return nil
}

func (b *Browser) Clear() {
	for k := range b.icons {
		delete(b.icons, k)
	}
	b.dataMap = make(map[uintptr]interface{})
	C.Fl_Browser_clear((*C.Fl_Browser)(b.ptr()))
}

func (b *Browser) Remove(line int) error {
	if line < 1 || line > b.Size() {
		return InvalidLine
	}
	delete(b.icons, line)

	// TODO: got the id from C++ is expensive, need a better way to delete go reference
	id := uintptr(C.Fl_Browser_data((*C.Fl_Browser)(b.ptr()), C.int(line)))
	delete(b.dataMap, id)

	C.Fl_Browser_remove((*C.Fl_Browser)(b.ptr()), C.int(line))
	return nil
}

func (b *Browser) ColumnChar() rune {
	return rune(C.Fl_Browser_column_char((*C.Fl_Browser)(b.ptr())))
}

func (b *Browser) SetColumnChar(r rune) {
	cStr := C.CString(string(r))
	defer C.free(unsafe.Pointer(cStr))

	C.Fl_Browser_set_column_char((*C.Fl_Browser)(b.ptr()), *cStr)
}

func (b *Browser) HideLine(line int) error {
	if line < 1 || line > b.Size() {
		return InvalidLine
	}

	C.Fl_Browser_hide_line((*C.Fl_Browser)(b.ptr()), C.int(line))
	return nil
}

func (b *Browser) Size() int {
	return int(C.Fl_Browser_size((*C.Fl_Browser)(b.ptr())))
}

func (b *Browser) Icon(line int) Image {
	return b.icons[line]
}

func (b *Browser) SetIcon(line int, i Image) {
	if i == nil {
		delete(b.icons, line)
		C.Fl_Browser_set_icon((*C.Fl_Browser)(b.ptr()), C.int(line), nil)
		return
	}

	b.icons[line] = i
	C.Fl_Browser_set_icon((*C.Fl_Browser)(b.ptr()), C.int(line), unsafe.Pointer(b.icons[line].getImage().ptr()))
}

func (b *Browser) FormatChar() rune {
	return rune(C.Fl_Browser_format_char((*C.Fl_Browser)(b.ptr())))
}

func (b *Browser) SetFormatChar(r rune) {
	cStr := C.CString(string(r))
	defer C.free(unsafe.Pointer(cStr))

	C.Fl_Browser_set_format_char((*C.Fl_Browser)(b.ptr()), *cStr)
}

func (b *Browser) Displayed(line int) bool {
	if C.Fl_Browser_displayed((*C.Fl_Browser)(b.ptr()), C.int(line)) == 1 {
		return true
	}

	return false
}

func (b *Browser) Text(line int) string {
	cStr := C.Fl_Browser_text((*C.Fl_Browser)(b.ptr()), C.int(line))
	//defer C.free(unsafe.Pointer(cStr))

	return C.GoString(cStr)
}

func (b *Browser) SetColumnWidths(widths ...int) {
	cArr := make([]C.int, len(widths), len(widths))
	for i, v := range widths {
		cArr[i] = C.int(v)
	}

	b.columnWidths = widths
	C.Fl_Browser_set_column_widths((*C.Fl_Browser)(b.ptr()), (*C.int)(&cArr[0]))
}

// Store column widths in Go instead of calling from C++ as it's complex and expensive
// to convert between them
func (b *Browser) ColumnWidths() []int {
	return b.columnWidths
}

func (b *Browser) Data(line int) interface{} {
	id := uintptr(C.Fl_Browser_data((*C.Fl_Browser)(b.ptr()), C.int(line)))
	return b.dataMap[id]
}

func (b *Browser) Value() int {
	return int(C.Fl_Browser_value((*C.Fl_Browser)(b.ptr())))
}

func (b *Browser) SetValue(line int) {
	C.Fl_Browser_select((*C.Fl_Browser)(b.ptr()), C.int(line))
}

type SelectBrowser struct {
	Browser
}

func NewSelectBrowser(x, y, w, h int, text ...string) *SelectBrowser {
	b := &SelectBrowser{}
	b.dataMap = make(map[uintptr]interface{})
	initWidget(b, unsafe.Pointer(C.Fl_Select_Browser_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	b.setDeletionCallback(b.onDelete)
	return b
}

func (w *SelectBrowser) setDeletionCallback(handler func()) {
	w.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Select_Browser_set_deletion_callback((*C.Fl_Select_Browser)(w.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(w.deletionHandlerId))
}

type HoldBrowser struct {
	Browser
}

func NewHoldBrowser(x, y, w, h int, text ...string) *HoldBrowser {
	b := &HoldBrowser{}
	b.dataMap = make(map[uintptr]interface{})
	initWidget(b, unsafe.Pointer(C.Fl_Hold_Browser_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	b.setDeletionCallback(b.onDelete)
	return b
}

func (w *HoldBrowser) setDeletionCallback(handler func()) {
	w.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Hold_Browser_set_deletion_callback((*C.Fl_Hold_Browser)(w.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(w.deletionHandlerId))
}

type MultiBrowser struct {
	Browser
}

func NewMultiBrowser(x, y, w, h int, text ...string) *MultiBrowser {
	b := &MultiBrowser{}
	b.dataMap = make(map[uintptr]interface{})
	initWidget(b, unsafe.Pointer(C.Fl_Multi_Browser_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	b.setDeletionCallback(b.onDelete)
	return b
}

func (w *MultiBrowser) setDeletionCallback(handler func()) {
	w.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Multi_Browser_set_deletion_callback((*C.Fl_Multi_Browser)(w.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(w.deletionHandlerId))
}