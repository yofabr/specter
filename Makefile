OS := $(shell uname -s)

ifeq ($(OS), Linux)
	EXECUTABLE = specter
else ifeq ($(OS), Darwin)
	EXECUTABLE = specter
else
	EXECUTABLE = specter.exe
endif

build:
	go build -o $(EXECUTABLE) main.go

install:
	go build -o $(EXECUTABLE) main.go
	sudo mv $(EXECUTABLE) /usr/local/bin

clean:
	rm -f $(EXECUTABLE)

.PHONY: build install clean