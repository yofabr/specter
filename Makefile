OS := $(shell uname -s)

ifeq ($(OS), Linux)
	EXECUTABLE = specter
else ifeq ($(OS), Darwin)
	EXECUTABLE = specter
else
	EXECUTABLE = specter.exe
endif

build:
	go build -o $(EXECUTABLE) ./cmd/specter

install:
	go build -o $(EXECUTABLE) ./cmd/specter
	sudo mv $(EXECUTABLE) /usr/local/bin

test:
	go test ./...

fmt:
	go fmt ./...

clean:
	rm -f $(EXECUTABLE)

.PHONY: build install clean