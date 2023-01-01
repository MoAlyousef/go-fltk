package fltk

/*
#include "cfltk/cfl_button.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Button struct {
	widget
}

func NewButton(x, y, w, h int, text ...string) *Button {
	b := &Button{}
	ptr := C.Fl_Button_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))
	initWidget(b, unsafe.Pointer(ptr))
	b.setDeletionCallback(b.onDelete)
	return b
}

func (b *Button) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Button_set_deletion_callback((*C.Fl_Button)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *Button) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Button_handle((*C.Fl_Button)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *Button) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Button_resize_callback((*C.Fl_Button)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *Button) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Button_draw((*C.Fl_Button)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

func (b *Button) Value() bool {
	return C.Fl_Button_value((*C.Fl_Button)(b.ptr())) != C.int(0)
}

func (b *Button) SetValue(val bool) {
	if val {
		C.Fl_Button_set_value((*C.Fl_Button)(b.ptr()), 1)
	} else {
		C.Fl_Button_set_value((*C.Fl_Button)(b.ptr()), 0)
	}
}

func (b *Button) SetDownBox(box BoxType) {
	C.Fl_Button_set_down_box((*C.Fl_Button)(b.ptr()), C.int(box))
}

type CheckButton struct {
	Button
}

func NewCheckButton(x, y, w, h int, text ...string) *CheckButton {
	i := &CheckButton{}
	initWidget(i, unsafe.Pointer(C.Fl_Check_Button_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	i.setDeletionCallback(i.onDelete)
	return i
}

func (b *CheckButton) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Check_Button_set_deletion_callback((*C.Fl_Check_Button)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *CheckButton) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Check_Button_handle((*C.Fl_Check_Button)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *CheckButton) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Check_Button_resize_callback((*C.Fl_Check_Button)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *CheckButton) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Check_Button_draw((*C.Fl_Check_Button)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

type RadioButton struct {
	Button
}

func NewRadioButton(x, y, w, h int, text ...string) *RadioButton {
	i := &RadioButton{}
	initWidget(i, unsafe.Pointer(C.Fl_Radio_Button_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	i.setDeletionCallback(i.onDelete)
	return i
}

func (b *RadioButton) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Radio_Button_set_deletion_callback((*C.Fl_Radio_Button)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *RadioButton) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Radio_Button_handle((*C.Fl_Radio_Button)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *RadioButton) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Radio_Button_resize_callback((*C.Fl_Radio_Button)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *RadioButton) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Radio_Button_draw((*C.Fl_Radio_Button)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

type RadioRoundButton struct {
	Button
}

func NewRadioRoundButton(x, y, w, h int, text ...string) *RadioRoundButton {
	i := &RadioRoundButton{}
	initWidget(i, unsafe.Pointer(C.Fl_Radio_Round_Button_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	i.setDeletionCallback(i.onDelete)
	return i
}

func (b *RadioRoundButton) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Radio_Round_Button_set_deletion_callback((*C.Fl_Radio_Round_Button)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *RadioRoundButton) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Radio_Round_Button_handle((*C.Fl_Radio_Round_Button)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *RadioRoundButton) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Radio_Round_Button_resize_callback((*C.Fl_Radio_Round_Button)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *RadioRoundButton) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Radio_Round_Button_draw((*C.Fl_Radio_Round_Button)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

type ReturnButton struct {
	Button
}

func NewReturnButton(x, y, w, h int, text ...string) *ReturnButton {
	i := &ReturnButton{}
	initWidget(i, unsafe.Pointer(C.Fl_Return_Button_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	i.setDeletionCallback(i.onDelete)
	return i
}

func (b *ReturnButton) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Return_Button_set_deletion_callback((*C.Fl_Return_Button)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *ReturnButton) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Return_Button_handle((*C.Fl_Return_Button)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *ReturnButton) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Return_Button_resize_callback((*C.Fl_Return_Button)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *ReturnButton) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Return_Button_draw((*C.Fl_Return_Button)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

type ToggleButton struct {
	Button
}

func NewToggleButton(x, y, w, h int, text ...string) *ToggleButton {
	i := &ToggleButton{}
	initWidget(i, unsafe.Pointer(C.Fl_Toggle_Button_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	i.setDeletionCallback(i.onDelete)
	return i
}

func (b *ToggleButton) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Toggle_Button_set_deletion_callback((*C.Fl_Toggle_Button)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *ToggleButton) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Toggle_Button_handle((*C.Fl_Toggle_Button)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *ToggleButton) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Toggle_Button_resize_callback((*C.Fl_Toggle_Button)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *ToggleButton) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Toggle_Button_draw((*C.Fl_Toggle_Button)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}
