package fltk

/*
#include <stdlib.h>
#include "cfltk/cfl_draw.h"
*/
import "C"
import "unsafe"

func SetDrawColor(color Color) {
	C.Fl_set_color_int(C.uint(color))
}

func Draw(text string, x, y, w, h int, align Align) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	C.Fl_draw5(textStr, C.int(x), C.int(y), C.int(w), C.int(h), C.int(align), nil, 0)
}

func DrawBox(boxType BoxType, x, y, w, h int, color Color) {
	C.Fl_draw_box(
		C.int(boxType), C.int(x), C.int(y), C.int(w), C.int(h), C.uint(color))
}

func SetDrawFont(font Font, size int) {
	C.Fl_set_draw_font(C.int(font), C.int(size))
}

func PushClip(x, y, w, h int) {
	C.Fl_push_clip(C.int(x), C.int(y), C.int(w), C.int(h))
}

func PushNoClip() {
	C.Fl_push_no_clip()
}

func PopClip() {
	C.Fl_pop_clip()
}

func DrawPoint(x, y int) {
	C.Fl_point(C.int(x), C.int(y))
}

func SetLineStyle(style LineStyle, width int) {
	C.Fl_line_style(C.int(style), C.int(width), nil)
}

func DrawRect(x, y, w, h int) {
	C.Fl_rect(C.int(x), C.int(y), C.int(w), C.int(h))
}

func DrawFocusRect(x, y, w, h int) {
	C.Fl_focus_rect(C.int(x), C.int(y), C.int(w), C.int(h))
}

func DrawRectWithColor(x, y, w, h int, col Color) {
	C.Fl_rect_with_color(C.int(x), C.int(y), C.int(w), C.int(h), C.uint(col))
}

func DrawRectf(x, y, w, h int) {
	C.Fl_rectf(C.int(x), C.int(y), C.int(w), C.int(h))
}

func DrawRectfWithColor(x, y, w, h int, col Color) {
	C.Fl_rectf_with_color(C.int(x), C.int(y), C.int(w), C.int(h), C.uint(col))
}

func DrawLine(x, y, x1, y1 int) {
	C.Fl_line(C.int(x), C.int(y), C.int(x1), C.int(y1))
}

func DrawLine2(x, y, x1, y1, x2, y2 int) {
	C.Fl_line2(C.int(x), C.int(y), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

func DrawLoop(x, y, x1, y1, x2, y2 int) {
	C.Fl_loop(C.int(x), C.int(y), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

func DrawLoop2(x, y, x1, y1, x2, y2, x3, y3 int) {
	C.Fl_loop2(C.int(x), C.int(y), C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(x3), C.int(y3))
}

func DrawPolygon(x, y, x1, y1, x2, y2 int) {
	C.Fl_polygon(C.int(x), C.int(y), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

func DrawPolygon2(x, y, x1, y1, x2, y2, x3, y3 int) {
	C.Fl_polygon2(C.int(x), C.int(y), C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(x3), C.int(y3))
}

func DrawXyLine(x, y, x1 int) {
	C.Fl_xyline(C.int(x), C.int(y), C.int(x1))
}

func DrawXyLine2(x, y, x1, y2 int) {
	C.Fl_xyline2(C.int(x), C.int(y), C.int(x1), C.int(y2))
}

func DrawXyLine3(x, y, x1, y2, x3 int) {
	C.Fl_xyline3(C.int(x), C.int(y), C.int(x1), C.int(y2), C.int(x3))
}

func DrawYxLine(x, y, y1 int) {
	C.Fl_yxline(C.int(x), C.int(y), C.int(y1))
}

func DrawYxLine2(x, y, y1, x2 int) {
	C.Fl_yxline2(C.int(x), C.int(y), C.int(y1), C.int(x2))
}

func DrawYxLine3(x, y, y1, x2, y3 int) {
	C.Fl_yxline3(C.int(x), C.int(y), C.int(y1), C.int(x2), C.int(y3))
}

func DrawArc(x, y, w, h int, a1, a2 float64) {
	C.Fl_arc(C.int(x), C.int(y), C.int(w), C.int(h), C.double(a1), C.double(a2))
}

func DrawPie(x, y, w, h int, a1, a2 float64) {
	C.Fl_pie(C.int(x), C.int(y), C.int(w), C.int(h), C.double(a1), C.double(a2))
}

func DrawArc2(x, y, r, start, end float64) {
	C.Fl_arc2(C.double(x), C.double(y), C.double(r), C.double(start), C.double(end))
}

func DrawCirlce(x, y, r float64) {
	C.Fl_circle(C.double(x), C.double(y), C.double(r))
}

// returns the dx, dy, w, h of the string
func TextExtents(text string) (int, int, int, int) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	dx := C.int(0)
	dy := C.int(0)
	w := C.int(0)
	h := C.int(0)
	C.Fl_text_extents(textStr, &dx, &dy, &w, &h)
	return int(dx), int(dy), int(w), int(h)
}

func DrawTextAngled(text string, x, y, angle int) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	C.Fl_draw2(C.int(angle), textStr, C.int(x), C.int(y))
}

func DrawRtl(text string, x, y int) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	C.Fl_rtl_draw(textStr, C.int(len(text)), C.int(x), C.int(y))
}

func MeasureText(text string, draw_symbols bool) (int, int) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	x := C.int(0)
	y := C.int(0)
	d := 0
	if draw_symbols {
		d = 1
	}
	C.Fl_measure(textStr, &x, &y, C.int(d))
	return int(x), int(y)
}

func DrawCheck(x, y, w, h int, col Color) {
	C.Fl_draw_check(C.int(x), C.int(y), C.int(w), C.int(h), C.uint(col))
}

type Offscreen struct {
	oPtr unsafe.Pointer
}

func NewOffscreen(w, h int) *Offscreen {
	o := &Offscreen{
		oPtr: C.Fl_create_offscreen(C.int(w), C.int(h)),
	}
	return o
}

func (offs *Offscreen) Begin() {
	C.Fl_begin_offscreen(offs.oPtr)
}

func (offs *Offscreen) End() {
	C.Fl_end_offscreen()
}

func (offs *Offscreen) Rescale() {
	C.Fl_rescale_offscreen(&offs.oPtr)
}

func (offs *Offscreen) Delete() {
	C.Fl_delete_offscreen(offs.oPtr)
}

func (offs *Offscreen) IsValid() bool {
	return offs.oPtr != nil
}

func (offs *Offscreen) Copy(x, y, w, h, srcx, srcy int) {
	C.Fl_copy_offscreen(C.int(x), C.int(y), C.int(w), C.int(h), offs.oPtr, C.int(srcx), C.int(srcy))
}
