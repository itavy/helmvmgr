GO           ?= go
FIRST_GOPATH := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))
GOOPTS       ?=

GO_VERSION        ?= $(shell $(GO) version)
GO_VERSION_NUMBER ?= $(word 3, $(GO_VERSION))

GIT          ?= git

GIT_COMMIT_HASH ?= $(word 1, $(shell $(GIT) show-ref))
# TODO extract from git tag
BUILD_VERSION ?= 0.0.1
BUILD_TIME ?= $(shell date +%FT%T%z)

.PHONY: all
all: test build

.PHONY: test
test:
	@echo ">> TODO testing"

.PHONY: build
build:
	$(GO) build \
	-ldflags="-X helmvmgr/pkg/cmdcli.Version=$(BUILD_VERSION) \
	-X helmvmgr/pkg/cmdcli.BuildTime=$(BUILD_TIME) \
	-X helmvmgr/pkg/cmdcli.CommitHash=$(GIT_COMMIT_HASH)"
