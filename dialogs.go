package fltk

/*
#include <stdlib.h>
#include "include/cfltk/cfl_dialog.h"
*/
import "C"
import "unsafe"

func MessageBox(title, message string) {
	titleStr := C.CString(title)
	defer C.free(unsafe.Pointer(titleStr))
	messageStr := C.CString(message)
	defer C.free(unsafe.Pointer(messageStr))
	C.Fl_message_title(titleStr)
	C.Fl_message2(messageStr)
}

func ChoiceDialog(message string, options ...string) int {
	if len(options) == 0 || len(options) >= 3 {
		panic("Unsupported number of ChoiceDialog options")
	}
	messageStr := C.CString(message)
	defer C.free(unsafe.Pointer(messageStr))
	option0 := C.CString(options[0])
	defer C.free(unsafe.Pointer(option0))
	option1 := (*C.char)(nil)
	if len(options) > 1 {
		option1 = C.CString(options[1])
		defer C.free(unsafe.Pointer(option1))
	}
	option2 := (*C.char)(nil)
	if len(options) > 2 {
		option2 = C.CString(options[2])
		defer C.free(unsafe.Pointer(option2))
	}
	// option3 := (*C.char)(nil)
	// if len(options) > 2 {
	// 	option3 = C.CString(options[3])
	// 	defer C.free(unsafe.Pointer(option3))
	// }
	return int(C.Fl_choice2(messageStr, option0, option1, option2))
}
