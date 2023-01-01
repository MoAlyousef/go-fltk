package fltk

/*
#include <stdlib.h>
#include "cfltk/cfl_menu.h"
#include "cfltk/cfl_enums.h"
#include "fltk.h"
*/
import "C"
import "unsafe"

type menu struct {
	widget
	deletionHandlerId uintptr
	itemCallbacks     []uintptr
}

func (m *menu) onDelete() {
	m.getWidget().onDelete()
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
	return int(C.Fl_Choice_add((*C.Fl_Choice)(m.ptr()), labelStr, C.int(0), (*C.Fl_Callback)(C.callback_handler),unsafe.Pointer(callbackId), C.int(0)))
}

func (m *menu) AddEx(label string, shortcut int, callback func(), flags int) int {
	callbackId := globalCallbackMap.register(callback)
	m.itemCallbacks = append(m.itemCallbacks, callbackId)
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	return int(C.Fl_Choice_add((*C.Fl_Choice)(m.ptr()), labelStr, C.int(shortcut), (*C.Fl_Callback)(C.callback_handler), unsafe.Pointer(callbackId), C.int(flags)))
}

func (m *menu) SetValue(value int) {
	C.Fl_Choice_set_value((*C.Fl_Choice)(m.ptr()), C.int(value))
}

func (m *menu) Value() int {
	return int(C.Fl_Choice_value((*C.Fl_Choice)(m.ptr())))
}

func (m *menu) Size() int {
	return int(C.Fl_Choice_size((*C.Fl_Choice)(m.ptr())))
}

type MenuButton struct {
	menu
}

func NewMenuButton(x, y, w, h int, text ...string) *MenuButton {
	m := &MenuButton{}
	ptr := unsafe.Pointer(C.Fl_Menu_Button_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text)))
	initWidget(m, ptr)
	m.setDeletionCallback(m.onDelete)
	return m
}

func (b *MenuButton) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Menu_Button_set_deletion_callback((*C.Fl_Menu_Button)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
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
	ptr := C.Fl_Menu_Bar_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))
	initWidget(m, unsafe.Pointer(ptr))
	m.setDeletionCallback(m.onDelete)
	return m
}

func (b *MenuBar) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Menu_Bar_set_deletion_callback((*C.Fl_Menu_Bar)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}