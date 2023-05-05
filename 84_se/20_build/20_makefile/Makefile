IMG ?= xxx:yyy


.PHONY: all
all: build


##@ Build
.PHONY: build
build:
	go build -o bin/manager main.go

.PHONY: run
run:
	go run ./main.go


##@ Docker
.PHONY: docker-build
docker-build:
	docker build -t ${IMG} .

.PHONY: docker-push
docker-push:
	docker push ${IMG}

.PHONY: docker-run
docker-run:
	docker run -d -p 8080:8080 ${IMG}

.PHONY: docker-run-limit
docker-run-limit: ## Run docker image with the server.
	docker run -d -p 8080:8080 --cpus=0.1 --memory=0.1g ${IMG}


##@ Test
.PHONY: test
test: test.unit test.api

.PHONY: test-unit
test-unit:
	go test ./... -coverprofile _output/cover.out

.PHONY: test-api
test-api:
	./scripts/tests/api.sh

.PHONY: test-api-wrk
test-api-wrk:
	wrk -t4 -c300 -d10s -T5s --latency http://127.0.0.1:8080/healthz

