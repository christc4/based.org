# Vim tutorial

<img src="/.pix/vim_cigarretes.avif" style="width:550px; height: auto;">

This is my .vimrc

	let mapleader=" " " 
	set clipboard=unnamedplus
	set ttyfast 

	autocmd BufWinLeave *.* mkview " Save Folds
	autocmd BufWinEnter *.* silent loadview " View Folds

	map <leader>c :w! \| !compiler "%:p"<CR>
	map <leader>p :!opout "%:p"<CR>

	map <leader>p :!opout "%:p"<CR>
	map <leader>d :w<CR>
	map <leader>h :wq<CR>
	map <leader>t :q!<CR>

## What you MUST do

### No. 1

Stop using default keybinds to save a file, don't bother with 

To default save a file one has to...

- hold shift
- press ':'
- enter 'w'

if you are not already in escape mode one has to 

- press 'esc'

that is too many keypresses,

even 'ZZ' is no good  

# Useful

- `gj` and `gk`
	- move cursor up and down to wrapped part of a line
- `g0` and `g$`
	- move cursos te first and last letter of wrapped line
- `gq`
	- turn long line into multiple lines
- `gu` and `gU` 
	- uncapitalize and capitalize words/lines
- `~` and `g~`
	- switch capitalization of a letter
- `gf` 
	- open highlighted text as file
- `gv`
	- reselect previous selected text
- `gJ`
	- conjoining lines without leaving spaces
- `g&` 
	- re-run substitute command for all lines

- `<`
	- remove indent, first highlight

- `>` 
	- apply indent, firsh highlight
