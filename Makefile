Name := key-generator
Version := $(shell git describe --tags --abbrev=0)
OWNER := youyo
.DEFAULT_GOAL := help

## Setup
setup:
	go get -u github.com/golang/dep/cmd/dep
	go get github.com/Songmu/make2help/cmd/make2help

## Install dependencies
deps: setup
	dep ensure -update

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: setup deps update vet lint test help
