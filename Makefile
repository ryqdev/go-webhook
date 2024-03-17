build:
	tmux send-keys -t unreal C-c
	tmux send-keys -t unreal 'git pull' C-m
	tmux send-keys -t unreal 'go run main.go' C-m
