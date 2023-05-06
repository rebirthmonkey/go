
# ==============================================================================
# Options

DOCKER := docker
DOCKER_SUPPORTED_API_VERSION ?= 1.31

REGISTRY_PREFIX ?= wukongsun
BASE_IMAGE = golang:1.18

EXTRA_ARGS ?= --no-cache
_DOCKER_BUILD_EXTRA_ARGS :=

ifdef HTTP_PROXY
_DOCKER_BUILD_EXTRA_ARGS += --build-arg HTTP_PROXY=${HTTP_PROXY}
endif

ifneq ($(EXTRA_ARGS), )
_DOCKER_BUILD_EXTRA_ARGS += $(EXTRA_ARGS)
endif

# Determine image files by looking into build/docker/*/Dockerfile
IMAGES_DIR ?= $(wildcard ${ROOT_DIR}/build/docker/*)

# Determine images names by stripping out the dir names
IMAGES ?= $(filter-out tools,$(foreach image,${IMAGES_DIR},$(notdir ${image})))
ifeq (${IMAGES},)
  $(error Could not determine IMAGES, set ROOT_DIR or run in source dir)
endif

IMAGE_VERSION := v1.0.0


# ==============================================================================
# Targets

.PHONY: docker.build
docker.build: $(addprefix docker.build., $(addprefix $(IMAGE_PLAT)., $(IMAGES)))

.PHONY: docker.build.multiarch
docker.build.multiarch: $(foreach p,$(PLATFORMS),$(addprefix docker.build., $(addprefix $(p)., $(IMAGES))))

.PHONY: docker.build.%
docker.build.%:
	$(eval IMAGE := $(word 2, $(subst ., ,$*)))
	$(eval ARCH := $(word 2, $(subst _, ,$(word 1, $(subst ., ,$*)))))
	$(eval IMAGE_OS := $(word 1, $(subst _, ,$(word 1, $(subst ., ,$*)))))
	@echo "===========> Building Docker image $(REGISTRY_PREFIX)/$(IMAGE)-$(ARCH):$(IMAGE_VERSION) for $(IMAGE_OS)"
	@mkdir -p $(TMP_DIR)/$(IMAGE)
	@cat $(ROOT_DIR)/build/docker/$(IMAGE)/Dockerfile\
		| sed "s#BASE_IMAGE#$(BASE_IMAGE)#g" >$(TMP_DIR)/$(IMAGE)/Dockerfile
	@cp $(OUTPUT_DIR)/platforms/$(IMAGE_OS)/$(ARCH)/$(IMAGE) $(TMP_DIR)/$(IMAGE)/
	@DST_DIR=$(TMP_DIR)/$(IMAGE) $(ROOT_DIR)/build/docker/$(IMAGE)/build.sh 2>/dev/null || true
	$(eval BUILD_SUFFIX := $(_DOCKER_BUILD_EXTRA_ARGS) --pull -t $(REGISTRY_PREFIX)/$(IMAGE)-$(ARCH):$(IMAGE_VERSION) $(TMP_DIR)/$(IMAGE))
	@if [ $(shell $(GO) env GOARCH) != $(ARCH) ] ; then \
		$(DOCKER) build --platform $(IMAGE_PLAT) $(BUILD_SUFFIX) ; \
	else \
		$(DOCKER) build $(BUILD_SUFFIX) ; \
	fi
	@rm -rf $(TMP_DIR)/$(IMAGE)

.PHONY: docker.push
docker.push: $(addprefix docker.push., $(addprefix $(IMAGE_PLAT)., $(IMAGES)))

.PHONY: docker.push.multiarch
docker.push.multiarch: $(foreach p,$(PLATFORMS),$(addprefix docker.push., $(addprefix $(p)., $(IMAGES))))

.PHONY: docker.push.%
docker.push.%:
	$(eval IMAGE := $(word 2, $(subst ., ,$*)))
	$(eval ARCH := $(word 2, $(subst _, ,$(word 1, $(subst ., ,$*)))))
	@echo "===========> Pushing Docker image $(IMAGE)-$(ARCH):$(IMAGE_VERSION) to $(REGISTRY_PREFIX)"
	$(DOCKER) push $(REGISTRY_PREFIX)/$(IMAGE)-$(ARCH):$(IMAGE_VERSION)

.PHONY: docker.clean
docker.clean: $(addprefix docker.clean., $(addprefix $(IMAGE_PLAT)., $(IMAGES)))

.PHONY: docker.clean.%
docker.clean.%:
	$(eval IMAGE := $(word 2, $(subst ., ,$*)))
	$(eval ARCH := $(word 2, $(subst _, ,$(word 1, $(subst ., ,$*)))))
	$(eval IMAGE_OS := $(word 1, $(subst _, ,$(word 1, $(subst ., ,$*)))))
	@echo "===========> Cleaning Docker image $(REGISTRY_PREFIX)/$(IMAGE)-$(ARCH):$(IMAGE_VERSION)"
	$(DOCKER) image rm $(REGISTRY_PREFIX)/$(IMAGE)-$(ARCH):$(IMAGE_VERSION)