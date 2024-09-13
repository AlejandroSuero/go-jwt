ifeq ($(OS),Windows_NT)
	GREEN= [00;32m
	RESTORE= [0m
	GO_FILES=$(shell dir /b /s *.go | findstr /v /i /e ".git")
else
	GREEN="\033[00;32m"
	RESTORE="\033[0m"
	GO_FILES=$(shell find . -name "*.go" -and  -not -name ".git")
endif

# make the output of the message appear green
define style_calls
	$(eval $@_msg = $(1))
	echo ${GREEN}${$@_msg}${RESTORE}
endef

default_target: help

build: release_jwt

release_jwt: $(GO_FILES)
	@$(call style_calls,"Building jwt release")
	@go build -o bin/go-jwt -ldflags "-s -w" .
	@$(call style_calls,"Done ✅")
	@echo ""
.PHONY: release_jwt

lint: $(GO_FILES) spell
	@$(call style_calls,"Running golangci-lint")
	@golangci-lint run
	@$(call style_calls,"Done ✅")
	@echo ""
	@$(call style_calls, "Running gofmt")
	@gofmt -s -w .
	@$(call style_calls,"Done ✅")
	@echo ""
	@$(call style_calls, "Running markdownlint")
	@markdownlint .
	@$(call style_calls,"Done ✅")
	@echo ""
	@$(call style_calls, "Running yamllint")
	@yamllint .
	@$(call style_calls,"Done ✅")
	@echo ""
.PHONY: lint

spell:
	@$(call style_calls,"Running codespell check")
	@codespell --quiet-level=2 --check-hidden --skip=./.git .
	@$(call style_calls,"Done ✅")
	@echo ""
.PHONY: spell

spell-write:
	@$(call style_calls,"Running codespell write")
	@codespell --quiet-level=2 --check-hidden --skip=./.git --write-changes .
	@$(call style_calls,"Done ✅")
	@echo ""
.PHONY: spell-write

all: lint spell build

help:
	@echo "Available targets:"
	@echo "  build: builds the jwt release"
	@echo "  release_jwt: builds the jwt release"
	@echo "  help: shows this help message"
.PHONY: help
