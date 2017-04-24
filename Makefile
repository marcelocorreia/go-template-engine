APP=go-template-engine
GOPATH?=/go
REPO_NAME=go-template-engine
OUTPUT_FILE=./bin/$(APP)
DOCKER_WORKING_DIR=$(GOPATH)/src/github.com/marcelocorreia/$(REPO_NAME)
NAMESPACE=marcelocorreia
IMAGE_GO_GLIDE=marcelocorreia/go-glide-builder:latest
TEST_OUTPUT_DIR=tmp

pipeline:
	fly -t dev set-pipeline \
		-n -p $(APP) \
		-c cicd/pipeline.yml \
		-l $(HOME)/.ssh/ci-credentials.yml \
		-v git_repo_url=git@github.com:$(NAMESPACE)/$(APP).git

	fly -t dev unpause-pipeline -p $(APP)

.PHONY: pipeline


default: deps

lint:
	@go fmt -x $$(glide nv)
.PHONY: lint

deps:
	glide install
.PHONY: deps

build:
	$(call docker_build,_build)
.PHONY: build

_build:
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_FILE)
.PHONY: _build

test:
	go test $$(glide nv)
.PHONY: test

clean:
	rm -rf ./bin/* ./dist/*
.PHONY: clean

package:
	$(call docker_build,_package)
	$(call create_tar)
.PHONY: package

_package: clean deps lint _build
	$(call create_tar)
.PHONY: _package

define docker_build
	@docker run --rm \
		-v $(PWD):$(DOCKER_WORKING_DIR) \
		-w $(DOCKER_WORKING_DIR) \
		$(IMAGE_GO_GLIDE) \
		bash -c "uname -a; make $1" && chown $(USER): .* *
endef

define create_tar
    @[ -f ./dist/linux ] && echo dist folder found, skipping creation || mkdir -p ./dist/linux
    tar -cvzf ./dist/linux/$(APP)-linux-amd64.tar.gz -C ./bin .
endef