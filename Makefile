SHELL = /bin/bash

.PHONY: deps
deps:
	dep ensure -v

.PHONY: setup
setup:
	go get github.com/golang/dep/cmd/dep

.PHONY: lint
lint: deps
	gometalinter

.PHONY: test
test: deps
	go test ./...
	go test -v -tags=integration -run=^./integration

.PHONY: integration-test
test: deps
	go test -tags=integration ./...

.PHONY: coverage
coverage: deps
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: codecov
codecov: deps coverage
	bash <(curl -s https://codecov.io/bash)

.PHONY: circleci
circleci:
	circleci build -e GITHUB_TOKEN=$GITHUB_TOKEN