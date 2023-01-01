package fltk

/*
#include "cfltk/cfl_valuator.h"
#include "cfltk/cfl_enums.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type Slider struct {
	valuator
}

type SliderType uint8

var (
	VERT_SLIDER      = SliderType(C.Fl_SliderType_Vertical)
	HOR_SLIDER       = SliderType(C.Fl_SliderType_Horizontal)
	VERT_FILL_SLIDER = SliderType(C.Fl_SliderType_VerticalFill)
	HOR_FILL_SLIDER  = SliderType(C.Fl_SliderType_HorizontalFill)
	VERT_NICE_SLIDER = SliderType(C.Fl_SliderType_VerticalNice)
	HOR_NICE_SLIDER  = SliderType(C.Fl_SliderType_HorizontalNice)
)

func NewSlider(x, y, w, h int, text ...string) *Slider {
	s := &Slider{}
	initWidget(s, unsafe.Pointer(C.Fl_Slider_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	s.setDeletionCallback(s.onDelete)
	return s
}

func (b *Slider) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Slider_set_deletion_callback((*C.Fl_Slider)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *Slider) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Slider_handle((*C.Fl_Slider)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *Slider) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Slider_resize_callback((*C.Fl_Slider)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *Slider) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Slider_draw((*C.Fl_Slider)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

func (s *Slider) SetType(sliderType SliderType) {
	s.widget.SetType(uint8(sliderType))
}

type ValueSlider struct {
	Slider
}

func NewValueSlider(x, y, w, h int, text ...string) *ValueSlider {
	s := &ValueSlider{}
	initWidget(s, unsafe.Pointer(C.Fl_Value_Slider_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	s.setDeletionCallback(s.onDelete)
	return s
}

func (b *ValueSlider) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Value_Slider_set_deletion_callback((*C.Fl_Value_Slider)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *ValueSlider) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Value_Slider_handle((*C.Fl_Value_Slider)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *ValueSlider) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Value_Slider_resize_callback((*C.Fl_Value_Slider)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *ValueSlider) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Value_Slider_draw((*C.Fl_Value_Slider)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

func (s *ValueSlider) SetTextFont(font Font) {
	C.Fl_Value_Slider_set_text_font((*C.Fl_Value_Slider)(s.ptr()), C.int(font))
}

func (s *ValueSlider) SetTextSize(size int) {
	C.Fl_Value_Slider_set_text_size((*C.Fl_Value_Slider)(s.ptr()), C.int(size))
}
