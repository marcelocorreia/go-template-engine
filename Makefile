include env.mk ci.mk pipeline.mk

test-full: clean_docs _docs-check cover-tests cover-out cover-html cover-cleanup

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

lint:
	@go fmt -x $$(glide nv)
.PHONY: lint

deps:
	@glide install
.PHONY: deps

test:
	@go test $$(glide nv) -cover -v
.PHONY: test

clean_full: clean_bin clean_dist clean_docs

clean_bin:
	@rm -rf ./bin/*

clean_dist:
	@rm -rf ./dist/*

clean_docs:
	@rm -rf ./docs/*

build:
	$(call build,GOOS=$(GOOS) GOARCH=$(GOARCH),$(APP_NAME))

define build
	$1 go build -o ./bin/$(APP_NAME) -ldflags "-X main.VERSION=dev" -v
endef

_validate-version:
ifndef VERSION
	$(error VERSION is required)
endif
_validate-file:
ifndef FILE
	$(error FILE is required)
endif

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

release: build_all test-full package

homebrew-tap:
	go-template-engine \
		--source ci/go-template-engine.rb \
		--var dist_file=dist/go-template-engine-darwin-amd64-1.39.0.zip \
		--var version=1.39.0 \
		--var hash_sum=123 \
		  > /Users/marcelo/IdeaProjects/tardis/homebrew-taps/go-template-engine.rb


get-version:
	@git checkout origin/version -- version && \
		cat version && \
		rm version

_docs-check:
	@[ -f $(DOCS_DIR) ] && echo $(DOCS_DIR) folder found || mkdir -p $(DOCS_DIR)



