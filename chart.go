package fltk

/*
#include "cfltk/cfl_misc.h"
#include "fltk.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

const (
	BAR_CHART        = iota /* Bar Chart variant */
	HORBAR_CHART     = iota /* Horizontal Bar Chart variant */
	LINE_CHART       = iota /* Line Chart variant */
	FILLED_CHART     = iota /* Fill Line Chart variant */
	SPIKE_CHART      = iota /* Spike Chart variant */
	PIE_CHART        = iota /* Pie Chart variant */
	SPECIALPIE_CHART = iota /* Special Pie Chart variant */
)

type Chart struct {
	widget
}

func NewChart(x, y, w, h int, text ...string) *Chart {
	c := &Chart{}
	initWidget(c, unsafe.Pointer(C.Fl_Chart_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	c.setDeletionCallback(c.onDelete)
	return c
}

func (b *Chart) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Chart_set_deletion_callback((*C.Fl_Chart)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *Chart) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Chart_handle((*C.Fl_Chart)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *Chart) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Chart_resize_callback((*C.Fl_Chart)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *Chart) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Chart_draw((*C.Fl_Chart)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

func (c *Chart) Clear() {
	C.Fl_Chart_clear((*C.Fl_Chart)(c.ptr()))
}

//Add the data value val with optional label text and color col to the chart.
//When color not needed just pass zero
func (c *Chart) Add(val float64, col Color, text ...string) {
	labelStr := cStringOpt(text)
	defer C.free(unsafe.Pointer(labelStr))

	C.Fl_Chart_add((*C.Fl_Chart)(c.ptr()), C.double(val), labelStr, C.uint(col))
}

// Insert inserts a data value val at the given position ind.
//
//Position 1 is the first data value.
func (c *Chart) Insert(index int, val float64, col Color, text ...string) {
	labelStr := cStringOpt(text)
	defer C.free(unsafe.Pointer(labelStr))

	C.Fl_Chart_insert((*C.Fl_Chart)(c.ptr()), C.int(index), C.double(val), labelStr, C.uint(col))
}

// Replace replace a data value val at the given position index.
//
//Position 1 is the first data value.
func (c *Chart) Replace(index int, val float64, col Color, text ...string) {
	labelStr := cStringOpt(text)
	defer C.free(unsafe.Pointer(labelStr))

	C.Fl_Chart_replace((*C.Fl_Chart)(c.ptr()), C.int(index), C.double(val), labelStr, C.uint(col))
}

// Bounds gets the lower and upper bounds of the chart values.
func (c *Chart) Bounds() (float64, float64) {
	var a, b float64
	C.Fl_Chart_bounds((*C.Fl_Chart)(c.ptr()), (*C.double)(unsafe.Pointer(&a)), (*C.double)(unsafe.Pointer(&b)))

	return a, b
}

// SetBounds sets the lower and upper bounds of the chart values.
func (c *Chart) SetBounds(a, b float64) {
	C.Fl_Chart_set_bounds((*C.Fl_Chart)(c.ptr()), C.double(a), C.double(b))
}

// Size returns the number of data values in the chart.
func (c *Chart) Size() int {
	return int(C.Fl_Chart_size((*C.Fl_Chart)(c.ptr())))
}

// func (c *Chart) SetSize(W, H int) {
// 	C.Fl_Chart_set_size((*C.Fl_Chart)(c.ptr()), C.int(W), C.int(H))
// }

// MaxSize gets the maximum number of data values for a chart.
func (c *Chart) MaxSize() int {
	return int(C.Fl_Chart_maxsize((*C.Fl_Chart)(c.ptr())))
}

// SetMaxSize sets the maximum number of data values for a chart.
//
//If you do not call this method then the chart will be allowed to grow to any size depending on available memory.
func (c *Chart) SetMaxSize(m int) {
	C.Fl_Chart_set_maxsize((*C.Fl_Chart)(c.ptr()), C.int(m))
}

// TextFont gets the chart's text font.
func (c *Chart) TextFont() Font {
	return Font(C.Fl_Chart_text_font((*C.Fl_Chart)(c.ptr())))
}

// SetTextFont sets the chart's text font to font.
func (c *Chart) SetTextFont(font Font) {
	C.Fl_Chart_set_text_font((*C.Fl_Chart)(c.ptr()), C.int(font))
}

// TextSize gets the chart's text size.
func (c *Chart) TextSize() int {
	return int(C.Fl_Chart_text_size((*C.Fl_Chart)(c.ptr())))
}

// SetTextSize sets the chart's text font to size.
func (c *Chart) SetTextSize(size int) {
	C.Fl_Chart_set_text_size((*C.Fl_Chart)(c.ptr()), C.int(size))
}

// TextColor gets the chart's text color.
func (c *Chart) TextColor() Color {
	return Color(C.Fl_Chart_text_color((*C.Fl_Chart)(c.ptr())))
}

// SetTextColor sets the chart's text color to color
func (c *Chart) SetTextColor(color Color) {
	C.Fl_Chart_set_text_color((*C.Fl_Chart)(c.ptr()), C.uint(color))
}

// Autosize gets whether the chart will automatically adjust the bounds of the chart.
func (c *Chart) Autosize() bool {
	return C.Fl_Chart_is_autosize((*C.Fl_Chart)(c.ptr())) != 0
}

// SetAutosize sets whether the chart will automatically adjust the bounds of the chart.
func (c *Chart) SetAutosize(flag bool) {
	var f uint8
	if flag {
		f = 1
	}
	C.Fl_Chart_make_autosize((*C.Fl_Chart)(c.ptr()), C.int(f))
}
