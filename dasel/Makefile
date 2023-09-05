ENV_FILE=env.json
OUT_FILE=env.toml

build:
	GOOS=wasip1 GOARCH=wasm go install github.com/tomwright/dasel/v2/cmd/dasel@v2.3.6
	mv $(HOME)/go/bin/wasip1_wasm/dasel dasel.wasm

write-env: 
	go env -json > $(ENV_FILE)

run: build write-env
	wasmtime --mapdir $(PWD)::. dasel.wasm \
		-- -f $(ENV_FILE) '.GOOS' --colour
	wasmtime --mapdir $(PWD)::. dasel.wasm \
		-- -f $(ENV_FILE) put -t string -v 'wasi' 'GOOS'
	wasmtime --mapdir $(PWD)::. dasel.wasm \
		-- -f $(ENV_FILE) '.GOOS' --colour