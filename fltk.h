#pragma once

typedef struct Fl_Widget Fl_Widget;

void my_timeout_handler(void* data);

void my_awake_handler(void* data);

void callback_handler(Fl_Widget *w, void* data);

int event_handler(Fl_Widget *w, int ev, void* data);

void resize_handler(Fl_Widget *wid, int x, int y, int w, int h, void *data);

void go_deleter(Fl_Widget *w, void *data);