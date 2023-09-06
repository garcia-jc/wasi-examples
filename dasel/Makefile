ENV_FILE=/env.json
OUT_FILE=/env.toml

build:
	GOOS=wasip1 GOARCH=wasm go install -v github.com/tomwright/dasel/v2/cmd/dasel@v2.3.6
	mv $(GOPATH)/bin/wasip1_wasm/dasel dasel.wasm

write-env: 
	go env -json > env.json

run: build write-env
	wasmedge run --dir /:$(PWD):readwrite -- dasel.wasm \
		-f $(ENV_FILE) '.GOOS' --colour
	wasmedge run --dir /:$(PWD):readwrite -- dasel.wasm \
		-f $(ENV_FILE) put -t string -v 'wasi' 'GOOS'
	wasmedge run --dir /:$(PWD):readwrite -- dasel.wasm \
		-f $(ENV_FILE) '.GOOS' --colour