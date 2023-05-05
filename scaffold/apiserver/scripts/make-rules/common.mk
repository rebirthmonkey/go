SHELL := /bin/bash


# ==============================================================================
# Directory options

# include the common make file
COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/../.. && pwd -P))
endif

ifeq ($(origin OUTPUT_DIR),undefined)
OUTPUT_DIR := $(ROOT_DIR)/_output
$(shell mkdir -p $(OUTPUT_DIR))
endif

ifeq ($(origin TOOLS_DIR),undefined)
TOOLS_DIR := $(OUTPUT_DIR)/tools
$(shell mkdir -p $(TOOLS_DIR))
endif

ifeq ($(origin TMP_DIR),undefined)
TMP_DIR := $(OUTPUT_DIR)/tmp
$(shell mkdir -p $(TMP_DIR))
endif


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
# Build options

ROOT_PACKAGE=$(shell pwd)
ifeq ($(ROOT_PACKAGE),)
	$(error the variable ROOT_PACKAGE must be set prior to including golang.mk)
endif

VERSION_PACKAGE=github.com/rebirthmonkey/go/pkg/version
# set the version number. you should not need to do this
# for the majority of scenarios.
ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --tags --always --match='v*')
endif

COMMANDS ?= $(filter-out %.md, $(wildcard ${ROOT_DIR}/cmd/*))
ifeq (${COMMANDS},) # like /Users/xxx/go/scaffold/apiserver/cmd/apiserver
  $(error Could not determine COMMANDS, set ROOT_DIR or run in source dir)
endif

BINS ?= $(foreach cmd,${COMMANDS},$(notdir ${cmd}))
ifeq (${BINS},) # like apiserver, authz
  $(error Could not determine BINS, set ROOT_DIR or run in source dir)
endif

# The OS must be linux when building docker images
PLATFORMS ?= linux_amd64
# PLATFORMS ?= darwin_amd64 windows_amd64 linux_amd64 linux_arm64

PLATFORM := linux_amd64
IMAGE_PLAT := linux_amd64

# Linux command settings
FIND := find . ! -path './third_party/*' ! -path './vendor/*'
XARGS := xargs --no-run-if-empty

# Makefile settings
ifndef V
MAKEFLAGS += --no-print-directory
endif

# ==============================================================================
# Test options

# Minimum test coverage
ifeq ($(origin COVERAGE),undefined)
COVERAGE := 60
endif


# ==============================================================================
# Git options

# Check if the tree is dirty.  default to dirty
GIT_TREE_STATE:="dirty"
ifeq (, $(shell git status --porcelain 2>/dev/null))
	GIT_TREE_STATE="clean"
endif
GIT_COMMIT:=$(shell git rev-parse HEAD)

# Copy githook scripts when execute makefile
#COPY_GITHOOK:=$(shell cp -f githooks/* .git/hooks/)


# ==============================================================================
# Tools options

# Specify tools severity, include: BLOCKER_TOOLS, CRITICAL_TOOLS, TRIVIAL_TOOLS.
# Missing BLOCKER_TOOLS can cause the CI flow execution failed, i.e. `make all` failed.
# Missing CRITICAL_TOOLS can lead to some necessary operations failed. i.e. `make release` failed.
# TRIVIAL_TOOLS are Optional tools, missing these tool have no affect.
BLOCKER_TOOLS ?= gsemver golines go-junit-report golangci-lint addlicense goimports codegen
CRITICAL_TOOLS ?= swagger mockgen gotests git-chglog github-release coscmd go-mod-outdated protoc-gen-go cfssl go-gitlint
TRIVIAL_TOOLS ?= depth go-callvis gothanks richgo rts kube-score coscli


# ==============================================================================
# Format options

COMMA := ,
SPACE :=
SPACE +=
