.PHONY: all
all: clean build


# ==============================================================================
# Includes

include scripts/make-rules/common.mk # make sure include common.mk at the first include line
include scripts/make-rules/tools.mk
include scripts/make-rules/golang.mk
include scripts/make-rules/docker.mk
include scripts/make-rules/k8s.mk


# ==============================================================================
# Targets

## build: Build source code for host platform.
.PHONY: build
build:
	@$(MAKE) go.build

## build.multiarch: Build source code for multiple platforms. See option PLATFORMS.
.PHONY: build.multiarch
build.multiarch:
	@$(MAKE) go.build.multiarch

## clean: Remove all files that are created by building.
.PHONY: clean
clean:
	@echo "===========> Cleaning all build output"
	@-rm -vrf $(OUTPUT_DIR)

## run
.PHONY: run
run:
	@$(MAKE) go.run

## Docker image
.PHONY: docker-build
docker-build:
	@$(MAKE) docker.build

## push
.PHONY: docker-push
docker-push:
	@$(MAKE) docker.push


## deploy
.PHONY: k8s-deploy
k8s-deploy: docker-build
	@$(MAKE) k8s.deploy

## undeploy
.PHONY: k8s-undeploy
k8s-undeploy:
	@$(MAKE) k8s.undeploy


## test: Run unit test.
.PHONY: test-unit
test-unit:
	@$(MAKE) go.test.unit

## test: Run unit test.
.PHONY: test-api
test-api:
	@./test/api/test.sh api::test::user


## tools: install dependent tools.
.PHONY: tools
tools:
	@$(MAKE) tools.install


## help: Show this help info.
.PHONY: help
help: Makefile
	@echo -e "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"
