#include "fltk.h"
#include <stdint.h>

extern void _go_timeoutHandler(uintptr_t);

void my_timeout_handler(void* data) {
	_go_timeoutHandler((uintptr_t)data);
}

extern void _go_awakeHandler(uintptr_t);

void my_awake_handler(void* data) {
	_go_awakeHandler((uintptr_t)data);
}

extern void _go_callbackHandler(uintptr_t);

void callback_handler(Fl_Widget *w, void* data) {
	_go_callbackHandler((uintptr_t)data);
}

extern int _go_eventHandler(uintptr_t, int ev);

int event_handler(Fl_Widget *w, int ev, void* data) {
	return _go_eventHandler((uintptr_t)data, ev);
}

void resize_handler(Fl_Widget *wid, int x, int y, int w, int h, void *data) {
	_go_callbackHandler((uintptr_t)data);
}

void go_deleter(Fl_Widget *w, void *data) {
	_go_callbackHandler((uintptr_t)data);
}
