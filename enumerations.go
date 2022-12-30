package fltk

/*
#include "cfltk/cfl_enums.h"
#include <stdint.h>
*/
import "C"

type Align uint

var (
	ALIGN_CENTER             = Align(C.Fl_Align_Center)
	ALIGN_TOP                = Align(C.Fl_Align_Top)
	ALIGN_BOTTOM             = Align(C.Fl_Align_Bottom)
	ALIGN_LEFT               = Align(C.Fl_Align_Left)
	ALIGN_RIGHT              = Align(C.Fl_Align_Right)
	ALIGN_INSIDE             = Align(C.Fl_Align_Inside)
	ALIGN_TEXT_OVER_IMAGE    = Align(C.Fl_Align_TextOverImage)
	ALIGN_IMAGE_OVER_TEXT    = Align(C.Fl_Align_ImageOverText)
	ALIGN_CLIP               = Align(C.Fl_Align_Clip)
	ALIGN_WRAP               = Align(C.Fl_Align_Wrap)
	ALIGN_IMAGE_NEXT_TO_TEXT = Align(C.Fl_Align_ImageNextToText)
	ALIGN_TEXT_NEXT_TO_IMAGE = Align(C.Fl_Align_TextNextToImage)
	ALIGN_IMAGE_BACKDROP     = Align(C.Fl_Align_ImageBackdrop)
	ALIGN_TOP_LEFT           = Align(C.Fl_Align_TopLeft)
	ALIGN_TOP_RIGHT          = Align(C.Fl_Align_TopRight)
	ALIGN_BOTTOM_LEFT        = Align(C.Fl_Align_BottomLeft)
	ALIGN_BOTTOM_RIGHT       = Align(C.Fl_Align_BottomRight)
	ALIGN_LEFT_TOP           = Align(C.Fl_Align_LeftTop)
	ALIGN_RIGHT_TOP          = Align(C.Fl_Align_RightTop)
	ALIGN_LEFT_BOTTOM        = Align(C.Fl_Align_LeftBottom)
	ALIGN_RIGHT_BOTTOM       = Align(C.Fl_Align_RightBottom)
	ALIGN_NOWRAP             = Align(C.Fl_Align_NoWrap)
	ALIGN_POSITION_MASK      = Align(C.Fl_Align_PositionMask)
	ALIGN_IMAGE_MASK         = Align(C.Fl_Align_ImageMask)
)

type BoxType int

const (
	NO_BOX                 = BoxType(0)
	FLAT_BOX               = BoxType(1)
	UP_BOX                 = BoxType(2)
	DOWN_BOX               = BoxType(3)
	UP_FRAME               = BoxType(4)
	DOWN_FRAME             = BoxType(5)
	THIN_UP_BOX            = BoxType(6)
	THIN_DOWN_BOX          = BoxType(7)
	THIN_UP_FRAME          = BoxType(8)
	THIN_DOWN_FRAME        = BoxType(9)
	ENGRAVED_BOX           = BoxType(10)
	EMBOSSED_BOX           = BoxType(11)
	ENGRAVED_FRAME         = BoxType(12)
	EMBOSSED_FRAME         = BoxType(13)
	BORDER_BOX             = BoxType(14)
	SHADOW_BOX             = BoxType(15)
	BORDER_FRAME           = BoxType(16)
	SHADOW_FRAME           = BoxType(17)
	ROUNDED_BOX            = BoxType(18)
	RSHADOW_BOX            = BoxType(19)
	ROUNDED_FRAME          = BoxType(20)
	RFLAT_BOX              = BoxType(21)
	ROUND_UP_BOX           = BoxType(22)
	ROUND_DOWN_BOX         = BoxType(23)
	DIAMOND_UP_BOX         = BoxType(24)
	DIAMOND_DOWN_BOX       = BoxType(25)
	OVAL_BOX               = BoxType(26)
	OSHADOW_BOX            = BoxType(27)
	OVAL_FRAME             = BoxType(28)
	OFLAT_BOX              = BoxType(29)
	PLASTIC_UP_BOX         = BoxType(30)
	PLASTIC_DOWN_BOX       = BoxType(31)
	PLASTIC_UP_FRAME       = BoxType(32)
	PLASTIC_DOWN_FRAME     = BoxType(33)
	PLASTIC_THIN_UP_BOX    = BoxType(34)
	PLASTIC_THIN_DOWN_BOX  = BoxType(35)
	PLASTIC_ROUND_UP_BOX   = BoxType(36)
	PLASTIC_ROUND_DOWN_BOX = BoxType(37)
	GTK_UP_BOX             = BoxType(38)
	GTK_DOWN_BOX           = BoxType(39)
	GTK_UP_FRAME           = BoxType(40)
	GTK_DOWN_FRAME         = BoxType(41)
	GTK_THIN_UP_BOX        = BoxType(42)
	GTK_THIN_DOWN_BOX      = BoxType(43)
	GTK_THIN_UP_FRAME      = BoxType(44)
	GTK_THIN_DOWN_FRAME    = BoxType(45)
	GTK_ROUND_UP_FRAME     = BoxType(46)
	GTK_ROUND_DOWN_FRAME   = BoxType(47)
	GLEAM_UP_BOX           = BoxType(48)
	GLEAM_DOWN_BOX         = BoxType(49)
	GLEAM_UP_FRAME         = BoxType(50)
	GLEAM_DOWN_FRAME       = BoxType(51)
	GLEAM_THIN_UP_BOX      = BoxType(52)
	GLEAM_THIN_DOWN_BOX    = BoxType(53)
	GLEAM_ROUND_UP_BOX     = BoxType(54)
	GLEAM_ROUND_DOWN_BOX   = BoxType(55)
	FREE_BOXTYPE           = BoxType(56)
)

type Font int

var (
	HELVETICA             = Font(C.Fl_Font_Helvetica)
	HELVETICA_BOLD        = Font(C.Fl_Font_HelveticaBold)
	HELVETICA_ITALIC      = Font(C.Fl_Font_HelveticaItalic)
	HELVETICA_BOLD_ITALIC = Font(C.Fl_Font_HelveticaBoldItalic)
	COURIER               = Font(C.Fl_Font_Courier)
	COURIER_BOLD          = Font(C.Fl_Font_CourierBold)
	COURIER_ITALIC        = Font(C.Fl_Font_CourierItalic)
	COURIER_BOLD_ITALIC   = Font(C.Fl_Font_CourierBoldItalic)
	TIMES                 = Font(C.Fl_Font_Times)
	TIMES_BOLD            = Font(C.Fl_Font_TimesBold)
	TIMES_ITALIC          = Font(C.Fl_Font_TimesItalic)
	TIMES_BOLD_ITALIC     = Font(C.Fl_Font_TimesBoldItalic)
	SYMBOL                = Font(C.Fl_Font_Symbol)
	SCREEN                = Font(C.Fl_Font_Screen)
	SCREEN_BOLD           = Font(C.Fl_Font_ScreenBold)
	ZAPF_DINGBATS         = Font(C.Fl_Font_Zapfdingbats)
)

type LabelType int

var (
	NORMAL_LABEL = LabelType(C.Fl_LabelType_Normal)
	NO_LABEL     = LabelType(C.Fl_LabelType_None)
)

type WrapMode int

const (
	WRAP_NONE      = WrapMode(0)
	WRAP_AT_COLUMN = WrapMode(1)
	WRAP_AT_PIXEL  = WrapMode(2)
	WRAP_AT_BOUNDS = WrapMode(3)
)

type Event int

var (
	NO_EVENT       = Event(C.Fl_Event_None)
	PUSH           = Event(C.Fl_Event_Push)
	DRAG           = Event(C.Fl_Event_Drag)
	RELEASE        = Event(C.Fl_Event_Released)
	MOVE           = Event(C.Fl_Event_Move)
	MOUSEWHEEL     = Event(C.Fl_Event_MouseWheel)
	ENTER          = Event(C.Fl_Event_Enter)
	LEAVE          = Event(C.Fl_Event_Leave)
	FOCUS          = Event(C.Fl_Event_Focus)
	UNFOCUS        = Event(C.Fl_Event_Unfocus)
	KEY            = Event(C.Fl_Event_KeyDown)
	KEYDOWN        = Event(C.Fl_Event_KeyDown)
	KEYUP          = Event(C.Fl_Event_KeyUp)
	SHORTCUT       = Event(C.Fl_Event_Shortcut)
	DEACTIVATE     = Event(C.Fl_Event_Deactivate)
	ACTIVATE       = Event(C.Fl_Event_Activate)
	HIDE           = Event(C.Fl_Event_Hide)
	SHOW           = Event(C.Fl_Event_Show)
	PASTE          = Event(C.Fl_Event_Paste)
	SELECTIONCLEAR = Event(C.Fl_Event_SelectionClear)
	DND_ENTER      = Event(C.Fl_Event_DndEnter)
	DND_DRAG       = Event(C.Fl_Event_DndDrag)
	DND_LEAVE      = Event(C.Fl_Event_DndLeave)
	DND_RELEASE    = Event(C.Fl_Event_DndRelease)
)

type CallbackCondition int

var (
	WhenNever           = CallbackCondition(C.Fl_CallbackTrigger_Never)
	WhenChanged         = CallbackCondition(C.Fl_CallbackTrigger_Changed)
	WhenNotChanged      = CallbackCondition(C.Fl_CallbackTrigger_NotChanged)
	WhenRelease         = CallbackCondition(C.Fl_CallbackTrigger_Release)
	WhenReleaseAlways   = CallbackCondition(C.Fl_CallbackTrigger_ReleaseAlways)
	WhenEnterKey        = CallbackCondition(C.Fl_CallbackTrigger_EnterKey)
	WhenEnterKeyAlways  = CallbackCondition(C.Fl_CallbackTrigger_EnterKeyAlways)
	WhenEnterKeyChanged = CallbackCondition(C.Fl_CallbackTrigger_EnterKeyChanged)
)

var (
	RGB         = int(C.Fl_Mode_Rgb)
	INDEX       = int(C.Fl_Mode_Index)
	DOUBLE      = int(C.Fl_Mode_Double)
	ACCUM       = int(C.Fl_Mode_Accum)
	ALPHA       = int(C.Fl_Mode_Alpha)
	DEPTH       = int(C.Fl_Mode_Depth)
	STENCIL     = int(C.Fl_Mode_Stencil)
	RGB8        = int(C.Fl_Mode_Rgb8)
	MULTISAMPLE = int(C.Fl_Mode_MultiSample)
	STEREO      = int(C.Fl_Mode_Stereo)
	FAKE_SINGLE = int(C.Fl_Mode_FakeSingle)
	OPENGL3     = int(C.Fl_Mode_Opengl3)
)

type Color uint

var (
	FOREGROUND_COLOR  = Color(C.Fl_Color_Foreground)
	BACKGROUND2_COLOR = Color(C.Fl_Color_Background2)
	INACTIVE_COLOR    = Color(C.Fl_Color_Inactive)
	SELECTION_COLOR   = Color(C.Fl_Color_Selection)
	GRAY0             = Color(C.Fl_Color_Gray0)
	DARK3             = Color(C.Fl_Color_Dark3)
	DARK2             = Color(C.Fl_Color_Dark2)
	DARK1             = Color(C.Fl_Color_Dark1)
	BACKGROUND_COLOR  = Color(C.Fl_Color_Background)
	LIGHT1            = Color(C.Fl_Color_Light1)
	LIGHT2            = Color(C.Fl_Color_Light2)
	LIGHT3            = Color(C.Fl_Color_Light3)
	BLACK             = Color(C.Fl_Color_Black)
	RED               = Color(C.Fl_Color_Red)
	GREEN             = Color(C.Fl_Color_Green)
	YELLOW            = Color(C.Fl_Color_Yellow)
	BLUE              = Color(C.Fl_Color_Blue)
	MAGENTA           = Color(C.Fl_Color_Magenta)
	CYAN              = Color(C.Fl_Color_Cyan)
	DARK_RED          = Color(C.Fl_Color_DarkRed)
	DARK_GREEN        = Color(C.Fl_Color_DarkGreen)
	DARK_YELLOW       = Color(C.Fl_Color_DarkYellow)
	DARK_BLUE         = Color(C.Fl_Color_DarkBlue)
	DARK_MAGENTA      = Color(C.Fl_Color_DarkMagenta)
	DARK_CYAN         = Color(C.Fl_Color_DarkCyan)
	WHITE             = Color(C.Fl_Color_White)
)

func ColorFromRgb(r, g, b uint8) Color {
	r1 := uint(r)
	g1 := uint(g)
	b1 := uint(b)
	return Color(((r1 & 0xff) << 24) + ((g1 & 0xff) << 16) + ((b1 & 0xff) << 8) + 0x00)
}

type LineStyle int

var (
	SOLID        = LineStyle(0)
	DASH         = LineStyle(1)
	Dot          = LineStyle(2)
	DASH_DOT     = LineStyle(3)
	DASH_DOT_DOT = LineStyle(4)
	CAP_FLAT     = LineStyle(100)
	CAP_ROUND    = LineStyle(200)
	CAP_SQUARE   = LineStyle(300)
	JOIN_MITER   = LineStyle(1000)
	JOIN_ROUND   = LineStyle(2000)
	JOIN_BEVEL   = LineStyle(3000)
)

type callbackMap struct {
	callbackMap map[uintptr]func()
	id          uintptr
}

func newCallbackMap() *callbackMap {
	return &callbackMap{
		callbackMap: make(map[uintptr]func()),
	}
}
func (m *callbackMap) register(fn func()) uintptr {
	m.id++
	m.callbackMap[m.id] = fn
	return m.id
}
func (m *callbackMap) unregister(id uintptr) {
	delete(m.callbackMap, id)
}
func (m *callbackMap) invoke(id uintptr) {
	if callback, ok := m.callbackMap[id]; ok && callback != nil {
		callback()
	}
}
func (m *callbackMap) isEmpty() bool {
	return len(m.callbackMap) == 0
}
func (m *callbackMap) size() int {
	return len(m.callbackMap)
}
func (m *callbackMap) clear() {
	for id := range m.callbackMap {
		delete(m.callbackMap, id)
	}
}

var globalCallbackMap = newCallbackMap()

//export _go_callbackHandler
func _go_callbackHandler(id uintptr) {
	globalCallbackMap.invoke(id)
}

type eventHandlerMap struct {
	eventHandlerMap map[uintptr]func(Event) bool
	id              uintptr
}

func newEventHandlerMap() *eventHandlerMap {
	return &eventHandlerMap{
		eventHandlerMap: make(map[uintptr]func(Event) bool),
	}
}
func (m *eventHandlerMap) register(fn func(Event) bool) uintptr {
	m.id++
	m.eventHandlerMap[m.id] = fn
	return m.id
}
func (m *eventHandlerMap) unregister(id uintptr) {
	delete(m.eventHandlerMap, id)
}
func (m *eventHandlerMap) invoke(id uintptr, event Event) bool {
	if handler, ok := m.eventHandlerMap[id]; ok && handler != nil {
		return handler(event)
	}
	return false
}
func (m *eventHandlerMap) isEmpty() bool {
	return len(m.eventHandlerMap) == 0
}
func (m *eventHandlerMap) size() int {
	return len(m.eventHandlerMap)
}
func (m *eventHandlerMap) clear() {
	for id := range m.eventHandlerMap {
		delete(m.eventHandlerMap, id)
	}
}

var globalEventHandlerMap = newEventHandlerMap()

//export _go_eventHandler
func _go_eventHandler(handlerId C.uintptr_t, event C.int) C.int {
	if globalEventHandlerMap.invoke(uintptr(handlerId), Event(event)) {
		return 1
	}
	return 0
}

var (
	ESCAPE = int(C.Fl_Key_Escape)
)
