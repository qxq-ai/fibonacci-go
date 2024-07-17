WASM_FILE ?= main.wasm
GO_FILE ?= api.go
DIRECTORY ?= .

build:
	GOOS=js GOARCH=wasm go build -o ./$(DIRECTORY)/$(WASM_FILE) ./$(DIRECTORY)/$(GO_FILE)
	cp $$(go env GOROOT)/misc/wasm/wasm_exec.js ./$(DIRECTORY)
	
server:
	go run server.go  ./$(DIRECTORY)