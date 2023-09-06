### wasi-examples

Este repositorio contiene algunos ejemplos para demostrar el soporte a [WebAssembly System Interface](https://wasi.dev/). Para ejecutar el WASM resultante de la compilación, necesitas un runtime de web assembly compatible con WASI.
En los ejemplos se usa WASMEdge, pero podrían ejecutarse en otros entornos como [Wasmtime](https://wasmtime.dev/) cambiando las banderas usadas.

Requerimientos globales:
* [Go 1.21 o superior](https://go.dev/dl/)
* [Wasmedge](https://wasmedge.org/docs/):
    > curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash
