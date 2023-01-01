package fltk

/*
#include "cfltk/cfl_table.h"
#include "cfltk/cfl_enums.h"
#include "fltk.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type table struct {
	Group
}

func (t *table) SetRowCount(rowCount int) {
	C.Fl_Table_set_rows((*C.Fl_Table)(t.ptr()), C.int(rowCount))
}

func (t *table) SetRowHeight(row, height int) {
	C.Fl_Table_set_row_height((*C.Fl_Table)(t.ptr()), C.int(row), C.int(height))
}

func (t *table) SetRowHeightAll(height int) {
	C.Fl_Table_set_row_height_all((*C.Fl_Table)(t.ptr()), C.int(height))
}

func (t *table) EnableRowHeaders() {
	C.Fl_Table_set_row_header((*C.Fl_Table)(t.ptr()), 1)
}

func (t *table) DisableRowHeaders() {
	C.Fl_Table_set_row_header((*C.Fl_Table)(t.ptr()), 0)
}

func (t *table) AllowRowResizing() {
	C.Fl_Table_set_row_resize((*C.Fl_Table)(t.ptr()), 1)
}

func (t *table) DisallowRowResizing() {
	C.Fl_Table_set_row_resize((*C.Fl_Table)(t.ptr()), 0)
}

func (t *table) SetColumnCount(columnCount int) {
	C.Fl_Table_set_cols((*C.Fl_Table)(t.ptr()), C.int(columnCount))
}

func (t *table) SetColumnWidth(column, width int) {
	C.Fl_Table_set_col_width((*C.Fl_Table)(t.ptr()), C.int(column), C.int(width))
}

func (t *table) SetColumnWidthAll(width int) {
	C.Fl_Table_set_col_width_all((*C.Fl_Table)(t.ptr()), C.int(width))
}

func (t *table) EnableColumnHeaders() {
	C.Fl_Table_set_col_header((*C.Fl_Table)(t.ptr()), 1)
}

func (t *table) DisableColumnHeaders() {
	C.Fl_Table_set_col_header((*C.Fl_Table)(t.ptr()), 0)
}

func (t *table) AllowColumnResizing() {
	C.Fl_Table_set_col_resize((*C.Fl_Table)(t.ptr()), 1)
}

func (t *table) DisallowColumnResizing() {
	C.Fl_Table_set_col_resize((*C.Fl_Table)(t.ptr()), 0)
}

func (t *table) CallbackRow() int {
	return int(C.Fl_Table_callback_row((*C.Fl_Table)(t.ptr())))
}

func (t *table) CallbackContext() TableContext {
	return TableContext(C.Fl_Table_callback_context((*C.Fl_Table)(t.ptr())))
}

func (t *table) Selection() (int, int, int, int) {
	var top, left, bottom, right C.int
	C.Fl_Table_get_selection((*C.Fl_Table)(t.ptr()), &top, &left, &bottom, &right)
	return int(top), int(left), int(bottom), int(right)
}

func (t *table) VisibleCells() (int, int, int, int) {
	var top, bottom, left, right C.int
	C.Fl_Table_visible_cells((*C.Fl_Table)(t.ptr()), &top, &bottom, &left, &right)
	return int(top), int(left), int(bottom), int(right)
}

func (t *table) SetTopRow(row int) {
	C.Fl_Table_set_top_row((*C.Fl_Table)(t.ptr()), C.int(row))
}

func (t *table) ScrollbarSize() int {
	return int(C.Fl_Table_scrollbar_size((*C.Fl_Table)(t.ptr())))
}

func (t *table) SetScrollbarSize(size int) {
	C.Fl_Table_set_scrollbar_size((*C.Fl_Table)(t.ptr()), C.int(size))
}

type TableRow struct {
	table
	deletionHandlerId  uintptr
	drawCellCallbackId int
}

type tableCallbackMap struct {
	callbackMap map[int]func(TableContext, int, int, int, int, int, int)
	id          int
}

func newTableCallbackMap() *tableCallbackMap {
	return &tableCallbackMap{
		callbackMap: make(map[int]func(TableContext, int, int, int, int, int, int)),
	}
}

func (m *tableCallbackMap) register(fn func(TableContext, int, int, int, int, int, int)) int {
	m.id++
	m.callbackMap[m.id] = fn
	return m.id
}

func (m *tableCallbackMap) unregister(id int) {
	delete(m.callbackMap, id)
}

func (m *tableCallbackMap) invoke(id int, context TableContext, r, c, x, y, w, h int) {
	if id == 0 {
		return
	}
	if callback, ok := m.callbackMap[id]; ok && callback != nil {
		callback(context, r, c, x, y, w, h)
	}
}

func (m *tableCallbackMap) isEmpty() bool {
	return len(m.callbackMap) == 0
}

func (m *tableCallbackMap) size() int {
	return len(m.callbackMap)
}

func (m *tableCallbackMap) clear() {
	for id := range m.callbackMap {
		delete(m.callbackMap, id)
	}
}

var globalTableCallbackMap = newTableCallbackMap()

type TableContext int

var (
	ContextNone      = TableContext(C.Fl_TableContext_None)
	ContextStartPage = TableContext(C.Fl_TableContext_StartPage)
	ContextEndPage   = TableContext(C.Fl_TableContext_EndPage)
	ContextRowHeader = TableContext(C.Fl_TableContext_RowHeader)
	ContextColHeader = TableContext(C.Fl_TableContext_ColHeader)
	ContextCell      = TableContext(C.Fl_TableContext_Cell)
	ContextTable     = TableContext(C.Fl_TableContext_Table)
	ContextRCResize  = TableContext(C.Fl_TableContext_RcResize)
)

func NewTableRow(x, y, w, h int, text ...string) *TableRow {
	t := &TableRow{}
	ptr := C.Fl_Table_Row_new(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))
	initWidget(t, unsafe.Pointer(ptr))
	t.setDeletionCallback(t.onDelete)
	return t
}

func (b *TableRow) setDeletionCallback(handler func()) {
	b.deletionHandlerId = globalCallbackMap.register(handler)
	C.Fl_Table_Row_set_deletion_callback((*C.Fl_Table_Row)(b.ptr()), (*[0]byte)(C.go_deleter), unsafe.Pointer(b.deletionHandlerId))
}

func (w *TableRow) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	C.Fl_Table_Row_handle((*C.Fl_Table_Row)(w.ptr()), (C.custom_handler_callback)(C.event_handler), unsafe.Pointer(w.eventHandlerId))
}

func (w *TableRow) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.Fl_Table_Row_resize_callback((*C.Fl_Table_Row)(w.ptr()), (*[0]byte)(C.resize_handler), unsafe.Pointer(w.resizeHandlerId))
}

func (w *TableRow) SetDrawHandler(handler func()) {
	if w.drawHandlerId > 0 {
		globalCallbackMap.unregister(w.drawHandlerId)
	}
	w.drawHandlerId = globalCallbackMap.register(handler)
	C.Fl_Table_Row_draw((*C.Fl_Table_Row)(w.ptr()), (C.custom_draw_callback)(C.callback_handler), unsafe.Pointer(w.drawHandlerId))
}

func (t *TableRow) onDelete() {
	t.getWidget().onDelete()
	if t.deletionHandlerId > 0 {
		globalCallbackMap.unregister(t.deletionHandlerId)
	}
	t.deletionHandlerId = 0
	if t.drawCellCallbackId > 0 {
		globalTableCallbackMap.unregister(t.drawCellCallbackId)
	}
	t.drawCellCallbackId = 0
}

func (t *TableRow) Destroy() {
	if t.drawCellCallbackId > 0 {
		globalTableCallbackMap.unregister(t.drawCellCallbackId)
	}
	t.drawCellCallbackId = 0
	t.table.Destroy()
}

func (t *TableRow) IsRowSelected(row int) bool {
	return C.Fl_Table_Row_row_selected((*C.Fl_Table_Row)(t.ptr()), C.int(row)) != 0
}

func (t *TableRow) SetDrawCellCallback(callback func(TableContext, int, int, int, int, int, int)) {
	if t.drawCellCallbackId > 0 {
		globalTableCallbackMap.unregister(t.drawCellCallbackId)
	}
	t.drawCellCallbackId = globalTableCallbackMap.register(callback)
	// C.Fl_Table_Row_set_draw_cell_data((*C.Fl_Table_Row)(t.ptr()), C.int(t.drawCellCallbackId))
}

type SelectionFlag int

var (
	Deselect        = SelectionFlag(C.Fl_TreeItemSelect_Deselect)
	Select          = SelectionFlag(C.Fl_TreeItemSelect_Select)
	ToggleSelection = SelectionFlag(C.Fl_TreeItemSelect_Toggle)
)

func (t *TableRow) SelectAllRows(flag SelectionFlag) {
	C.Fl_Table_Row_select_all_rows((*C.Fl_Table_Row)(t.ptr()), C.int(flag))
}

func (t *TableRow) SelectRow(row int, flag SelectionFlag) {
	C.Fl_Table_Row_select_row((*C.Fl_Table_Row)(t.ptr()), C.int(row), C.int(flag))
}

func (t *TableRow) FindCell(ctx TableContext, row int, col int) (int, int, int, int, error) {
	var x, y, w, h C.int
	ret := C.Fl_Table_Row_find_cell((*C.Fl_Table_Row)(t.ptr()), C.int(ctx), C.int(row), C.int(col), &x, &y, &w, &h)
	err := errors.New("no cell was found")
	if ret == 0 {
		err = nil
	}
	return int(x), int(y), int(w), int(h), err
}

type RowSelectMode int

var (
	SelectNone   = RowSelectMode(C.Fl_TableRowSelectMode_None)
	SelectSingle = RowSelectMode(C.Fl_TableRowSelectMode_Single)
	SelectMulti  = RowSelectMode(C.Fl_TableRowSelectMode_Multi)
)

func (t *TableRow) SetType(tableType RowSelectMode) {
	C.Fl_Table_Row_set_type((*C.Fl_Table_Row)(t.ptr()), C.int(tableType))
}

//export _go_drawTableHandler
func _go_drawTableHandler(id, context, r, c, x, y, w, h C.int) {
	globalTableCallbackMap.invoke(int(id), TableContext(context), int(r), int(c), int(x), int(y), int(w), int(h))
}
