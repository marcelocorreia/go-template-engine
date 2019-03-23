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
	$1 go build -o ./bin/$(APP_NAME) -ldflags "-X main.VERSION=$(CURRENT_VERSION)" -v ./main.go
endef

DISTDIRS=$(shell ls dist/)
build_all: _setup-versions 
	gox -ldflags "-X main.VERSION=$(NEXT_VERSION)" \
		--arch amd64 \
		--output ./dist/{{.Dir}}-{{.OS}}-{{.Arch}}-$(NEXT_VERSION)/{{.Dir}}
package:
	-@for dir in $(DISTDIRS); do \
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

release: _release 

_release: _setup-versions build_all package _git-push _release-warning _setup-versions ;$(info $(M) Releasing version $(NEXT_VERSION)...)## Release by adding a new tag. RELEASE_TYPE is 'patch' by default, and can be set to 'minor' or 'major'.
	github-release release \
		-u marcelocorreia \
		-r go-template-engine \
		--tag $(NEXT_VERSION) \
		--name $(NEXT_VERSION) \
		--description "Template engine em Golang full of goodies"

	github-release upload -u marcelocorreia -r go-template-engine --tag $(NEXT_VERSION) --name go-template-engine-darwin-amd64-$(NEXT_VERSION).zip --file ./dist/go-template-engine-darwin-amd64-$(NEXT_VERSION).zip
	github-release upload -u marcelocorreia -r go-template-engine --tag $(NEXT_VERSION) --name go-template-engine-darwin-amd64-$(NEXT_VERSION).zip --file ./dist/go-template-engine-freebsd-amd64-$(NEXT_VERSION).zip
	github-release upload -u marcelocorreia -r go-template-engine --tag $(NEXT_VERSION) --name go-template-engine-darwin-amd64-$(NEXT_VERSION).zip --file ./dist/go-template-engine-linux-amd64-$(NEXT_VERSION).zip
	github-release upload -u marcelocorreia -r go-template-engine --tag $(NEXT_VERSION) --name go-template-engine-darwin-amd64-$(NEXT_VERSION).zip --file ./dist/go-template-engine-netbsd-amd64-$(NEXT_VERSION).zip
	github-release upload -u marcelocorreia -r go-template-engine --tag $(NEXT_VERSION) --name go-template-engine-darwin-amd64-$(NEXT_VERSION).zip --file ./dist/go-template-engine-openbsd-amd64-$(NEXT_VERSION).zip
	github-release upload -u marcelocorreia -r go-template-engine --tag $(NEXT_VERSION) --name go-template-engine-darwin-amd64-$(NEXT_VERSION).zip --file ./dist/go-template-engine-windows-amd64-$(NEXT_VERSION).zip


_release-warning: ;$(info $(M) Release - Warning...)
	@cowsay -f mario "Make sure evertyhing is pushed"
	


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

concourse-up: _ci-params
	$(call concourse,up -d)

concourse-logs: _ci-params
	$(call concourse,logs -f)

concourse-down: _ci-params
	$(call concourse,kill)
	$(call concourse,down)

peido-%:
	echo $@

_ci-params:
	@$(eval export CONCOURSE_EXTERNAL_URL=$(CONCOURSE_EXTERNAL_URL))

define concourse
	cd ci && docker-compose $1
endef

_git-push:
	-@git add .
	-@git commit -m "Release $(NEXT_VERSION)"
	-@git push