# Fuzzily cd

	cd $(find /var/www/werc/sites -type d \( -name .git -prune \) -o \( -type d -print \) | fzf) && clear
