include env.mk ci.mk pipeline.mk

lint:
	@go fmt -x $$(glide nv)
.PHONY: lint

deps:
	glide install
.PHONY: deps

test:
	go test $$(glide nv) -v
.PHONY: test

clean_full: clean_bin clean_dist

clean_bin:
	rm -rf ./bin/*

clean_dist:
	rm -rf ./dist/*

build:
	$(call build,GOOS=$(GOOS) GOARCH=$(GOARCH),$(APP_NAME))

define build
	$1 go build -o ./bin/$2 -ldflags "-X main.VERSION=dev" -v
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
		-parallel 2 \
		-ldflags "-X main.VERSION=$(VERSION)" \
		--arch amd64 \
		--output ./dist/{{.Dir}}-{{.OS}}-{{.Arch}}-$(VERSION)/{{.Dir}}
.PHONY: build_all


DISTDIRS=$(shell ls dist/)
package: build_all
	for dir in $(DISTDIRS) ; do \
       cd dist/$$dir/; \
       tar -cvzf ../$$dir.tar.gz * ; \
       cd -;\
       rm -rf dist/$$dir/;\
    done
.PHONY: package

release: build_all package homebrew-tap

homebrew-tap:
	@go-template-engine \
		--source ci/go-template-engine.rb \
        --var dist_file=$(HOMEBREW_BINARY) \
        --var version=$(VERSION) \
        --var hash_sum=$(HOMEBREW_BINARY_SUM) \
        > $(HOMEBREW_REPO_PATH)/go-template-engine.rb
#	@cd /tmp/brew-repo && ls -l && pwd && git push

get-version:
	@git checkout origin/version -- version && \
		cat version && \
		rm version