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

release: clean_full
	make package GOOS=linux VERSION=$(VERSION)
	make package GOOS=darwin VERSION=$(VERSION)
	make package GOOS=windows VERSION=$(VERSION)
	make clean_bin
	echo "<><><><><><><><><>"
	pwd
	echo "<><><><><><><><><>"
	ls -lR ../
	echo "<><><><><><><><><>"
	ls -l /
	echo "<><><><><><><><><>"
build:
	$(call build,GOOS=$(GOOS) GOARCH=$(GOARCH),tardis)

package: clean_bin lint build
	 $(call package,$(APP_NAME),$(GOOS),$(GOARCH),$(VERSION))

define package
	echo "Creating Package $(shell pwd)/dist/$1-$2-$3-$4.tar.gz"
	tar -cvzf ./dist/$1-$2-$3-$4.tar.gz -C ./bin .
endef

define build
	$1 go build -o ./bin/$2 -ldflags "-X main.VERSION=$(VERSION)" -v
endef

_validate-version:
ifndef VERSION
	$(error VERSION is required)
endif