# Uppercase to lower

`for file in *; do mv "$file" "$(echo "$file" | tr '[:upper:]' '[:lower:]')"; done`

