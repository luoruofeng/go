all: deps build

.PHONY: deps
deps:
	go get -d -v github.com/dustin/go-broadcast/...
	go get -d -v github.com/gin-gonic/gin

.PHONY: build
build: deps
	go build -o realtime-chat main.go room.go

.PHONY: clean
clean:
	go clean

