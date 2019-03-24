include config.mk

release: _release

all-versions:
	@git ls-remote --tags $(GIT_REMOTE)

current-version: _setup-versions
	@echo $(CURRENT_VERSION)

next-version: _setup-versions
	@echo $(NEXT_VERSION)

build: _build

build_all: _build_all

#####
#tests: _setup-versions cover-tests cover-out cover-html

_build: _setup-versions
	@go fmt -x $$(glide nv)
	@export GOOS=$(GOOS) GOARCH=$(GOARCH) && \
		go build -o ./bin/$(APP_NAME) -ldflags "-X main.VERSION=$(CURRENT_VERSION)-dev" -v ./main.go


_build_all: _setup-versions
	@gox -ldflags "-X main.VERSION=$(NEXT_VERSION)" \
		--arch amd64 \
		--output ./dist/{{.Dir}}-{{.OS}}-{{.Arch}}-$(NEXT_VERSION)/{{.Dir}}

_package:
	@for dir in $(DISTDIRS); do \
		echo $$dir; \
		if [[ -d "dist/$$dir" ]];then \
			cd dist/$$dir/; \
		   zip ../$$dir.zip * ; \
		   cd -;\
		   rm -rf dist/$$dir/;\
		fi \
    done

_xota:
		echo yay;\
	fi


#define _zipIt
#endef



_release: _setup-versions _build_all _package _git-push ;$(info $(M) Releasing version $(NEXT_VERSION)...)## Release by adding a new tag. RELEASE_TYPE is 'patch' by default, and can be set to 'minor' or 'major'.
	@github-release release \
		-u marcelocorreia \
		-r go-template-engine \
		--tag $(NEXT_VERSION) \
		--name $(NEXT_VERSION) \
		--description "Template engine em Golang full of goodies"
	@$(foreach plat,$(PLATFORMS),github-release upload -u marcelocorreia -r go-template-engine --tag $(NEXT_VERSION) --name go-template-engine-$(plat)-amd64-$(NEXT_VERSION).zip --file ./dist/go-template-engine-$(plat)-amd64-$(NEXT_VERSION).zip;)
	@make _update_brew
	@make _docker-build
	@make _docker-push

_setup-versions:
	$(eval export CURRENT_VERSION=$(shell git ls-remote --tags $(GIT_REMOTE) | grep -v latest | awk '{ print $$2}'|grep -v 'stable'| sort -r --version-sort | head -n1|sed 's/refs\/tags\///g'))
	$(eval export NEXT_VERSION=$(shell docker run --rm --entrypoint=semver $(SEMVER_DOCKER) -c -i $(RELEASE_TYPE) $(CURRENT_VERSION)))

cover-tests:
	@go test . -coverprofile docs/main-cover.out -v
	@$(foreach var,$(shell glide nv | sed 's/\.//g' | sed 's/\///g'),go test ./$(var)/... -coverprofile docs/$(var)-cover.out || exit 1;)

cover-out:
	@echo "mode: set" > docs/coverage.out
	@$(foreach f,$(shell ls docs/**out),cat $(f) | sed 's/mode: set//g' | perl -p -e 's/^\s*$$//mg' >> docs/coverage.out || exit 1;)

cover-html:
	@go tool cover -html=docs/coverage.out -o docs/index.html
	@$(foreach f,$(shell ls docs/**out),go tool cover -html=$(f) -o $(f).html  || exit 1;)
	@rm docs/coverage.out.html

cover-cleanup:
	-@mkdir docs/out
	@$(foreach f,$(shell ls docs/**out),$(shell echo mv $(f) docs/out/)  || exit 1;)

_docker-build: _setup-versions
	sed -i .bk 's/ARG gte.*/ARG gte_version\=\"$(CURRENT_VERSION)\"/' resources/Dockerfile
	docker build -t marcelocorreia/go-template-engine -f resources/Dockerfile .
	docker build -t marcelocorreia/go-template-engine:$(CURRENT_VERSION) -f resources/Dockerfile .

_docker-push: _setup-versions
	docker push marcelocorreia/go-template-engine
	docker push marcelocorreia/go-template-engine:$(CURRENT_VERSION)

_git-push:
	-@git add .
	-@git commit -m "Release $(NEXT_VERSION)"
	-@git push


_update_brew: _setup-versions
	-rm -rf /tmp/homebrew-gte
	git clone git@github.com:marcelocorreia/homebrew-taps.git /tmp/homebrew-gte
	/Volumes/work/go/src/github.com/marcelocorreia/go-template-engine/bin/go-template-engine -s resources/go-template-engine.rb \
		--var hash_sum=$(shell shasum -a 256 dist/go-template-engine-darwin-amd64-$(CURRENT_VERSION).zip | awk {'print $$1'}) \
		--var version=$(CURRENT_VERSION) \
		--var dist_file=go-template-engine-darwin-amd64-$(CURRENT_VERSION).zip > \
		/tmp/homebrew-gte/go-template-engine.rb

	cd /tmp/homebrew-gte && \
		git add go-template-engine.rb && \
		git commit -m "Release go-template-engine $(CURRENT_VERSION)" \
		&& git push

_clean_bin:
	@rm -rf ./bin/*

#_clean_all: _clean_bin
#	@rm -rf ./dist/*
#	@rm -rf ./docs/*