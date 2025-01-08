test:
	tmux new-session -s golang-testcontainers

git:
	git add .
	git commit -m "Version"
	git push -u origin
