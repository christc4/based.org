#  Fuzzily find files

	clear &&
	vim "$(find /var/www/werc/sites -type d -name '.git' -prune -o -type f \
	-not -name '*.webp' \
	-not -name '*.webm' \
	-not -name '*.pdf' \
	-not -name '*.jpg' \
	-not -name '*.jpeg' \
	-not -name '*.gif' \
	-not -name '*.png' \
	-not -name '*.mp4' -print | fzf --info=inline --preview 'cat {}')"

## Alternatives

You can change 'vim' to '$EDITOR'

If you want more color in your life, try 'bat' instead of 'cat'

	bat --style=numbers --color=always --line-range :500 {}

## explanation

this script excludes everything that's in a '.git' directory and excludes images or uneditable documents
