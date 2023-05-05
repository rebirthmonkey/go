GO := go
GO_SUPPORTED_VERSIONS ?= 1.13|1.14|1.15|1.16|1.17|1.18|1.19|1.20
GO_LDFLAGS += -X $(VERSION_PACKAGE).GitVersion=$(VERSION) \
	-X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PACKAGE).GitTreeState=$(GIT_TREE_STATE) \
	-X $(VERSION_PACKAGE).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

ifneq ($(DLV),)
	GO_BUILD_FLAGS += -gcflags "all=-N -l"
	LDFLAGS = ""
endif
GO_BUILD_FLAGS += -ldflags "$(GO_LDFLAGS)"

ifeq ($(GOOS),windows)
	GO_OUT_EXT := .exe
endif

GOPATH := $(shell go env GOPATH)
ifeq ($(origin GOBIN), undefined)
	GOBIN := $(GOPATH)/bin
endif

EXCLUDE_TESTS=github.com/rebirthmonkey/go/test github.com/rebirthmonkey/go/pkg/log github.com/rebirthmonkey/go/third_party github.com/rebirthmonkey/go/internal/pump/storage github.com/rebirthmonkey/go/internal/pump github.com/rebirthmonkey/go/internal/pkg/logger


# ==============================================================================
# Targets

.PHONY: go.build.%
go.build.%:
	$(eval BIN := $(word 2,$(subst ., ,$*)))
	$(eval OS := $(word 1,$(subst _, ,$(PLATFORM))))
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	@echo "===========> Building binary $(BIN) $(VERSION) for $(OS) $(ARCH)"
	@mkdir -p $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)
	@CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) $(GO) build $(GO_BUILD_FLAGS) -o $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)/$(BIN)$(GO_OUT_EXT) $(ROOT_PACKAGE)/cmd/$(BIN)

.PHONY: go.build
go.build: $(addprefix go.build., $(addprefix $(PLATFORM)., $(BINS)))

.PHONY: go.build.multiarch
go.build.multiarch: $(foreach p,$(PLATFORMS),$(addprefix go.build., $(addprefix $(p)., $(BINS))))

.PHONY: go.clean
go.clean:
	@echo "===========> Cleaning all build output"
	@-rm -vrf $(OUTPUT_DIR)

.PHONY: go.run.%
go.run.%:
	@echo "===========> Running the application $*"
	@$(GO) run cmd/$*/$*.go -c configs/$*.yaml --mysql.host 0.0.0.0:3306 --gin.secure.tls.cert-file ./configs/cert/server.pem --gin.secure.tls.private-key-file ./configs/cert/server.key


.PHONY: go.test.unit
go.test.unit:
	@echo "===========> Running unit test"
	@set -o pipefail;$(GO) test -race -cover -coverprofile=$(OUTPUT_DIR)/coverage.out \
		-timeout=10m -shuffle=on -short -v `go list ./...|\
		egrep -v $(subst $(SPACE),'|',$(sort $(EXCLUDE_TESTS)))` 2>&1 | \
		tee >(go-junit-report --set-exit-code >$(OUTPUT_DIR)/report.xml)
	@#sed -i '/mock_.*.go/d' $(OUTPUT_DIR)/coverage.out # remove mock_.*.go files from test coverage
	@#$(GO) tool cover -html=$(OUTPUT_DIR)/coverage.out -o $(OUTPUT_DIR)/coverage.html

