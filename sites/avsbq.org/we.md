	#!/bin/sh

	# Check if at least one argument (the markdown file) is provided
	if [ "$#" -lt 1 ]; then
	  echo "Usage: $0 markdown_file.md"
	  exit 1
	fi

	# Assign the first argument to the markdown file variable
	markdown_file="$1"

	# Check if the markdown file exists; if not, exit with an error
	if [ ! -f "$markdown_file" ]; then
	  echo "Error: $markdown_file does not exist."
	  exit 1
	fi

	# Loop through all .avif files in the current directory
	for img in *.avif; do
	  # Check if there are any .avif files
	  if [ -e "$img" ]; then
	    # Append the markdown image link to the markdown file
	    echo "<img src=$img>" >> "$markdown_file"
	  fi
	done
