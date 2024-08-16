# [Variables](variables) in Bash

Up: [Bash Shell Scripts](bash_shell_scripts)
Brother(s):
TARGET DECK

This is just like standard variable assignment in programming, for example: 

`count=0` would assign the value 0 to the `count` variable.

The output of a [UNIX](unix) command (like `cat`) can be used as the value for a variable!

Usage: `data=$(command)`

Example: `data=$(cat helloWorld.txt); echo helloWorld.txt`

Output: `SIKE YOU THOUGHT THIS HAD HELLO WORLD IN IT SHURRUP`

(The dollar sign ($) is crucial in specifying that the variable is being referenced, and not assigned.)


































#### Why:
#### How:









