package fltk

/*
#include "include/cfltk/cfl_enums.h"
#include "include/cfltk/cfl.h"
*/
import "C"

type MouseButton int

var (
	LeftMouse   = MouseButton(C.Fl_Shortcut_Button1)
	MiddleMouse = MouseButton(C.Fl_Shortcut_Button2)
	RightMouse  = MouseButton(C.Fl_Shortcut_Button3)
)

func EventType() Event {
	return Event(C.Fl_event())
}

func EventButton() MouseButton {
	return MouseButton(C.Fl_event_button())
}

func EventButton1() bool {
	return C.Fl_event_button() != 0
}
func EventX() int {
	return int(C.Fl_event_x())
}
func EventY() int {
	return int(C.Fl_event_y())
}
func EventXRoot() int {
	return int(C.Fl_event_x_root())
}
func EventYRoot() int {
	return int(C.Fl_event_y_root())
}
func EventDX() int {
	return int(C.Fl_event_dx())
}
func EventDY() int {
	return int(C.Fl_event_dy())
}
func EventKey() int {
	return int(C.Fl_event_key())
}
func EventIsClick() bool {
	return C.Fl_event_is_click() != 0
}
func EventText() string {
	return C.GoString(C.Fl_event_text())
}

var (
	SHIFT       = int(C.Fl_Shortcut_Shift)
	CAPS_LOCK   = int(C.Fl_Shortcut_CapsLock)
	CTRL        = int(C.Fl_Shortcut_Ctrl)
	ALT         = int(C.Fl_Shortcut_Alt)
	NUM_LOCK    = int(C.Fl_Key_NumLock)
	META        = int(C.Fl_Shortcut_Meta)
	SCROLL_LOCK = int(C.Fl_Key_ScrollLock)
	BUTTON1     = int(C.Fl_Shortcut_Button1)
	BUTTON2     = int(C.Fl_Shortcut_Button2)
	BUTTON3     = int(C.Fl_Shortcut_Button3)
)

func EventState() int {
	return int(C.Fl_event_state())
}
