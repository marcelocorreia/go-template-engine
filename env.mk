APP=go-template-engine
GOPATH?=/go
REPO_NAME=go-template-engine
OUTPUT_FILE=./bin/$(APP)
DOCKER_WORKING_DIR=$(GOPATH)/src/github.com/marcelocorreia/$(REPO_NAME)
NAMESPACE=marcelocorreia
IMAGE_GO_GLIDE=marcelocorreia/go-glide-builder:latest
TEST_OUTPUT_DIR=tmp
