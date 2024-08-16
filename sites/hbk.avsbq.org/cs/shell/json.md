# Json

	#!/bin/sh

	# Initialize an empty JSON array
	echo "[" > s.json

	# Find all .md files in the current directory and subdirectories
	find . -type f -name "*.md" | sed 's|^\./||; s|\.md$||' > temp_files.txt

	# Read the list of files and format them into JSON
	awk '{
	  if (NR == 1) {
	    printf "  \"%s\"", $0
	  } else {
	    printf ",\n  \"%s\"", $0
	  }
	} END {
	  print "\n]"
	}' temp_files.txt >> s.json

	# Clean up temporary file
	rm temp_files.txt

	# Inform the user
	echo "s.json created successfully."

