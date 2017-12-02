APP_NAME=go-template-engine
GITHUB_USER=marcelocorreia
GOARCH?=amd64
GOOS?=darwin
GOPATH?=/go
NAMESPACE=github.com/marcelocorreia
OUTPUT_FILE=./bin/$(APP_NAME)
REPO_NAME=$(APP_NAME)
REPO_URL=git@github.com:$(GITHUB_USER)/$(APP_NAME).git
TEST_OUTPUT_DIR=tmp
VERSION?=$(shell make get-version)
WORKDIR=$(GOPATH)/src/$(NAMESPACE)/$(REPO_NAME)