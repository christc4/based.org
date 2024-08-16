# Tmux

## My personal config

	#Reload conf by Ctrl + B, then write :source-file ~/.tmux.conf
	set-option -g default-shell "/bin/zsh"

	# Status
	set -g message-style bg=default,fg='#B6E354'
	set -g status-style bg=default,fg='#B6E354'
	set -g status-interval 1


	set-option -g status-interval 5
	set-option -g automatic-rename on
	set-option -g automatic-rename-format '#{b:pane_current_path}'

	set -g status-left ''
	set -g status-right ''

	set-window-option -g window-status-style fg='#a9a9a9',bg=default
	set-window-option -g window-status-current-style fg='#000000',bg=default

	set -g status-justify centre

	# bind -n DoubleClick1Pane run-shell "dragon -x '#{pane_current_path}/#{mouse_word}'"

	# Status end

	set -g prefix C-b

	set -g base-index 1
	set-option -g mouse on
	set -g set-clipboard on
	set -g history-limit 10000

	set -g default-terminal "xterm-256color"
	set -g terminal-overrides "xterm-256color"
	set -ga terminal-overrides ",xterm-256color:Tc"


	# Keybinds

	# Movement and resize vim controls
	bind -r C-k resize-pane -U
	bind -r C-j resize-pane -D
	bind -r C-h resize-pane -L
	bind -r C-l resize-pane -R

	bind -r k select-pane -U
	bind -r j select-pane -D
	bind -r h select-pane -L
	bind -r l select-pane -R

	# Rebind splits
	bind v split-window -h -c '#{pane_current_path}'
	bind b split-window -v -c '#{pane_current_path}'

	unbind Up
	unbind Down
	unbind Left
	unbind Right

	unbind C-Up
	unbind C-Down
	unbind C-Left
	unbind C-Right

	# Rebind rename to R
	bind-key r command-prompt -I "#W" "rename-window '%%'"

