# tested with wasmedge version 0.13.3
install-wasmedge:
	curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash
build:
	GOOS=wasip1 GOARCH=wasm go build -o build.out -v main.go 
run: build
	wasmedge build.out
