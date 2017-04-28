include env.mk

pipeline:
	git add . ; git commit -m "lazy dev"; git push
	fly -t dev set-pipeline \
		-n -p $(APP) \
		-c cicd/pipeline.yml \
		-l $(HOME)/.ssh/ci-credentials.yml \
		-v git_repo_url=git@github.com:$(NAMESPACE)/$(APP).git \
		-v s3_regex="go-template-engine/go-template-engine-(.*).tar.gz" \
		-v git_repo=$(APP)

	fly -t dev unpause-pipeline -p $(APP)

	fly -t dev trigger-job -j $(APP)/go-template-engine

	fly -t dev watch -j $(APP)/go-template-engine
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
	cp -Rv /go/src/github.com/$(NAMESPACE)/$(APP)/dist/* ../package/
.PHONY: package


tar:
	@[ -f ./dist ] && echo dist folder found, skipping creation || mkdir -p ./dist
	tar -cvzf ./dist/$(APP)-linux-amd64.tar.gz -C ./bin .
.PHONY: tar
