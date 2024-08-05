# Define the variables with default values
WASM_FILE ?= main.wasm
GO_FILE ?= api.go
DIRECTORY ?= .

# Build rule
build:
	@if [ ! -f "$(DIRECTORY)/go.mod" ]; then \
		echo "Initializing Go module in $(DIRECTORY)"; \
		cd $(DIRECTORY) && go mod init $(basename $(DIRECTORY)) && go mod tidy; \
	fi
	GOOS=js GOARCH=wasm go build -o $(DIRECTORY)/$(WASM_FILE) $(DIRECTORY)/$(GO_FILE)
	cp $$(go env GOROOT)/misc/wasm/wasm_exec.js $(DIRECTORY)

# Server rule
server:
	go run server.go $(DIRECTORY)
