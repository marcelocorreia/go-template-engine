APP_NAME := go-template-engine
GITHUB_USER := marcelocorreia
GOARCH :=amd64
GOOS := darwin
GOPATH := /go
NAMESPACE := github.com/marcelocorreia
OUTPUT_FILE := ./bin/$(APP_NAME)
REPO_NAME := $(APP_NAME)
REPO_URL := git@github.com:$(GITHUB_USER)/$(APP_NAME).git
TEST_OUTPUT_DIR := tmp
#VERSION := $(shell make get-last-release)
WORKDIR := $(GOPATH)/src/$(NAMESPACE)/$(REPO_NAME)
HOMEBREW_REPO := git@github.com:marcelocorreia/homebrew-taps.git
HOMEBREW_BINARY := dist/$(APP_NAME)-darwin-amd64-$(VERSION).zip
#HOMEBREW_BINARY_SUM := $(shell shasum -a 256 $(HOMEBREW_BINARY) | awk '{print $$1}')
HOMEBREW_REPO_PATH ?= /Users/marcelo/IdeaProjects/tardis/homebrew-taps
DOCS_DIR := docs

include go.mk



pipeline: git-push
	fly -t main set-pipeline \
		-n -p $(APP_NAME) \
		-c ./ci/pipeline.yml \
		-l $(HOME)/.ssh/ci-credentials.yml \
		-l ci/properties.yml

	fly -t main unpause-pipeline -p $(APP_NAME)
.PHONY: pipeline

test-full: clean_docs _docs-check cover-tests cover-out cover-html cover-cleanup

clean_full: clean_bin clean_dist clean_docs

view-doc:
	grip -b

clean_bin:
	@rm -rf ./bin/*

clean_dist:
	@rm -rf ./dist/*

clean_docs:
	@rm -rf ./docs/*

build:
	$(call build,GOOS=$(GOOS) GOARCH=$(GOARCH),$(APP_NAME))

define build
	$1 go build -o ./bin/$(APP_NAME) -ldflags "-X main.VERSION=dev-mc2" -v ./main.go
endef

_validate-version:
ifndef VERSION
	$(error VERSION is required)
endif
_validate-file:
ifndef FILE
	$(error FILE is required)
endif

APP_NAME := go-template-engine
GITHUB_USER := marcelocorreia

get-last-release:
	@curl -s https://api.github.com/repos/$(GITHUB_USER)/$(APP_NAME)/tags | jq ".[]|.name" | head -n1 | sed 's/\"//g' | sed 's/v*//g'


get-version:
	@git checkout origin/version -- version && \
		cat version && \
		rm version

_docs-check:
	@[ -f $(DOCS_DIR) ] && echo $(DOCS_DIR) folder found || mkdir -p $(DOCS_DIR)


_validate-app-name:
ifndef APP_NAME
	$(error APP_NAME is required)
endif

update_brew: _validate-app-name
	./update-brew.sh $(APP_NAME)

git-push:
	git add . ; git commit -m "updating pipeline"; git push

pipeline-full: git-push pipeline

_prepare:
	@echo $(GOPATH) - $(shell pwd)
	@mkdir -p /go/src/$(NAMESPACE)/$(APP_NAME)/dist
	@cp -R * /go/src/$(NAMESPACE)/$(APP_NAME)/
	@$(call ci_make,deps)

_build:
	@$(call ci_make,lint build GOOS=linux)

_test:
	@$(call ci_make, test GOOS=linux)

_release: _validate-version
	@$(call ci_make,release)
	pwd
	cp $(GOPATH)/src/$(NAMESPACE)/$(APP_NAME)/dist/*zip ../output/

define ci_make
	echo ""
	echo "*** $1::Begin ***"
	cd $(GOPATH)/src/$(NAMESPACE)/$(APP_NAME) && \
    		make $1
	echo "*** $1::End ***"
	echo ""
	cd -
endef
