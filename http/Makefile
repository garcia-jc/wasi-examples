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
	docker buildx build --platform $(PLATFORM_NAME) -t $(IMAGE_TAG) .
docker-run: package
	docker run --runtime=$(RUNTIME_NAME) --platform=$(PLATFORM_NAME) \ 
		--rm -p 3000:3000 $(IMAGE_TAG)

# # tested with wasmedge version 0.13.3
# install-wasmedge:
# 	curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash

# SHIM_PACKAGE=containerd-shim-wasmedge-linux-$(ARCH).tar.gz
# build-wasi-shim:
# 	wget https://github.com/containerd/runwasi/releases/download/containerd-shim-wasmedge%2Fv0.1.1/$(SHIM_PACKAGE)
# 	tar -xvf $(SHIM_PACKAGE)
# 	rm $(SHIM_PACKAGE) containerd-wasmedged containerd-shim-wasmedged-v1
# 	sudo mv containerd-shim-wasmedge-v1 /usr/local/bin
# 	echo "{\"runtimes\":{\"wasmedge\": {\"path\": \"/usr/local/bin/containerd-shim-wasmedge-v1\"}}}" | sudo tee /etc/docker/daemon.json
