GO ?= go
CASK ?= cask

all: go-base64.so

go-base64.so: main.go init.c
	$(GO) build -buildmode=c-shared -o $@

clean:
	$(RM) go-base64.so go-base64.h

test: go-base64.so
	$(CASK) exec ert-runner
