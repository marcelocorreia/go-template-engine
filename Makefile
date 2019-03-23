APP_NAME := go-template-engine
GITHUB_USER := marcelocorreia
GOARCH := amd64
GOOS := darwin
GOPATH := /go
NAMESPACE := github.com/marcelocorreia
OUTPUT_FILE := ./bin/$(APP_NAME)
REPO_NAME := $(APP_NAME)
REPO_URL := git@github.com:$(GITHUB_USER)/$(APP_NAME).git
TEST_OUTPUT_DIR := tmp
#VERSION := $(shell make get-version)
VERSION := 2.5.8
WORKDIR := $(GOPATH)/src/$(NAMESPACE)/$(REPO_NAME)
HOMEBREW_REPO := git@github.com:marcelocorreia/homebrew-taps.git
HOMEBREW_BINARY := dist/$(APP_NAME)-darwin-amd64-$(VERSION).zip
#HOMEBREW_BINARY_SUM := $(shell shasum -a 256 $(HOMEBREW_BINARY) | awk '{print $$1}')
HOMEBREW_REPO_PATH ?= /Users/marcelo/IdeaProjects/tardis/homebrew-taps
DOCS_DIR := docs
CONCOURSE_EXTERNAL_URL ?= http://localhost:8080
SEMVER_DOCKER ?= marcelocorreia/semver
SEMVER_DOCKER ?= marcelocorreia/semver
GIT_BRANCH ?= master
GIT_REMOTE ?= origin
RELEASE_TYPE ?= patch

build:
	$(call build,GOOS=$(GOOS) GOARCH=$(GOARCH),$(APP_NAME))

define build
	$1 go build -o ./bin/$(APP_NAME) -ldflags "-X main.VERSION=dev" -v ./main.go
endef

DISTDIRS=$(shell ls dist/)
build_all: package
	gox -ldflags "-X main.VERSION=$(VERSION)" \
		--arch amd64 \
		--output ./dist/{{.Dir}}-{{.OS}}-{{.Arch}}-$(VERSION)/{{.Dir}}
package:
	for dir in $(DISTDIRS); do \
    	cd dist/$$dir/; \
    	zip ../$$dir.zip * ; \
        cd -;\
        rm -rf dist/$$dir/;\
    done

lint:
	@go fmt -x $$(glide nv)
.PHONY: lint

all-versions: ## Show all versions and the commit
	@git ls-remote --tags $(GIT_REMOTE)

current-version: _setup-versions## Show the current version.
	@echo $(CURRENT_VERSION)

next-version: _setup-versions## Show the current version.
	@echo $(NEXT_VERSION)


_setup-versions:
	$(eval export CURRENT_VERSION=$(shell git ls-remote --tags $(GIT_REMOTE) | grep -v latest | awk '{ print $$2}'|grep -v 'stable'| sort -r --version-sort | head -n1|sed 's/refs\/tags\///g'))
	$(eval export NEXT_VERSION=$(shell docker run --rm --entrypoint=semver $(SEMVER_DOCKER) -c -i $(RELEASE_TYPE) $(CURRENT_VERSION)))

cover-tests:
	@go test . -coverprofile docs/main-cover.out -v
	@$(foreach var,$(shell glide nv | sed 's/\.//g' | sed 's/\///g' ),go test ./$(var)/... -coverprofile docs/$(var)-cover.out || exit 1;)

cover-out:
	@echo "mode: set" > docs/coverage.out
	@$(foreach f,$(shell ls docs/**out),cat $(f) | sed 's/mode: set//g' | perl -p -e 's/^\s*$$//mg' >> docs/coverage.out || exit 1;)

cover-html:
	@go tool cover -html=docs/coverage.out -o docs/index.html
	@$(foreach f,$(shell ls docs/**out),go tool cover -html=$(f) -o $(f).html  || exit 1;)
	@rm docs/coverage.out.html

cover-cleanup:
	@mkdir docs/out
	@$(foreach f,$(shell ls docs/**out),$(shell echo mv $(f) docs/out/)  || exit 1;)

docker-build:
	docker build -t marcelocorreia/go-template-engine .
