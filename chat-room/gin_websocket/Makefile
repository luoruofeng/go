all: deps build

.PHONY: deps
deps:
	go get -d -v github.com/dustin/go-broadcast/...
	go get -d -v github.com/gorilla/websocket

.PHONY: build
build: deps
	go build -o realtime-chat main.go room.go tmp.go
