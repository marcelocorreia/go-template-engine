fly-login:
	fly -t dev login -n dev -c https://ci.correia.io

git-push:
	git add . ; git commit -m "updating pipeline"; git push

pipeline-full: git-push pipeline

pipeline:
	fly -t dev set-pipeline \
		-n -p $(APP_NAME) \
		-c ./ci/pipeline.yml \
		-l $(HOME)/.ssh/ci-credentials.yml \
		-l ci/properties.yml

	fly -t dev unpause-pipeline -p $(APP_NAME)
.PHONY: pipeline

pipeline-destroy:
	fly -t dev destroy-pipeline -p $(APP_NAME)
.PHONY: pipeline-destroy

_prepare:
	@echo $(GOPATH) - $(shell pwd)
	@mkdir -p /go/src/$(NAMESPACE)/$(APP_NAME)/dist
	@cp -R * /go/src/$(NAMESPACE)/$(APP_NAME)/
	@$(call ci_make,deps)

_build:
	@$(call ci_make,lint build GOOS=linux)

_test:
	@$(call ci_make,lint test GOOS=linux)

_release: _validate-version
	@$(call ci_make,release)
	pwd
	cp $(GOPATH)/src/$(NAMESPACE)/$(APP_NAME)/dist/*zip ../output/

define ci_make
	echo ""
	echo "*** $1::Begin ***"
	cd $(GOPATH)/src/$(NAMESPACE)/$(APP_NAME) && \
    		make $1
	echo "*** $1::End ***"
	echo ""
	cd -
endef