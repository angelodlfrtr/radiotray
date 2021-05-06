.PHONY: build
build:
	go build -o ./build/radiotray

.PHONY: clean
clean:
	rm -Rf ./build

.PHONY: win
win:
	go build -o ./build/radiotray.win.exe -ldflags "-H=windowsgui"

.PHONY: lint
lint:
	find . -name "*.go" | xargs misspell -error
	golint -set_exit_status ./...
	go vet ./...
	staticcheck ./...

.PHONY: darwin-app
darwin-app:
	make clean
	make build
	mkdir -p ./build
	cp -r ./lib/RadioTrayBase.app ./build/RadioTray.app
	cp ./build/radiotray ./build/RadioTray.app/Contents/MacOS
