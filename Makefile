.PHONY: build
build:
	go build -o ./build/radiotray

.PHONY: win
win:
	go build -o ./build/radiotray.win.exe -ldflags "-H=windowsgui"
