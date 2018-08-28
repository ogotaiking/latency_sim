# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
LDFLAGS = -s -w
BINARY_NAME=client
BINARY_PATH=.


all: build
build:  linux mac win rpi

linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags="$(LDFLAGS)" -o $(BINARY_PATH)/$(BINARY_NAME)_linux -v
mac:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -ldflags="$(LDFLAGS)" -o $(BINARY_PATH)/$(BINARY_NAME)_mac -v
win:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags="$(LDFLAGS)" -o $(BINARY_PATH)/$(BINARY_NAME)_win.exe -v
rpi:
	GOOS=linux GOARCH=arm $(GOBUILD) -ldflags="$(LDFLAGS)" -o $(BINARY_PATH)/$(BINARY_NAME)_rpi -v
