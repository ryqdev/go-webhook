.PHONY: run
run:
	tmux send-keys -t mdbook C-c
	tmux send-keys -t mdbook 'git pull' C-m
	tmux send-keys -t mdbook 'make' C-m
