	#!/bin/bash

	input_folder=$(pwd)
	output_folder=$(pwd)
	background_color="#0c1214"

	# Ensure the output folder exists
	mkdir -p "$output_folder"

	# Iterate through all GIF files in the input folder
	for input_file in "$input_folder"/*.gif; do
	    # Extract the filename without the extension
	    filename=$(basename -- "$input_file")
	    filename_no_extension="${filename%.*}"

	    # Set the output file path
	    output_file="$output_folder/${filename_no_extension}.gif"

	    # Apply background removal
	    convert "$input_file" -transparent "$background_color" "$output_file"
	done
