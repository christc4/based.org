NAME:
xid - an Xtensible input device

DESCRIPTION:
xid is a terminal multi-pane line-oriented text editor combinging the strengths of vi(m) and acme.

EXAMPLE:
By design, xid is designed to print very little - that is to say text is displayed sparingly. This hails from the days of teletype, wherein the venerable ed thrived. But more relavantly, this is to prevent users from being bombarbed with text, this way users are forced to think what they want to edit.

Running xid with no arguments, spawns a user pane,

		| $					|

This is xid's command interface, your primary mode of editing text.
Here xid accepts two arguments n and q, name and quit.

		| $ n foobar.txt			|	# Providing file name

Once a file name is provided, xid creates and saves the empty file.

Typing a, puts you in append mode like ed.

		| $ a		 			|	# Appending file
		| $ Buy groceries		        |
		| $ .				        |	# Saving file

The delineator "." saves the file, launching a display pane.

		| 1# 1 Buy groceries...	        |

It may not look like much has happened, but ed has created a fold. By default lines or paragraphs, seperated by empty lines create folds. #1, means this is fold 1.  1, shows you are on line 1.

Were you to provide more lines through the user pane.

		| $ a		 			|	# Appending file
		| $ Wash car			     	|
		| $ Clean ears			        |
		| $ .				        |	# Saving file

Your fold will now display

		| 1# 1 [3] - Buy groceries...	        |

[3] shows how many lines are on a fold, folds appear displaying the first line, although you may wish to rename this


		| $ rn f1 "To-do list|			|	# Renames f

Your display pane should now look like this

		| 1# 1 [3] - *To-do list*	        |

Named folds, will appear italic. Jumping back to pane 1, you can perform regex expressions

You can unfold - folds - with the uf command e.g. "uf f1"

Like acme, the display pane has a bar. This bar will only appear/exist, if in the config.h, you specify your scripts path, here are some example scripts.

Dict | Thes | eBay |

Left clicking text in the display pane, then middle clicking an option in the display bar will run the script on your selected text. Say for example you left click "car" and middle click eBay, xId, will launch eBay listings for car, this is up to you how you run it from hereon. However, if you run scripts which don't require graphical browser or use terminal tools, in your user pane, the output will be printed.

NOTE, xid, provides no scripts.

This displays all text on the display pane, in the display pane you can use plumbing. Although you cannot directly edit or type in the display pane, you can perform plumbing. Mouse chording, left clicking highlights all text








xid notes.txt, if file is empty xid will only display one pane, as soon as the user saves new text onto the file, a new pane is summoned. By default,


FEATURES:
	- TWO PANES
	- AUTO FOLDING TEXT
	- PLUMBING

PANE 1: (text display & plumbing)

	- All text is folded, lines with
	- all text is folded, with comments one line above, e.g. #P1 Once upon a time...
	- piping, with mouse, middle click brings menu, scriptable

PANE 2: (command line)
	- input ed commands, regex expressions,
	- vi mode, accept line numbers e.g. 1v
	- brings line 1 into buffer, and like zsh vi-mode, vi mode is enabled
	- ,n "compiles" text, auto-refreshing PANE 1

