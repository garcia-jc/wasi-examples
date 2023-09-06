build:
	GOOS=wasip1 GOARCH=wasm go build -o chans.wasm -v main.go 
run: build
	wasmedge run chans.wasm