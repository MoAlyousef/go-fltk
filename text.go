package fltk

/*
#include <stdlib.h>
#include "include/cfltk/cfl_text.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type TextBuffer struct {
	cPtr *C.Fl_Text_Buffer
}

var ErrTextBufferDestroyed = errors.New("text buffer is destroyed")

func NewTextBuffer() *TextBuffer {
	ptr := C.Fl_Text_Buffer_new()
	return &TextBuffer{ptr}
}

func (b *TextBuffer) ptr() *C.Fl_Text_Buffer {
	if b.cPtr == nil {
		panic(ErrTextBufferDestroyed)
	}
	return b.cPtr
}
func (b *TextBuffer) Destroy() {
	C.Fl_Text_Buffer_delete(b.ptr())
	b.cPtr = nil
}

func (b *TextBuffer) SetText(txt string) {
	txtstr := C.CString(txt)
	defer C.free(unsafe.Pointer(txtstr))
	C.Fl_Text_Buffer_set_text(b.ptr(), txtstr)
}

func (b *TextBuffer) Append(txt string) {
	txtstr := C.CString(txt)
	defer C.free(unsafe.Pointer(txtstr))
	C.Fl_Text_Buffer_append(b.ptr(), txtstr)
}

func (b *TextBuffer) Text() string {
	return C.GoString(C.Fl_Text_Buffer_text(b.ptr()))
}

type TextDisplay struct {
	widget
}

func NewTextDisplay(x, y, w, h int, text ...string) *TextDisplay {
	t := &TextDisplay{}
	initWidget(t, unsafe.Pointer(C.Fl_Text_Display_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return t
}

func (t *TextDisplay) SetBuffer(buf *TextBuffer) {
	C.Fl_Text_Display_set_buffer((*C.Fl_Text_Display)(t.ptr()), buf.ptr())
}

// wrapMargin is not needed if WrapMode is WRAP_NONE or WRAP_AT_BOUNDS
func (t *TextDisplay) SetWrapMode(wrap WrapMode, wrapMargin ...int) {
	if len(wrapMargin) < 1 {
		wrapMargin = append(wrapMargin, 0)
	}

	C.Fl_Text_Display_wrap_mode((*C.Fl_Text_Display)(t.ptr()), C.int(wrap), C.int(wrapMargin[0]))
}

func (t *TextDisplay) Buffer() *TextBuffer {
	ptr := C.Fl_Text_Display_get_buffer((*C.Fl_Text_Display)(t.ptr()))
	return &TextBuffer{ptr}
}

// MoveRight moves the current insert position right one character.
//Returns true if the cursor moved, false if the end of the text was reached
func (t *TextDisplay) MoveRight() bool {
	return (int)(C.Fl_Text_Display_move_right((*C.Fl_Text_Display)(t.ptr()))) != 0
}

//MoveLeft moves the current insert position left one character.
func (t *TextDisplay) MoveLeft() bool {
	return (int)(C.Fl_Text_Display_move_left((*C.Fl_Text_Display)(t.ptr()))) != 0
}

//MoveUp moves the current insert position up one line.
func (t *TextDisplay) MoveUp() bool {
	return (int)(C.Fl_Text_Display_move_up((*C.Fl_Text_Display)(t.ptr()))) != 0
}

//MoveDown moves the current insert position down one line.
func (t *TextDisplay) MoveDown() bool {
	return (int)(C.Fl_Text_Display_move_down((*C.Fl_Text_Display)(t.ptr()))) != 0
}

//ShowInsertPosition scrolls the text buffer to show the current insert position.
func (t *TextDisplay) ShowInsertPosition() {
	C.Fl_Text_Display_show_insert_position((*C.Fl_Text_Display)(t.ptr()))
}

//TextSize gets the default size of text in the widget
func (t *TextDisplay) TextSize() int {
	return (int)(C.Fl_Text_Display_text_size((*C.Fl_Text_Display)(t.ptr())))
}

//SetTextSize sets the default size of text in the widget
func (t *TextDisplay) SetTextSize(size int) {
	C.Fl_Text_Display_set_text_size((*C.Fl_Text_Display)(t.ptr()), C.int(size))
}

type TextEditor struct {
	TextDisplay
}

func NewTextEditor(x, y, w, h int, text ...string) *TextEditor {
	t := &TextEditor{}
	initWidget(t, unsafe.Pointer(C.Fl_Text_Editor_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return t
}
