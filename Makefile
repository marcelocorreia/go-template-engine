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
RELEASE_TYPE ?= patch

# ####
.PHONY: default
default: available-targets

available-targets: _available-targets
dep-ensure-update: _dep-ensure-update
dep-ensure: _dep-ensure
dep-init: _dep-init
doc-readme: _readme
go-build: _go-build
go-test: _go-test
go-release: _release
go-report: _go-report
go-snapshot: _snapshot
go-open-coverage: _open-coverage
github-open-page: _open-page
git-push: ;$(call git_push,Updating...)
godoc-open: _godoc-open
godoc-server-start: _godoc-server-start
godoc-server-stop: _godoc-server-stop
grip-open: _grip
version-all: _all-versions
version-current: _all-current
version-next: _all-next
hammer-forge-all: hammer-forge-make hammer-forge-readme
hammer-forge-readme:
	-@hammer forge addon --name README.tpl.md .
hammer-forge-make:
	-@hammer forge addon --name Makefile.tpl .

# --
# Please do not call the targets below directly, instead create a call as the ones above.
# --

# Builds the application
_go-build:
	go build -o ./bin/$(PROJECT_NAME) -ldflags "-X main.VERSION=dev" -v ./cmd/$(PROJECT_NAME)/

# Tests the application
_go-test:
	go test -v ./...

# Starts Go Doc Server
_godoc-server-start:  ;$(info $(M) - Starting Godoc server)
	godoc -v

# Stops Go Doc Server
_godoc-server-stop:  ;$(info $(M) - Stopping Godoc server)
	killall godoc

# Open Go Doc Package page
_godoc-open: ;$(info $(M) - Opening Godoc page)
	open http://localhost:6060/pkg/github.com/$(GITHUB_USER)/$(PROJECT_NAME)

# Shows all versions
_all-versions: ;$(info $(M) - Showing $(PROJECT_NAME) all versions)
	@git ls-remote --tags $(GIT_REMOTE)

# Shows current version
_current-version: _setup-versions
	@echo $(CURRENT_VERSION)

# Shows next  version
_next-version: _setup-versions
	@echo $(NEXT_VERSION)

# Builds a snapshot
_snapshot: ;$(info $(M) - Releasing $(PROJECT_NAME)-snapshot)
	-@mkdir -p dist coverage
	goreleaser  release --snapshot  --rm-dist --debug

# Releases the application
_release: _require-github-token _setup-versions _tag-push ;$(info $(M) - Releasing $(PROJECT_NAME)-$(NEXT_VERSION))
	goreleaser release  --rm-dist

# Builds a dry run of the app
_dry-run:
	goreleaser release  --skip-publish

# Prepares for release
_tag-push: ;$(call git_push,Releasing $(PROJECT_NAME)"-"$(NEXT_VERSION)) ;$(info $(M) Tagging $(PROJECT_NAME)-$(NEXT_VERSION))
	git tag $(NEXT_VERSION)
	git push --tags

# Prepares for release
_setup-versions:
	$(eval export CURRENT_VERSION=$(shell git ls-remote --tags $(GIT_REMOTE) | grep -v latest | awk '{ print $$2}'|grep -v 'stable'| sort -r --version-sort | head -n1|sed 's/refs\/tags\///g'))
	$(eval export NEXT_VERSION=$(shell docker run --rm --entrypoint=semver $(SEMVER_DOCKER) -c -i $(RELEASE_TYPE) $(CURRENT_VERSION)))

# Dep Support - begin
# Runs dep init
_dep-init: ;$(info $(M) - Running "dep init")
	dep init

# Runs dep ensure
_dep-ensure: ;$(info $(M) - Running "dep dep-ensure")
	dep ensure

# Runs dep ensure -update
_dep-ensure-update: ;$(info $(M) - Running "dep dep-ensure")
	dep ensure -update

_dep-ensure-add: ;$(info $(M) - Running "dep dep-ensure -add $(PACKAGE)")
ifndef PACKAGE
	$(error PACKAGE is required)
endif
	dep ensure -add $(PACKAGE)


# Opens coverage page using default browser
_open-coverage: ;$(info $(M) - Opening $(PROJECT_NAME)-$(NEXT_VERSION) Test Coverage Report)
	open ./coverage/index.html

# Opens github page using default browser
_open-page: ;$(info $(M) - Opening $(PROJECT_NAME) Github Page)
	open https://github.com/$(GITHUB_USER)/$(GIT_REPO_NAME).git

# Generates README.md file, refer to README.yml
_readme: ;$(info $(M) - Generates $(PROJECT_NAME) README.md Page)
	$(HAMMER_CMD) generate --resource-type readme .

# Opens Go Report Card
_go-report: ;$(info $(M) - Opening $(PROJECT_NAME) Go Report Card)
	open https://goreportcard.com/report/github.com/$(GITHUB_USER)/$(GIT_REPO_NAME)

# Opens README page in local browser using Github renderer. Requires grip. pip install grip
_grip:
	grip -b

# Shows available targets
_available-targets:
	@make -pn -C $(PWD) | grep -A1 '^# makefile'| egrep -v '^--|^define|^#' | sort | uniq

# Exports AWS creds environment variables. Profiles are recommended but not mandatory
_set-aws-creds:
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