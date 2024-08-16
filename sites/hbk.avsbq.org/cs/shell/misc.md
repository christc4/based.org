for file in *; do mv "$file" "$(echo "$file" | tr -cd '0-9.png' | sed 's/[^0-9]*//g').png"; done

