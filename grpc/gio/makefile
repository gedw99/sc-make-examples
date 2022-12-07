# https://git.sr.ht/~whereswaldon/rosebud

REPO_NAME=rosebud

print:

dep:
	git clone https://git.sr.ht/~whereswaldon/rosebud
dep-delete:
	rm -rf $(REPO_NAME)


mod-upgrade:
	go install github.com/oligot/go-mod-upgrade@latest
	cd $(REPO_NAME) && go-mod-upgrade
	$(MAKE) mod-tidy 
mod-tidy:
	cd $(REPO_NAME) && go mod tidy



example-run:
	# just the tree widget
	cd $(REPO_NAME)/example && go run .

WASM_TEST_BIN_NAME=wasm-headless
WASM_TEST_BIN=$(GOPATH)/bin/$(WASM_TEST_BIN_NAME)
wasm-test-print:
	@echo
	@echo 'WASM_TEST_BIN_NAME:           $(WASM_TEST_BIN_NAME)'
	@echo 'WASM_TEST_BIN:                $(WASM_TEST_BIN)'
	@echo 'WASM_TEST_BIN installed ?:    $(shell which $(WASM_TEST_BIN_NAME))'
wasm-test-dep:
	go install github.com/agnivade/wasmbrowsertest@latest
	mv $(GOPATH)/bin/wasmbrowsertest $(WASM_TEST_BIN)
wasm-test-dep-delete:
	rm -f $(WASM_TEST_BIN)
wasm-test-run:
	cd $(REPO_NAME) && GOOS=js GOARCH=wasm go test -exec="$(WASM_TEST_BIN)" "$@"
	cd $(REPO_NAME) && WASM_HEADLESS=off GOOS=js GOARCH=wasm go test -exec="$(WASM_TEST_BIN)" "$@"



go-test:
	cd $(REPO_NAME) && go test ./... -v -cover

server-run:
	# builds the client into the server
	cd $(REPO_NAME)/demo/server && chmod +x ./build.bash 
	cd $(REPO_NAME)/demo/server && ./build.bash 
	cd $(REPO_NAME)/demo/server && ./server
