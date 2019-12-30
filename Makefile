.PHONY: build

GOBIN = ./build/bin
GO ?= latest

build:
	go build -o build/gproton ./cmd/gproton
	@echo "Done building."