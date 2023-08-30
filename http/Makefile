build:
	GOOS=wasip1 GOARCH=wasm go build -o build.out -v main.go 
run: build
	wasmedge --dir .:. build.out
