include env.mk

git-push:
	git add . ; git commit -m "updating pipeline"; git push

pipeline: git-push
	git add .; git commit -m "Pipeline WIP"; git push
	fly -t dev set-pipeline \
		-n -p $(APP) \
		-c cicd/pipeline.yml \
		-l $(HOME)/.ssh/ci-credentials.yml \
		-v git_repo_url=git@github.com:$(NAMESPACE)/$(APP).git \
		-v git_repo=$(APP)

	fly -t dev unpause-pipeline -p $(APP)

#
#	fly -t dev watch -j $(APP)/go-template-engine
.PHONY: pipeline

pipeline-destroy:
	fly -t dev destroy-pipeline -p $(APP)
.PHONY: pipeline-destroy

pipeline-login:
	fly -t dev login -n dev -c https://ci.correia.io

lint:
	@go fmt -x $$(glide nv)
.PHONY: lint

deps:
	glide install
.PHONY: deps

build:
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_FILE)-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_FILE)-darwin-amd64
.PHONY: build

test:
	go test $$(glide nv)
.PHONY: test

clean:
	rm -rf ./bin/* ./dist/*
.PHONY: clean

# Concourse targets
_deps: _prepare
	cd $(GOPATH)/src/github.com/$(NAMESPACE)/$(APP); glide install
.PHONY: _deps


_build: _prepare _deps
	cd $(GOPATH)/src/github.com/$(NAMESPACE)/$(APP); GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_FILE)
.PHONY: _build

_test: _prepare
	cd $(GOPATH)/src/github.com/$(NAMESPACE)/$(APP);	go test $$(glide nv)
.PHONY: _test

_clean:
	cd $(GOPATH)/src/github.com/$(NAMESPACE)/$(APP); rm -rf ./bin/* ./dist/*
.PHONY: _clean

package: _prepare
	@[ -f ./package ] && echo dist folder found, skipping creation || mkdir -p ./package
	cd $(GOPATH)/src/github.com/$(NAMESPACE)/$(APP) ; GOPATH=/go make deps lint test build tar
	cp -Rv $(GOPATH)/src/github.com/$(NAMESPACE)/$(APP)/dist/* ../package/
.PHONY: package

_prepare:
	@[ -f $(GOPATH)/src/github.com/$(NAMESPACE)/$(APP) ] && echo $(GOPATH)/src/github.com/$(NAMESPACE)/$(APP) folder found, skipping creation || mkdir -p $(GOPATH)/src/github.com/$(NAMESPACE)/$(APP)

tar:
	@[ -f ./dist ] && echo dist folder found, skipping creation || mkdir -p ./dist
	tar -cvzf ./dist/$(APP)-linux-amd64.tar.gz -C ./bin .
.PHONY: tar

create-make:
	echo '#!/usr/bin/env bash\n' > make.sh
	echo 'dir=$$(dirname $$0)' >> make.sh
	echo 'cd $$dir' >> make.sh
	echo 'make $$1' >> make.sh

version:
	@git show version:version