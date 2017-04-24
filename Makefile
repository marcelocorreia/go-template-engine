APP=go-template-engine
GOPATH?=/go
REPO_NAME=go-template-engine
OUTPUT_FILE=./bin/$(APP)
DOCKER_WORKING_DIR=$(GOPATH)/src/github.com/marcelocorreia/$(REPO_NAME)
NAMESPACE=marcelocorreia
IMAGE_GO_GLIDE=marcelocorreia/go-glide-builder:latest
TEST_OUTPUT_DIR=tmp

pipeline:
	git add . ; git commit -m "lazy dev"; git push
	fly -t dev set-pipeline \
		-n -p $(APP) \
		-c cicd/pipeline.yml \
		-l $(HOME)/.ssh/ci-credentials.yml \
		-v git_repo_url=git@github.com:$(NAMESPACE)/$(APP).git

	fly -t dev unpause-pipeline -p $(APP)

	fly -t dev trigger-job -j $(APP)/package


.PHONY: pipeline


default: deps

lint:
	@go fmt -x $$(glide nv)
.PHONY: lint

deps:
	glide install
.PHONY: deps

build:
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_FILE)
.PHONY: _build

test:
	go test $$(glide nv)
.PHONY: test

clean:
	rm -rf ./bin/* ./dist/*
.PHONY: clean

package:
	mkdir -p /go/src/github.com/$(NAMESPACE)/$(APP)
	rsync -avz --exclude 'vendor' ./* /go/src/github.com/$(NAMESPACE)/$(APP)/
	cd /go/src/github.com/$(NAMESPACE)/$(APP) ; GOPATH=/go make clean deps lint test build tar
	cd /go/src/github.com/$(NAMESPACE)/$(APP); ls -l dist
.PHONY: package

list:
	cd /go/src/github.com/$(NAMESPACE)/$(APP); ls -l
.PHONY: list


tar:
	@[ -f ./dist ] && echo dist folder found, skipping creation || mkdir -p ./dist
	tar -cvzf ./dist/$(APP)-linux-amd64.tar.gz -C ./bin .
	@[ -f ../dist ] && echo dist folder found, skipping creation || mkdir -p ../dist
	cp ./dist/* ../dist/
.PHONY: tar


