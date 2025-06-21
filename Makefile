.PHONY: build install clean

build:
	go build -o bin/bm main.go

install:
	sudo mv bin/bm /usr/local/go/bin/bm

clean:
	sudo rm -rf /usr/local/go/bin/bm