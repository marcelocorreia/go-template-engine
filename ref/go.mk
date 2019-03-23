cover-tests:
	@go test . -coverprofile docs/main-cover.out -v
	@$(foreach var,$(shell /go/bin/glide nv | sed 's/\.//g' | sed 's/\///g' ),go test ./$(var)/... -coverprofile docs/$(var)-cover.out || exit 1;)

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

open_coverage:
	@open docs/index.html

build_all: clean_dist
	@gox \
		-ldflags "-X main.VERSION=$(VERSION)" \
		--arch amd64 \
		--output ./dist/{{.Dir}}-{{.OS}}-{{.Arch}}-$(VERSION)/{{.Dir}}
.PHONY: build_all


DISTDIRS=$(shell ls dist/)
package: build_all
	for dir in $(DISTDIRS); do \
       cd dist/$$dir/; \
       zip ../$$dir.zip * ; \
       cd -;\
       rm -rf dist/$$dir/;\
    done
.PHONY: package

release: build_all  package

lint:
	@go fmt -x $$(glide nv)
.PHONY: lint

deps:
	@glide install
.PHONY: deps

test:
	@go test $$(glide nv) -cover -v
.PHONY: test

all-versions: ## Show all versions and the commit
	@git ls-remote --tags $(GIT_REMOTE)

current-version: _setup-versions## Show the current version.
	@echo $(CURRENT_VERSION)

next-version: _setup-versions## Show the current version.
	@echo $(NEXT_VERSION)

GIT_BRANCH ?= master
GIT_REMOTE ?= origin
RELEASE_TYPE ?= patch
_setup-versions:
	$(eval export CURRENT_VERSION=$(shell git ls-remote --tags $(GIT_REMOTE) | grep -v latest | awk '{ print $$2}'|grep -v 'stable'| sort -r --version-sort | head -n1|sed 's/refs\/tags\///g'))
	$(eval export NEXT_VERSION=$(shell docker run --rm --entrypoint=semver $(SEMVER_DOCKER) -c -i $(RELEASE_TYPE) $(CURRENT_VERSION)))

d: _setup-versions
	echo "docker run --rm --entrypoint=semver $(SEMVER_DOCKER) -c -i $(RELEASE_TYPE) $(CURRENT_VERSION)"