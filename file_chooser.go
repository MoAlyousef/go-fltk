package fltk

/*
#include <stdlib.h>
#include "include/cfltk/cfl_dialog.h"
#include "include/cfltk/cfl_enums.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type FileChooser struct {
	cPtr                     *C.Fl_File_Chooser
	pathname, pattern, title *C.char
	callbackId               uintptr
}

var ErrFileChooserDestroyed = errors.New("file chooser is destroyed")

type FileChooserType int

var (
	FileChooser_SINGLE    = FileChooserType(C.Fl_FileChooserType_Single)
	FileChooser_MULTI     = FileChooserType(C.Fl_FileChooserType_Multi)
	FileChooser_CREATE    = FileChooserType(C.Fl_FileChooserType_Create)
	FileChooser_DIRECTORY = FileChooserType(C.Fl_FileChooserType_Directory)
)

func NewFileChooser(pathname, pattern string, fctype FileChooserType, title string) *FileChooser {
	c := &FileChooser{}
	c.pathname = C.CString(pathname)
	c.pattern = C.CString(pattern)
	c.title = C.CString(title)
	c.cPtr = C.Fl_File_Chooser_new(c.pathname, c.pattern, C.int(fctype), c.title)
	return c
}

func (c *FileChooser) ptr() *C.Fl_File_Chooser {
	if c.cPtr == nil {
		panic(ErrFileChooserDestroyed)
	}
	return c.cPtr
}
func (c *FileChooser) Destroy() {
	if c.pathname != nil {
		C.free(unsafe.Pointer(c.pathname))
	}
	c.pathname = nil
	if c.pattern != nil {
		C.free(unsafe.Pointer(c.pattern))
	}
	c.pattern = nil
	if c.title != nil {
		C.free(unsafe.Pointer(c.title))
	}
	c.title = nil
	if c.callbackId > 0 {
		globalCallbackMap.unregister(c.callbackId)
	}
	C.Fl_File_Chooser_delete(c.ptr())
	c.cPtr = nil
}

func (c *FileChooser) SetCallback(callback func()) {
	if c.callbackId > 0 {
		globalCallbackMap.unregister(c.callbackId)
	}
	c.callbackId = globalCallbackMap.register(callback)
	// C.Fl_File_Chooser_set_callback(c.ptr(), C.uintptr_t(c.callbackId))
}
func (c *FileChooser) Show() {
	C.Fl_File_Chooser_show((*C.Fl_File_Chooser)(c.ptr()))
}
func (c *FileChooser) Popup() {
	C.Fl_File_Chooser_show((*C.Fl_File_Chooser)(c.ptr()))
}
func (c *FileChooser) Shown() bool {
	return C.Fl_File_Chooser_shown(c.ptr()) != 0
}
func (c *FileChooser) SetPreview(enable bool) {
	if enable {
		C.Fl_File_Chooser_preview(c.ptr())
	} else {
		// C.Fl_File_Chooser_preview(c.ptr(), 0)
	}
}
func (c *FileChooser) Selection() []string {
	count := int(C.Fl_File_Chooser_count(c.ptr()))
	var selection []string
	for i := 1; i <= count; i++ {
		value := C.GoString(C.Fl_File_Chooser_value(c.ptr(), C.int(i)))
		selection = append(selection, value)
	}
	return selection
}

func ChooseFile(message, pattern, initialFilename string, relative bool) (string, bool) {
	var rel int
	if relative {
		rel = 1
	}
	messageStr := C.CString(message)
	defer C.free(unsafe.Pointer(messageStr))
	patternStr := C.CString(pattern)
	defer C.free(unsafe.Pointer(patternStr))
	initialFilenameStr := C.CString(initialFilename)
	defer C.free(unsafe.Pointer(initialFilenameStr))
	res := C.Fl_file_chooser(messageStr, patternStr, initialFilenameStr, C.int(rel))
	if res == nil {
		return "", false
	}
	return C.GoString(res), true
}
