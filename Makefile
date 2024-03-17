build:
	tmux new -s unreal
	tmux send-keys -t unreal C-c
	git pull
	tmux send-keys -t unreal 'go run main.go' C-m
