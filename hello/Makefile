build:
	GOOS=wasip1 GOARCH=wasm go build -o hello.wasm -v main.go 
run: build
	wasmedge run --dir /:$(PWD):readwrite -- hello.wasm
test: build 
	wasmedge run hello.wasm