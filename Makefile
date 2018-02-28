include env.mk pipeline.mk go.mk

test-full: clean_docs _docs-check cover-tests cover-out cover-html cover-cleanup

clean_full: clean_bin clean_dist clean_docs

view-doc:
	grip -b

clean_bin:
	@rm -rf ./bin/*

clean_dist:
	@rm -rf ./dist/*

clean_docs:
	@rm -rf ./docs/*

build:
	$(call build,GOOS=$(GOOS) GOARCH=$(GOARCH),$(APP_NAME))

define build
	$1 go build -o ./bin/$(APP_NAME) -ldflags "-X main.VERSION=dev" -v
endef

_validate-version:
ifndef VERSION
	$(error VERSION is required)
endif
_validate-file:
ifndef FILE
	$(error FILE is required)
endif

APP_NAME := go-template-engine
GITHUB_USER := marcelocorreia

get-last-release:
	@curl -s https://api.github.com/repos/$(GITHUB_USER)/$(APP_NAME)/tags | jq ".[]|.name" | head -n1 | sed 's/\"//g' | sed 's/v*//g'

homebrew-tap:
	go-template-engine \
		--source ci/go-template-engine.rb \
		--var dist_file=dist/go-template-engine-darwin-amd64-1.39.0.zip \
		--var version=1.39.0 \
		--var hash_sum=123 \
		  > /Users/marcelo/IdeaProjects/tardis/homebrew-taps/go-template-engine.rb


get-version:
	@git checkout origin/version -- version && \
		cat version && \
		rm version

_docs-check:
	@[ -f $(DOCS_DIR) ] && echo $(DOCS_DIR) folder found || mkdir -p $(DOCS_DIR)

concourse-pull:
	cd ci && docker-compose pull
concourse-up:
	cd ci && CONCOURSE_EXTERNAL_URL=http://localhost:8080 docker-compose up -d

concourse-down:
	cd ci && docker-compose down

concourse-stop:
	cd ci && docker-compose stop

concourse-start:
	cd ci && docker-compose start

concourse-logs:
	cd ci && docker-compose logs -f

concourse-keys:
	@[ -f ./ci/keys ] && echo ./ci/keys folder found || $(call create-concourse-keys)

define create-concourse-keys
	echo "Creating Concourse keys"
	mkdir -p ./ci/keys/web ./ci/keys/worker;
	ssh-keygen -t rsa -f ./ci/keys/web/tsa_host_key -N ''
	ssh-keygen -t rsa -f ./ci/keys/web/session_signing_key -N ''
	ssh-keygen -t rsa -f ./ci/keys/worker/worker_key -N ''
	cp ./ci/keys/worker/worker_key.pub ./ci/keys/web/authorized_worker_keys
	cp ./ci/keys/web/tsa_host_key.pub ./ci/keys/worker
endef
