include env.mk

lint:
	@go fmt -x $$(glide nv)
.PHONY: lint

deps:
	@glide install
.PHONY: deps

_prepare:
	@echo $(GOPATH) - $(shell pwd)
	@mkdir -p /go/src/$(NAMESPACE)/$(APP_NAME)/dist
	@cp -R * /go/src/$(NAMESPACE)/$(APP_NAME)/
	@$(call ci_make,deps)

_build:
	@$(call ci_make,lint build GOOS=linux)

define ci_make
	echo ""
	echo "*** $1::Begin ***"
	cd $(GOPATH)/src/$(NAMESPACE)/$(APP_NAME) && \
    		make $1
	echo "*** $1::End ***"
	echo ""
endef

merda:
	@echo "MERDA"
	@echo " MERDA"
	@echo "  MERDA"
	@echo "   MERDA"
	@echo "    MERDA"
	@echo "     MERDA"
	@echo "    MERDA"
	@echo "   MERDA"
	@echo "  MERDA"
	@echo " MERDA"
	@echo "MERDA"
	@echo "MERDA"
	@echo " MERDA"
	@echo "  MERDA"
	@echo "   MERDA"
	@echo "    MERDA"
	@echo "     MERDA"
	@echo "    MERDA"
	@echo "   MERDA"
	@echo "  MERDA"
	@echo " MERDA"
	@echo "MERDA"
	@echo "MERDA"
	@echo " MERDA"
	@echo "  MERDA"
	@echo "   MERDA"
	@echo "    MERDA"
	@echo "     MERDA"
	@echo "    MERDA"
	@echo "   MERDA"
	@echo "  MERDA"
	@echo " MERDA"
	@echo "MERDA"
	@echo "MERDA"
	@echo " MERDA"
	@echo "  MERDA"
	@echo "   MERDA"
	@echo "    MERDA"
	@echo "     MERDA"
	@echo "    MERDA"
	@echo "   MERDA"
	@echo "  MERDA"
	@echo " MERDA"
	@echo "MERDA"

