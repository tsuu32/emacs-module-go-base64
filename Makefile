GO ?= go
EMACS ?= emacs

all: go-base64.so

go-base64.so: main.go init.c
	go build -buildmode=c-shared -o $@

clean:
	rm -f go-base64.so go-base64.h

test: go-base64.so
	$(EMACS) --batch -L . -l go-base64 --eval \
	'(progn \
	   (message (b64-encode "hello")) \
	   (message (b64-decode "aGVsbG8=")))'
