# Auto generated
M := $(shell printf "\033[34;1m▶\033[0m")
#
PROJECT_HOME := $(shell pwd)
PROJECT_NAME ?= go-template-engine
AWS_PROFILE ?= aws-profile
#AWS_ACCESS_KEY_ID ?= AWS_ACCESS_KEY_ID
#AWS_SECRET_ACCESS_KEY ?= AWS_SECRET_ACCESS_KEY
AWS_DEFAULT_REGION ?= ap-southeast-2
GITHUB_USER ?= marcelocorreia
GIT_REPO_NAME ?= go-template-engine
SEMVER_DOCKER ?= marcelocorreia/semver
HAMMER_CMD := hammer
RELEASE_TYPE ?= minor

#
.PHONY: default
default: hammer-targets


fmt:
	go fmt ./awstools ./cmd/go-template-engine ./templateengine

wrap-up:
	go mod tidy
	go mod vendor
	go mod download

# Builds the application
go-build: fmt
	go build -o ./bin/$(PROJECT_NAME) -ldflags "-X main.VERSION=dev" -v ./cmd/$(PROJECT_NAME)/

# Tests the application
go-test:
	go test -v ./...

# Starts Go Doc Server
godoc-server-start:  ;$(info $(M) - Starting Godoc server)
	godoc -v

# Stops Go Doc Server
godoc-server-stop:  ;$(info $(M) - Stopping Godoc server)
	killall godoc

# Open Go Doc Package page
godoc-open: ;$(info $(M) - Opening Godoc page)
	open http://localhost:6060/pkg/github.com/$(GITHUB_USER)/$(PROJECT_NAME)

# Shows all versions
all-versions: ;$(info $(M) - Showing $(PROJECT_NAME) all versions)
	@git ls-remote --tags $(GIT_REMOTE)

# Shows current version
current-version: _setup-versions
	@echo $(CURRENT_VERSION)

# Shows next  version
next-version: _setup-versions
	@echo $(NEXT_VERSION)

# Builds a snapshot
go-snapshot: fmt ;$(info $(M) - Releasing $(PROJECT_NAME)-snapshot)
	-@mkdir -p dist coverage
	goreleaser  release --snapshot  --rm-dist --debug

# Releases the application
go-release: fmt _require-github-token _setup-versions tag-push ;$(info $(M) - Releasing $(PROJECT_NAME)-$(NEXT_VERSION))
	goreleaser release  --rm-dist

# Builds a dry run of the app
go-release-dry-run: fmt
	goreleaser release  --skip-publish

go-reporter:
	goreporter -p . -r goreporter/ -e vendor -f html

# Prepares for release
tag-push: _setup-versions ;$(call git_push,Releasing $(PROJECT_NAME)"-"$(NEXT_VERSION)) ;$(info $(M) Tagging $(PROJECT_NAME)-$(NEXT_VERSION))
	git tag $(NEXT_VERSION)
	git tag v$(NEXT_VERSION)
	git tag go/v$(NEXT_VERSION)
	git push --tags

# Prepares for release
_setup-versions:
	$(eval export CURRENT_VERSION=$(shell git ls-remote --tags $(GIT_REMOTE) | grep -v latest | awk '{ print $$2}'|grep -v 'stable'| sort -r --version-sort | head -n1|sed 's/refs\/tags\///g'))
	$(eval export NEXT_VERSION=$(shell docker run --rm --entrypoint=semver $(SEMVER_DOCKER) -c -i $(RELEASE_TYPE) $(CURRENT_VERSION)))

# Dep Support - begin
# Runs dep init
dep-init: ;$(info $(M) - Running "dep init")
	dep init

# Runs dep ensure
dep-ensure: ;$(info $(M) - Running "dep dep-ensure")
	dep ensure

# Runs dep ensure -update
dep-ensure-update: ;$(info $(M) - Running "dep dep-ensure")
	dep ensure -update

dep-ensure-add: ;$(info $(M) - Running "dep dep-ensure -add $(PACKAGE)")
ifndef PACKAGE
	$(error PACKAGE is required)
endif
	dep ensure -add $(PACKAGE)


# Opens coverage page using default browser
open-coverage: ;$(info $(M) - Opening $(PROJECT_NAME)-$(NEXT_VERSION) Test Coverage Report)
	open ./coverage/index.html

# Opens github page using default browser
open-page: ;$(info $(M) - Opening $(PROJECT_NAME) Github Page)
	open https://github.com/$(GITHUB_USER)/$(GIT_REPO_NAME).git

# Opens Go Report Card
go-report: ;$(info $(M) - Opening $(PROJECT_NAME) Go Report Card)
	open https://goreportcard.com/report/github.com/$(GITHUB_USER)/$(GIT_REPO_NAME)

# Opens README page in local browser using Github renderer. Requires grip. pip install grip
grip:
	grip -b

# Exports AWS creds environment variables. Profiles are recommended but not mandatory
_aws_init:
ifdef AWS_ACCESS_KEY_ID
	$(eval export AWS_ACCESS_KEY_ID=$(AWS_ACCESS_KEY_ID))
	$(eval export AWS_SECRET_ACCESS_KEY=$(AWS_SECRET_ACCESS_KEY))
endif
ifdef AWS_PROFILE
	$(eval export AWS_DEFAULT_REGION=$(AWS_DEFAULT_REGION))
	$(eval export AWS_PROFILE=$(AWS_PROFILE))
endif
#

define git_push
	-git add .
	-git commit -m "$1"
	-git push
endef

_require-github-token:
ifndef GITHUB_TOKEN
	$(error GITHUB_TOKEN is required)
endif

# Hammer auto generated target
hammer-banner:
	-@hammer minion banner $(AWS_PROFILE)

# Hammer auto generated target
hammer-forge-make:
	-@hammer forge addon --name Makefile.tpl .

# Hammer auto generated target
hammer-forge-readme:
	-@hammer forge addon --name README.tpl.md .

# Hammer auto generated target
hammer-parasameters:
	@make variables | egrep -v ':=|SHELL|MAKEFLAGS'

# Hammer auto generated target
hammer-targets:
	@make -npRq | egrep -i -v 'makefile|^#|=|^\t|^\.|->|^_' | grep ":" | sort | uniq | awk '{print $$1}'|sed 's/://g'

# Hammer auto generated target
hammer-variables:
	@make -pn | grep -A1 '^# makefile'| egrep -v '^--|^define|^#' | sort | uniq

hammer-doctor:
	hammer make doctor .

_test-bin:
	go-template-engine -s /Volumes/work/go/src/github.com/marcelocorreia/go-template-engine/templateengine/test_fixtures/config/dev/app1 \
	-o /Volumes/work/go/src/github.com/marcelocorreia/go-template-engine/tmp \
	--var-file /Volumes/work/go/src/github.com/marcelocorreia/go-template-engine/templateengine/test_fixtures/config/dev/config.yaml \
	--log debug
