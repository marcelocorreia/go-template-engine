include env.mk pipeline.mk go.mk

test-full: clean_docs _docs-check cover-tests cover-out cover-html cover-cleanup

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



