build:
	GOOS=wasip1 GOARCH=wasm go build -v main.go 
run: build
	wasmtime run --mapdir $(PWD)::. main
test: build 
	wasmtime run main