#!/usr/bin/make -f

.ONESHELL:
.SHELL := /usr/bin/bash

AUTHOR := "noelruault"
PROJECTNAME := $(shell basename "$$(pwd)")
PROJECTPATH := $(shell pwd)
GOFLAGS :=

help:
	echo "Usage: make [options] [arguments]\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

go-build: ## Compiles packages and dependencies. Builds a binary under bin/. *Accepts GOFLAGS and LDFLAGS
	@[ -d bin ] || mkdir bin
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(GOFLAGS) -o bin/ "$(PROJECTPATH)/cmd/..."

start: ## Starts go-fibonacci service. *Accepts GOFLAGS and LDFLAGS
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOFLAGS) "$(PROJECTPATH)/cmd/main.go" $(LDFLAGS)

go-doc: ## Generates static docs
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) godoc -http=localhost:6060

go-vendor: ## Updates vendor dependencies
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go mod vendor && go mod tidy

test: ## Runs the tests
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) GOFLAGS="-count=1" go test ./...

bench: ## Runs the benchmarks
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -benchmem -bench=. ./internal/models -count=1

docker-build: ## Builds the project binary inside a docker image
	@docker build -t "$(AUTHOR)/$(PROJECTNAME)" .

k8-minikube-start: docker-build ## Builds the service image, deploys and starts the go-fibonacci service inside a kubernetes cluster, checks if minikube is running and points your shell to minikube's docker-daemon
	@if [ `minikube status --format "{{.Host}}"` != "Running" ]; then \
        printf "\033[0;33mminikube needs to be available and running\033[0m %s\n"; \
        exit 1; \
    fi
	@eval $(minikube -p minikube docker-env)
	@kubectl apply -f fib.yaml


k8-start: ## Deploys and starts the go-fibonacci service inside a kubernetes cluster
	@kubectl apply -f fib.yaml


k8-stop: ##  Deletes the minikube go-fibonacci deployment and service
	@kubectl delete -f fib.yaml
