	if [ -z "$BROWSER" ]; then
	    echo "The BROWSER environment variable is not set. Please set it to your web browser command."
	    exit 1
	fi

	if [ $# -ne 1 ]; then
	    echo "Usage: $0 <chinese_character>"
	    exit 1
	fi

	chinese_character="$1"

	url="https://tatoeba.org/en/sentences/search?from=cmn&query=$chinese_character&to=eng"
