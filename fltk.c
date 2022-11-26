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