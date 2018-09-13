.PHONY: all aerod clean install uninstall run test test-all
DEST = $(shell pwd)/build/bin

all: aerod

aerod:
	./env.sh go install ./cmd/aerod
	@echo "$(DEST)/aerod"

clean:
	@rm -rf build/

install: aerod
	@cp -f $(DEST)/aerod $GOPATH/bin/

uninstall:
	@rm -f $GOPATH/bin/aerod

run: aerod
	@$(DEST)/aerod

test: test-all

test-all:
	@go test -v ./...
