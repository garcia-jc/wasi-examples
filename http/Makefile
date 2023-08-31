BINARY_NAME=http-server.wasm
RUNTIME_NAME=io.containerd.wasmedge.v1
PLATFORM_NAME=wasi/wasm
IMAGE_TAG=hello-wasi:latest
# ARCH?=amd64

build:
	GOOS=wasip1 GOARCH=wasm go build -o $(BINARY_NAME) -v main.go 
run: build
	wasmedge $(BINARY_NAME)

package: build 
	docker build --platform $(PLATFORM_NAME) -t $(IMAGE_TAG) .
docker-run: package
	docker run --runtime=$(RUNTIME_NAME) --platform=$(PLATFORM_NAME) --rm -p 3000:3000 $(IMAGE_TAG)