# XorgRaiseWindow

[The original author](https://www.linuxquestions.org/questions/linux-general-1/how-to-bring-up-application-window-to-front-from-shell-script-83545/)

In some xorg window managers, users have reported wmctrl not working out of the box. Since we are already using the x11/xlib library for our respective window managers, we can write a c++ program (see above for
original author), to raise windows. We then call this with a simple shell script and map this to a keypress.

	#include <stdio.h>
	#include <stdlib.h>
	#include <X11/Xlib.h>

	int main(int argc, char **argv)
	{
	if ( argc != 2 ) {
	printf("Usage:\n\ttotop <window id>\n");
	return 1;
	}

	Display *dsp = XOpenDisplay(NULL);

	long id = strtol(argv[1], NULL, 16);
	XRaiseWindow ( dsp, id );
	XSetInputFocus ( dsp, id, RevertToNone, CurrentTime );

	XCloseDisplay ( dsp );

	return 0;
	}


When done run,

	c++ totop.cpp -L/usr/X11R6/lib -lX11 -o totop

We have named this program 'totop' as in to the top of the window stack

The shell script,

	id=$(xwininfo -name "st" | grep id: | awk "{ print \$4 }")
	totop $id

Alternatively you can ommit the second line, and only use the third with the id if you know it
