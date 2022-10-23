# Build all by default, even if it's not first
.DEFAULT_GOAL := all

.PHONY: all
all: tools tidy clean build

# ==============================================================================
# Build options

ROOT_PACKAGE=~/workspace/go/84_se/20_build/30_make/80_server
VERSION_PACKAGE=github.com/rebirthmonkey/go/pkg/version

# ==============================================================================
# Includes

include scripts/make-rules/common.mk # make sure include common.mk at the first include line
include scripts/make-rules/tools.mk
include scripts/make-rules/golang.mk
include scripts/make-rules/image.mk
include scripts/make-rules/deploy.mk
#include scripts/make-rules/copyright.mk
#include scripts/make-rules/release.mk


# ==============================================================================
# Usage

define USAGE_OPTIONS

Options:
  DEBUG            Whether to generate debug symbols. Default is 0.
  BINS             The binaries to build. Default is all of cmd.
                   This option is available when using: make build/build.multiarch
                   Example: make build BINS="iam-apiserver iam-authz-server"
  IMAGES           Backend images to make. Default is all of cmd starting with iam-.
                   This option is available when using: make image/image.multiarch/push/push.multiarch
                   Example: make image.multiarch IMAGES="iam-apiserver iam-authz-server"
  REGISTRY_PREFIX  Docker registry prefix. Default is marmotedu. 
                   Example: make push REGISTRY_PREFIX=ccr.ccs.tencentyun.com/marmotedu VERSION=v1.6.2
  PLATFORMS        The multiple platforms to build. Default is linux_amd64 and linux_arm64.
                   This option is available when using: make build.multiarch/image.multiarch/push.multiarch
                   Example: make image.multiarch IMAGES="iam-apiserver iam-pump" PLATFORMS="linux_amd64 linux_arm64"
  VERSION          The version information compiled into binaries.
                   The default is obtained from gsemver or git.
  V                Set to 1 enable verbose build. Default is 0.
endef
export USAGE_OPTIONS

# ==============================================================================
# Targets

## deploy: Deploy updated components to development env.
.PHONY: deploy
deploy: image
	@$(MAKE) deploy.run

## undeploy: undeploy updated components to development env.
.PHONY: undeploy
undeploy:
	@$(MAKE) deploy.clean

## image: Build docker images for host arch.
.PHONY: image
image:
	@$(MAKE) image.build

## push: Build docker images for host arch and push images to registry.
.PHONY: push
push:
	@$(MAKE) image.push


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

## run: Run the application.
.PHONY: run
run:
	@$(MAKE) go.run

## test: Run unit test.
.PHONY: test
test:
	@$(MAKE) go.test

## test: Run unit test.
.PHONY: test.api
test.api:
	@./test/api/test.sh api::test::user

## tidy: run go mod tidy
.PHONY: tidy
tidy:
	@$(GO) mod tidy


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
