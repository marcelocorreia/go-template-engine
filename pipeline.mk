git-push:
	git add . ; git commit -m "updating pipeline"; git push

pipeline-full: git-push pipeline

_prepare:
	@echo $(GOPATH) - $(shell pwd)
	@mkdir -p /go/src/$(NAMESPACE)/$(APP_NAME)/dist
	@cp -R * /go/src/$(NAMESPACE)/$(APP_NAME)/
	@$(call ci_make,deps)

_build:
	@$(call ci_make,lint build GOOS=linux)

_test:
	@$(call ci_make, test GOOS=linux)

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

concourse-up: _ci-params
	cd ci && docker-compose up -d

concourse-logs:
	cd ci && docker-compose logs -f
concourse-down:
	cd ci && docker-compose kill; docker-compose down

_ci-params:
	@$(eval export CONCOURSE_EXTERNAL_URL=$(CONCOURSE_EXTERNAL_URL))

define concourse
	cd ci && docker-compose $1
endef
