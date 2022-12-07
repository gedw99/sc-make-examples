# https://github.com/tarndt/wasmws

print:

dep:
	git clone https://github.com/tarndt/wasmws
dep-delete:
	rm -rf wasmws



mod-upgrade:
	go install github.com/oligot/go-mod-upgrade@latest
	cd wasmws && go-mod-upgrade
	$(MAKE) mod-tidy 
mod-tidy:
	cd wasmws && go mod tidy



wasm-test:
	cd wasmws && chmod +x ./test.bash 
	cd wasmws && ./test.bash -v -cover

go-test:
	cd wasmws && go test ./... -v -cover

server-run:
	# builds the client into the server
	cd wasmws/demo/server && chmod +x ./build.bash 
	cd wasmws/demo/server && ./build.bash 
	cd wasmws/demo/server && ./server
