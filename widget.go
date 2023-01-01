package fltk

/*
#include <stdlib.h>
#include <stdint.h>
#include "cfltk/cfl_widget.h"
#include "cfltk/cfl.h"
#include "fltk.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type widget struct {
	tracker           *C.Fl_Widget_Tracker
	callbackId        uintptr
	deletionHandlerId uintptr
	resizeHandlerId   uintptr
	drawHandlerId     uintptr
	eventHandlerId    uintptr
	image             Image
	deimage           Image
}

type Widget interface {
	getWidget() *widget
}

var ErrDestroyed = errors.New("widget is destroyed")

func initWidget(w Widget, p unsafe.Pointer) {
	ww := w.getWidget()
	ww.tracker = C.Fl_Widget_Tracker_new((*C.Fl_Widget)(p))
}

func (w *widget) ptr() *C.Fl_Widget {
	if !w.exists() {
		panic(ErrDestroyed)
	}
	return C.Fl_Widget_Tracker_widget(w.tracker)
}

func (w *widget) exists() bool {
	if w.tracker == nil {
		return false
	}
	return C.Fl_Widget_Tracker_exists(w.tracker) == 1
}

func (w *widget) SetCallback(f func()) {
	if w.callbackId > 0 {
		globalCallbackMap.unregister(w.callbackId)
	}
	w.callbackId = globalCallbackMap.register(f)
	C.Fl_Widget_set_callback(w.ptr(), (*C.Fl_Callback)(C.callback_handler), unsafe.Pointer(w.callbackId))
}

func (w *widget) SetCallbackCondition(when CallbackCondition) {
	C.Fl_Widget_set_when(w.ptr(), C.int(when))
}

func (w *widget) getWidget() *widget {
	return w
}

func (w *widget) onDelete() {
	if w.deletionHandlerId > 0 {
		globalCallbackMap.unregister(w.deletionHandlerId)
	}
	w.deletionHandlerId = 0
	if w.callbackId > 0 {
		globalCallbackMap.unregister(w.callbackId)
	}
	w.callbackId = 0
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = 0
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = 0
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = 0
	w.image = nil
	w.deimage = nil
	C.Fl_Widget_Tracker_delete(w.tracker)
	w.tracker = nil
}

func (w *widget) Destroy() {
	C.Fl_delete_widget(w.ptr())
}

func (w *widget) SetBox(box BoxType) {
	C.Fl_Widget_set_box(w.ptr(), C.int(box))
}

func (w *widget) SetLabelFont(font Font) {
	C.Fl_Widget_set_label_font(w.ptr(), C.int(font))
}

func (w *widget) SetLabelSize(size int) {
	C.Fl_Widget_set_label_size(w.ptr(), C.int(size))
}

func (w *widget) SetLabelType(ltype LabelType) {
	C.Fl_Widget_set_label_type(w.ptr(), C.int(ltype))
}

func (w *widget) SetLabelColor(col Color) {
	C.Fl_Widget_set_label_color(w.ptr(), C.uint(col))
}

func (w *widget) ClearVisibleFocus() {
	C.Fl_Widget_clear_visible_focus(w.ptr())
}

func (w *widget) X() int {
	return int(C.Fl_Widget_x(w.ptr()))
}

func (w *widget) Y() int {
	return int(C.Fl_Widget_y(w.ptr()))
}

func (w *widget) W() int {
	return int(C.Fl_Widget_width(w.ptr()))
}

func (w *widget) H() int {
	return int(C.Fl_Widget_height(w.ptr()))
}

func (w *widget) SetAlign(align Align) {
	C.Fl_Widget_set_align(w.ptr(), C.int(align))
}

func (w *widget) MeasureLabel() (int, int) {
	var width, height C.int
	C.Fl_Widget_measure_label(w.ptr(), &width, &height)
	return int(width), int(height)
}

func (w *widget) SetPosition(x, y int) {
	C.Fl_Widget_resize(w.ptr(), C.int(w.X()), C.int(w.Y()), C.int(x), C.int(y))
}

func (w *widget) Resize(x, y, width, height int) {
	C.Fl_Widget_resize(w.ptr(), C.int(x), C.int(y), C.int(width), C.int(height))
}

func (w *widget) Redraw() {
	C.Fl_Widget_redraw(w.ptr())
}

func (w *widget) Deactivate() {
	C.Fl_Widget_deactivate(w.ptr())
}

func (w *widget) Activate() {
	C.Fl_Widget_activate(w.ptr())
}

func (w *widget) SetType(widgetType uint8) {
	C.Fl_Widget_set_type(w.ptr(), C.int(widgetType))
}

func (w *widget) Show() {
	C.Fl_Widget_show(w.ptr())
}

func (w *widget) Hide() {
	C.Fl_Widget_hide(w.ptr())
}

func (w *widget) SelectionColor() Color {
	return Color(C.Fl_Widget_selection_color(w.ptr()))
}

func (w *widget) SetSelectionColor(color Color) {
	C.Fl_Widget_set_selection_color(w.ptr(), C.uint(color))
}

func (w *widget) SetColor(color Color) {
	C.Fl_Widget_set_color(w.ptr(), C.uint(color))
}

func (w *widget) SetLabel(label string) {
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	C.Fl_Widget_set_label(w.ptr(), labelStr)
}

func (w *widget) SetImage(i Image) {
	C.Fl_Widget_set_image(w.ptr(), unsafe.Pointer(i.getImage().ptr()))
	w.image = i
}

func (w *widget) SetDeimage(i Image) {
	C.Fl_Widget_set_deimage(w.ptr(), unsafe.Pointer(i.getImage().ptr()))
	w.deimage = i
}

func (w *widget) Box() BoxType {
	return BoxType(C.Fl_Widget_box(w.ptr()))
}

func (w *widget) LabelColor() Color {
	return Color(C.Fl_Widget_label_color(w.ptr()))
}

func (w *widget) Align() Align {
	return Align(C.Fl_Widget_align(w.ptr()))
}

func (w *widget) Type() uint8 {
	return uint8(C.Fl_Widget_get_type(w.ptr()))
}

func (w *widget) Label() string {
	return C.GoString(C.Fl_Widget_label(w.ptr()))
}

func (w *widget) Color() Color {
	return Color(C.Fl_Widget_color(w.ptr()))
}

func (w *widget) LabelFont() Font {
	return Font(C.Fl_Widget_label_font(w.ptr()))
}

func (w *widget) LabelSize() int {
	return int(C.Fl_Widget_label_size(w.ptr()))
}

func (w *widget) LabelType() LabelType {
	return LabelType(C.Fl_Widget_label_type(w.ptr()))
}

func (w *widget) SetTooltip(text string) {
	tooltipStr := C.CString(text)
	defer C.free(unsafe.Pointer(tooltipStr))
	C.Fl_Widget_set_tooltip(w.ptr(), tooltipStr)
}

func (w *widget) Parent() *Group {
	grp := &Group{}
	ptr := C.Fl_Widget_parent(w.ptr())
	initWidget(grp, ptr)
	return grp
}
