
# ==============================================================================
# Options

DOCKER := docker
DOCKER_SUPPORTED_API_VERSION ?= 1.31

REGISTRY_PREFIX ?= wukongsun
#BASE_IMAGE = centos:centos8
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


# ==============================================================================
# Targets

.PHONY: docker.build
docker.build: $(addprefix docker.build., $(addprefix $(IMAGE_PLAT)., $(IMAGES)))

.PHONY: docker.build.multiarch
docker.build.multiarch: $(foreach p,$(PLATFORMS),$(addprefix docker.build., $(addprefix $(p)., $(IMAGES))))

.PHONY: docker.build.%
docker.build.%: go.build.%
	$(eval IMAGE := $(COMMAND))
	$(eval IMAGE_PLAT := $(subst _,/,$(PLATFORM)))
	@echo "===========> Building Docker image $(IMAGE) $(VERSION) for $(IMAGE_PLAT)"
	@mkdir -p $(TMP_DIR)/$(IMAGE)
	@cat $(ROOT_DIR)/build/docker/$(IMAGE)/Dockerfile\
		| sed "s#BASE_IMAGE#$(BASE_IMAGE)#g" >$(TMP_DIR)/$(IMAGE)/Dockerfile
	@cp $(OUTPUT_DIR)/platforms/$(IMAGE_PLAT)/$(IMAGE) $(TMP_DIR)/$(IMAGE)/
	@DST_DIR=$(TMP_DIR)/$(IMAGE) $(ROOT_DIR)/build/docker/$(IMAGE)/build.sh 2>/dev/null || true
	$(eval BUILD_SUFFIX := $(_DOCKER_BUILD_EXTRA_ARGS) --pull -t $(REGISTRY_PREFIX)/$(IMAGE)-$(ARCH):$(VERSION) $(TMP_DIR)/$(IMAGE))
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
docker.push.%: docker.build.%
	@echo "===========> Pushing Docker image $(IMAGE) $(VERSION) to $(REGISTRY_PREFIX)"
	$(DOCKER) push $(REGISTRY_PREFIX)/$(IMAGE)-$(ARCH):$(VERSION)

