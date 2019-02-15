# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run

# App info
BINARY_NAME := baton
VERSION := 0.1.6

# Folders
OUTPUT_FOLDER := build
BINARY_OUTPUT := $(OUTPUT_FOLDER)/$(BINARY_NAME)

# Output platforms
PLATFORMS := windows linux darwin
os = $(word 1, $@)

# Make for all the platforms
.PHONY: $(PLATFORMS)
ifeq ($(OS),Windows_NT)
$(PLATFORMS): # If host is windows
	set GOOS=$(os)&& set GOARCH=amd64&& $(GOBUILD) -o $(BINARY_OUTPUT)-$(VERSION)-$(os)-amd64$(if $(filter $(os),windows),.exe,) main.go
else
$(PLATFORMS): # else it's Linux/MacOS
	GOOS=$(os) GOARCH=amd64 $(GOBUILD) -o $(BINARY_OUTPUT)-$(VERSION)-$(os)-amd64$(if $(filter $(os),windows),.exe,) main.go
endif

.PHONY: build
build: windows linux darwin

.PHONY: run
run:
	$(GORUN) main.go

.DEFAULT_GOAL := build