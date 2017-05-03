APP=sample
pipeline:
	fly -t dev set-pipeline \
		-n -p $(APP) \
		-l $(HOME)/.ssh/ci-credentials.yml \
		-v git_repo_url=git@github.com:marcelocorreia/go-template-engine.git \
		-c pipeline.yml

	fly -t dev unpause-pipeline -p $(APP)
