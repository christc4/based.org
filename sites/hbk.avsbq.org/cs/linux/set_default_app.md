# Set default apps

	export BROWSER="/usr/bin/browser"

Some programs that are graphical, are under no obligation to use the $BROWSER variable

Write your own desktop entry

	$ xdg-mime default browser.desktop x-scheme-handler/http
	$ xdg-mime default browser.desktop x-scheme-handler/https 
