package fltk

/*
#include <stdlib.h>
#include "cfltk/cfl_menu.h"
#include "cfltk/cfl_enums.h"
*/
import "C"
import "unsafe"

type menu struct {
	widget
	deletionHandlerId uintptr
	itemCallbacks     []uintptr
}

func (m *menu) init() {
	if m.deletionHandlerId > 0 {
		panic("menu already initialized")
	}
	// m.setDeletionHandler(m.onDelete)
}
func (m *menu) onDelete() {
	if m.deletionHandlerId > 0 {
		globalCallbackMap.unregister(m.deletionHandlerId)
	}
	m.deletionHandlerId = 0
	for _, itemCallbackId := range m.itemCallbacks {
		globalCallbackMap.unregister(itemCallbackId)
	}
	m.itemCallbacks = m.itemCallbacks[:0]
}
func (m *menu) Destroy() {
	for _, itemCallbackId := range m.itemCallbacks {
		globalCallbackMap.unregister(itemCallbackId)
	}
	m.itemCallbacks = m.itemCallbacks[:0]
	m.widget.Destroy()
}
func (m *menu) Add(label string, callback func()) int {
	callbackId := globalCallbackMap.register(callback)
	m.itemCallbacks = append(m.itemCallbacks, callbackId)
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	// return int(C.Fl_Menu_Bar_add((*C.Fl_Menu_Bar)(m.ptr()), labelStr, 0, C.int(callbackId), 0))
	return 0
}
func (m *menu) AddEx(label string, shortcut int, callback func(), flags int) int {
	callbackId := globalCallbackMap.register(callback)
	m.itemCallbacks = append(m.itemCallbacks, callbackId)
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	// return int(C.Fl_Menu_Bar_add((*C.Fl_Menu_Bar)(m.ptr()), labelStr, C.int(shortcut), C.int(callbackId), C.int(flags)))
	return 0
}

func (m *menu) SetValue(value int) {
	C.Fl_Menu_Bar_set_value((*C.Fl_Menu_Bar)(m.ptr()), C.int(value))
}
func (m *menu) Value() int {
	return int(C.Fl_Menu_Bar_value((*C.Fl_Menu_Bar)(m.ptr())))
}
func (m *menu) Size() int {
	return int(C.Fl_Menu_Bar_size((*C.Fl_Menu_Bar)(m.ptr())))
}

type MenuButton struct {
	menu
}

func NewMenuButton(x, y, w, h int, text ...string) *MenuButton {
	m := &MenuButton{}
	initWidget(m, unsafe.Pointer(C.Fl_Menu_Button_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	m.menu.init()
	return m
}

type MenuType int

var (
	POPUP1   = MenuType(C.Popup1)
	POPUP2   = MenuType(C.Popup2)
	POPUP12  = MenuType(C.Popup12)
	POPUP3   = MenuType(C.Popup3)
	POPUP13  = MenuType(C.Popup13)
	POPUP23  = MenuType(C.Popup23)
	POPUP123 = MenuType(C.Popup123)
)

var (
	MENU_INACTIVE  = int(C.Fl_MenuFlag_Inactive)
	MENU_TOGGLE    = int(C.Fl_MenuFlag_Toggle)
	MENU_VALUE     = int(C.Fl_MenuFlag_Value)
	MENU_RADIO     = int(C.Fl_MenuFlag_Radio)
	MENU_INVISIBLE = int(C.Fl_MenuFlag_Invisible)
	SUBMENU        = int(C.Fl_MenuFlag_Submenu)
	MENU_DIVIDER   = int(C.Fl_MenuFlag_MenuDivider)
)

func (m *MenuButton) SetType(menuType MenuType) {
	C.Fl_Menu_Button_set_type((*C.Fl_Menu_Button)(m.ptr()), C.int(menuType))
}
func (m *MenuButton) Popup() {
	C.Fl_Menu_Button_popup((*C.Fl_Menu_Button)(m.ptr()))
}
func (m *MenuButton) Destroy() {
	m.menu.Destroy()
}

type MenuBar struct {
	menu
}

func NewMenuBar(x, y, w, h int, text ...string) *MenuBar {
	m := &MenuBar{}
	initWidget(m, unsafe.Pointer(C.Fl_Menu_Bar_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	m.menu.init()
	return m
}
