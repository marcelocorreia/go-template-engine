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
	fly -t dev trigger-job -j $(APP)/integration
#
#	fly -t dev watch -j $(APP)/go-template-engine
.PHONY: pipeline

pipeline-destroy:
	fly -t dev destroy-pipeline -p $(APP)
.PHONY: pipeline-destroy

lint:
	@go fmt -x $$(glide nv)
.PHONY: lint

deps:
	glide install
.PHONY: deps

build:
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_FILE)
.PHONY: build

test:
	go test $$(glide nv)
.PHONY: test


clean:
	rm -rf ./bin/* ./dist/*
.PHONY: clean

# Concourse targets
_deps: _prepare
	cd /go/src/github.com/$(NAMESPACE)/$(APP); glide install
.PHONY: _deps


_build: _prepare _deps
	cd /go/src/github.com/$(NAMESPACE)/$(APP); GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_FILE)
.PHONY: _build

_test:
	cd /go/src/github.com/$(NAMESPACE)/$(APP);	go test $$(glide nv)
.PHONY: _test

_clean:
	cd /go/src/github.com/$(NAMESPACE)/$(APP); rm -rf ./bin/* ./dist/*
.PHONY: _clean

package: _prepare
	cd /go/src/github.com/$(NAMESPACE)/$(APP) ; GOPATH=/go make deps lint test build tar
	cp -Rv /go/src/github.com/$(NAMESPACE)/$(APP)/dist/* ../package/
.PHONY: package

_prepare:
	@[ -f /go/src/github.com/$(NAMESPACE)/$(APP) ] && echo /go/src/github.com/$(NAMESPACE)/$(APP) folder found, skipping creation || mkdir -p /go/src/github.com/$(NAMESPACE)/$(APP); rsync -avz --exclude 'vendor' ./* /go/src/github.com/$(NAMESPACE)/$(APP)/

tar:
	@[ -f ./dist ] && echo dist folder found, skipping creation || mkdir -p ./dist
	tar -cvzf ./dist/$(APP)-linux-amd64.tar.gz -C ./bin .
.PHONY: tar
