all: go-base64.so

go-base64.so: main.go init.c
	go build -buildmode=c-shared -o go-base64.so

clean:
	rm go-base64.so go-base64.h
