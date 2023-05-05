.PHONY: all
all: clean build


# ==============================================================================
# Includes

include scripts/make-rules/common.mk
include scripts/make-rules/golang.mk
include scripts/make-rules/docker.mk
include scripts/make-rules/k8s.mk
include scripts/make-rules/tools.mk


# ==============================================================================
# Build

.PHONY: build
build:
	@echo "===========> Building Go source code for host platform"
	@$(MAKE) go.build

.PHONY: build.multiarch
build.multiarch:
	@echo "===========> Building Go source code for all platforms"
	@$(MAKE) go.build.multiarch

.PHONY: clean
clean:
	@echo "===========> Cleaning all build output"
	@-rm -vrf $(OUTPUT_DIR)

.PHONY: run
run:
	@echo "===========> Running Go source code"
	@$(MAKE) go.run


# ==============================================================================
# Docker & k8s

.PHONY: docker-build
docker-build:
	@$(MAKE) docker.build

.PHONY: docker-push
docker-push:
	@$(MAKE) docker.push

.PHONY: k8s-deploy
k8s-deploy:
	@$(MAKE) k8s.deploy

.PHONY: k8s-undeploy
k8s-undeploy:
	@$(MAKE) k8s.undeploy


# ==============================================================================
# Test

.PHONY: test-unit
test-unit:
	@$(MAKE) go.test.unit

## test: Run unit test.
.PHONY: test-api
test-api:
	@./test/api/test.sh api::test::user

.PHONY: test-api-wrk
test-api-wrk:
	@wrk -t4 -c300 -d10s -T5s --latency http://127.0.0.1:30080/healthz


# ==============================================================================
# Tools & Help

.PHONY: tools
tools:
	@$(MAKE) tools.install

.PHONY: help
help: Makefile
	@echo -e "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"
