APP=sample
pipeline:
	fly -t dev set-pipeline \
		-n -p $(APP) \
		-c pipeline.yml

	fly -t dev unpause-pipeline -p $(APP)

	fly -t dev trigger-job -j $(APP)/job-pass-files