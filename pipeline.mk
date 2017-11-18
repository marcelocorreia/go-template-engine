include env.mk


fly_login:
	fly -t dev login -n dev -c https://ci.correia.io


git-push:
	git add . ; git commit -m "updating pipeline"; git push

pipeline-full: git-push pipeline


pipeline:
	fly -t dev set-pipeline \
		-n -p $(APP_NAME) \
		-c ci/pipeline.yml \
		-l $(HOME)/.ssh/ci-credentials.yml \
		-l ci/properties.yml

	fly -t dev unpause-pipeline -p $(APP_NAME)

.PHONY: pipeline

pipeline-destroy:
	fly -t dev destroy-pipeline -p $(APP_NAME)
.PHONY: pipeline-destroy