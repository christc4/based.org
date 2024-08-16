# Basic dmenu bookmarking

	declare -A bookmarks=(
	  ["yt"]="https://youtube.com"
	  ["pm"]="https://mail.protonmail.com/u/0/all-mail"
	  ["cc"]="https://christc4.github.io/"
	)

	declare -a options=(
	  "yt"
	  "pm"
	  "cc"
	)

	choice=$(printf '%s\n' "${options[@]}" | dmenu -l 20 -p 'Bookmarks')

	if [ "$choice" ]; then
	  name=$(printf '%s\n' "${choice}" | awk '{print $NF}')
	  url="${bookmarks[$name]}"
	  $BROWSER "$url"
	else
	  exit 1
	fi

## Don't do this

No. 1, you need to edit the script over and over, if you want to add/remove bookmarks.

## Even better

Use external text file

Incorporate fuzzy finding, I recommend fzf

Allow for non-urls to be stored, needs to know how to handle non web-links and what to do with non-links
