build:
	GOOS=wasip1 GOARCH=wasm go build -o build.out -v main.go 
run: build
	wasmedge run --dir /:$(PWD):readwrite -- build.out
test: build 
	wasmedge run build.out