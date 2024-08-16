# Batch image conversion

*Last update: 2024/07/24*

	#!/bin/sh

	# Function to convert .webp to .avif using ImageMagick
	convert_webp_to_avif() {
	    input_file="$1"
	    output_file="${input_file%.webp}.avif"
	    convert "$input_file" "$output_file"
	}

	# Find all .webp files and convert them
	find . -type f -name '*.webp' -exec sh -c '
	    for file; do
		output_file="${file%.webp}.avif"
		convert "$file" "$output_file"
	    done
	' sh {} +

	echo "Conversion complete!"

