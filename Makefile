.DEFAULT_GOAL := help

## Setup
setup:
	go get -v -u github.com/golang/dep/cmd/dep

## Install dependencies
deps:
	dep ensure

## Start Server
run:
	go run *.go

## Deploy to heroku
deploy:
	heroku container:push web

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: setup deps run deploy help
