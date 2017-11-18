APP_NAME=go-template-engine
GOPATH?=/go
REPO_NAME=$(APP_NAME)
OUTPUT_FILE=./bin/$(APP_NAME)
GITHUB_USER=marcelocorreia
NAMESPACE=github.com/marcelocorreia
REPO_URL=git@github.com:$(GITHUB_USER)/$(APP_NAME).git
TEST_OUTPUT_DIR=tmp
WORKDIR=$(GOPATH)/src/$(NAMESPACE)/$(REPO_NAME)
VERSION?=0.0.0
VERSION?=$(shell cat version)
GOOS?=darwin
GOARCH?=amd64

