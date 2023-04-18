# Image URL to use all building/pushing image targets
IMG ?= xxx:yyy

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: build


##@ Build

.PHONY: build
build: ## Build manager binary.
	go build -o bin/manager main.go

.PHONY: run
run: ## Run a controller from your host.
	go run ./main.go

.PHONY: docker-build
docker-build: ## Build docker image with the server.
	docker build -t ${IMG} .

.PHONY: docker-push
docker-push: ## Push docker image with the server.
	docker push ${IMG}

.PHONY: docker-run
docker-run: ## Run docker image with the server.
	docker run -d -p 8080:8080 ${IMG}

.PHONY: docker-run-limit
docker-run-limit: ## Run docker image with the server.
	docker run -d -p 8080:8080 --cpus=0.1 --memory=0.1g ${IMG}


##@ Test
.PHONY: test
test: test-unit test-api ## Run tests.


.PHONY: test-unit
test-unit:
	go test ./... -coverprofile _output/cover.out

.PHONY: test-api
test-api:
	./scripts/tests/api.sh

.PHONY: test-api-wrk
test-api-wrk:
	wrk -t4 -c300 -d10s -T5s --latency http://127.0.0.1:8080/healthz

