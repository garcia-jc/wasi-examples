### wasi-examples

Este repositorio contiene algunos ejemplos para demostrar el soporte a [WebAssembly System Interface](https://wasi.dev/). Para ejecutar el WASM resultante de la compilación, necesitas un runtime de web assembly compatible con WASI.
En los ejemplos se usa WASMEdge, pero podrían ejecutarse en otros entornos como [Wasmtime](https://wasmtime.dev/) cambiando las banderas usadas.

[![asciicast](https://asciinema.org/a/5xKtQ78BAGVAiVP9TVLdXcNTQ.svg)](https://asciinema.org/a/5xKtQ78BAGVAiVP9TVLdXcNTQ) 

Requerimientos globales:
* [Go 1.21 o superior](https://go.dev/dl/)
* make
* [Wasmedge](https://wasmedge.org/docs/):
    > curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash

Cada folder es un proyecto ejecutable. Para ejecutarlos, abre una terminal en alguno de los proyectos, y usa `make run`
```bash
cd hello
make run 
==========
GOOS=wasip1 GOARCH=wasm go build -o hello.wasm -v main.go 
wasmedge run --dir /:/home/ubuntu/exp/wasi-examples/hello:readwrite -- hello.wasm
2023/09/06 19:58:38 ---------- Makefile
2023/09/06 19:58:38 ---------- build.out
2023/09/06 19:58:38 ---------- hello.wasm
2023/09/06 19:58:38 ---------- main.go
duracion 3.178519ms
```

* hello: lista los archivos a los que tiene acceso el runtime de WASI
* chans: Demuestra el uso de canales y `random`
* http: Inicializa un servidor HTTP
* dasel: Compila el proyecto [dasel](https://github.com/TomWright/dasel), (una alternativa a jq), y lo usa para hacer algunas modificaciones
