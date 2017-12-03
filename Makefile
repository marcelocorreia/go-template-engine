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

package: clean_dist
	@gox -ldflags "-X main.VERSION=$(VERSION)" \
		--arch amd64 --arch arm \
		--output ./dist/{{.Dir}}-{{.OS}}-{{.Arch}}-$(VERSION)/{{.Dir}}
.PHONY: package

DISTDIRS=$(shell ls dist/)
release: package
	for dir in $(DISTDIRS) ; do \
       cd dist/$$dir/; \
       tar -cvzf ../$$dir.tar.gz * ; \
       cd -;\
       rm -rf dist/$$dir/;\
    done
.PHONY: release



homebrew-tap:
	@go-template-engine \
		--source ci/go-template-engine.rb \
        --var dist_file=$(HOMEBREW_BINARY) \
        --var version=$(VERSION) \
        --var hash_sum=$(HOMEBREW_BINARY_SUM) \
        > homebrew-repo/go-template-engine.rb
#	@cd /tmp/brew-repo && ls -l && pwd && git push

get-version:
	@git checkout origin/version -- version && \
		cat version && \
		rm version