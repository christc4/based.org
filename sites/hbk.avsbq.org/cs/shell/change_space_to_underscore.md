# Change '\ '  to '&#95;'

> Tested on OpenBSD 7.5, ksh

`find . -name '* *' | while IFS= read -r f; do mv "$f" "$(echo "$f" | tr ' ' '_')"; done`

