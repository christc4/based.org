# Acme not Acne

<details><summary>Resources used:</summary>
[evbogdanox](https://github.com/evbogdanov/acme/blob/master/README.md)
</details>

<img src=".pix/acme_chord.avif" style="width:410px; height: auto;"> Image by Tom Lieber

`Edit =`

Find current line number

`:13`

Go to the 13th line

`:0`

Go to file beginning

`:$` 

Go to file end

:1,$ or :, or Edit 1,$ or Edit , select all lines

:1,5 or Edit 1,5 select lines 1..5
Edit , d clear window
Edit , < echo hello world replace window body with some text
Edit , < erl -man maps replace window body with erlang manual
Edit , s/text/TEXT/g or Edit , | sed 's/text/TEXT/g' global replace
$% or $samfile current file name
$winid current window id
echo some text | 9p write acme/$winid/body append to the end of current window
keyboard shortcuts:
ctrl-u delete from cursor to start of line
ctrl-w delete word before the cursor
ctrl-h delete character before the cursor
ctrl-a move cursor to start of the line
ctrl-e move cursor to end of the line
ctrl-i tab
ctrl-j enter
ctrl-f filepath autocompletion
fn-*left arrow* go home (macOS)
fn-*right arrow* go end (macOS)
search with right click:
:+/foobar, :/foobar and just foobar search forward
:-/foobar search backwards
press esc to select the last typed text
press esc again to delete any selected text
Font switch between fonts
:/^hel regexp match: lines starting with 'hel'
:/lo\n/ regexp match: lines ending with 'lo'
:/^b/,/^e/ regexp match: lines between starting with 'b' and starting with 'e'
Dump write the state of acme to the file
Load restore from the dump
Edit , > python pipe window body through python interpreter
three-finger tap emulates middle click (macOS)
Sam commands
Edit +/hello search 'hello' forward
Edit -/hello search 'hello' backward
Edit , > wc -l count lines in file
Edit , | sort sort lines
Edit 3,5p print lines 3..5 in new window
Edit 3,5 | upper lines 3..5 upper cased
Edit 3,5 s/HE/he/g replace on 3..5 lines only
Edit 2 d delete 2nd line
Edit 2 c/new change 2nd line
Edit 2 a/new append text after 2nd line
Edit 2 i/new insert text before 2nd line
Edit 2 < date replace 2nd line with the output of date
Edit ,x/^TODAY$/ < date replace matching lines with the output of date
Edit ,x/Plan9/ |tr a-z A-Z replace all instances of Plan9 with upper case
Edit 3,5x/^/ a/	/ indent lines 3..5 with 1 tab
You can do amazing things with Sam commands:

Edit ,x/Acme/ {
  i/I love 
  c/Sam
  a/ editor!
}
Cut / Copy selection to a file
select some text
cut: | sed '' > file.txt
copy: > sed '' > file.txt
pipe selection to a file: > awk '{ print(toupper($1)) }' | sort | nl > file.txt
Games!
Tic-tac-toe
Mini sea battle
Other Plan9 goodies
win start shell in a new window
page FILE view graphics files
web URL open url in your browser
