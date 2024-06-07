unreal:
	tmux send-keys -t unreal C-c
	tmux send-keys -t unreal 'git pull' C-m
	tmux send-keys -t unreal 'go run main.go' C-m

mdbook:
	tmux send-keys -t mdbook C-c
	tmux send-keys -t mdbook 'git pull' C-m
	tmux send-keys -t mdbook 'mdbook serve -p 8888 -n 0.0.0.0' C-m
