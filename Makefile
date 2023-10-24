BINARY_NAME := bar-mail-body

build:
		-mkdir _build
		go build -ldflags="-extldflags=-static" -o _build/$(BINARY_NAME) ./cmd/main.go

install: build
		sudo cp _build/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
		sudo cp bar-mail.sh /usr/local/bin/
