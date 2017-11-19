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
	cp $(GOPATH)/src/$(NAMESPACE)/$(APP_NAME)/dist/*tar.gz output/

define ci_make
	echo ""
	echo "*** $1::Begin ***"
	cd $(GOPATH)/src/$(NAMESPACE)/$(APP_NAME) && \
    		make $1
	echo "*** $1::End ***"
	echo ""
	cd -

endef