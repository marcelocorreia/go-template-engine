include variables.tfvars

TF_IMAGE?=marcelocorreia/terraform
VARS_FILE?=variables.tfvars

# TF
tf-plan: tf-version tf-init
	$(call terraform, plan -var-file $(VARS_FILE) -input=false)
.PHONY: tf-plan


tf-state-file-ckeck:
	@[ -f state.tf ] && echo "terraform state file found, skipping creation" || $(call create_state)

tf-get:
	$(call terraform, get)
.PHONY: tf-get

tf-get-update:
	$(call terraform, -update=true)
.PHONY: tf-get-update

tf-apply: tf-plan
	$(call terraform, apply -var-file $(VARS_FILE) -input=false)
.PHONY: tf-apply

tf-destroy:
	$(call terraform, destroy -force -var-file $(VARS_FILE))
.PHONY: tf-destroy

tf-deploy: tf-apply tf-state-list
.PHONY: tf-deploy

tf-refresh:
	$(call terraform, refresh -var-file $(VARS_FILE))
.PHONY: tf-refresh

tf-output: tf-init
	$(call terraform, output)
.PHONY: tf-output

tf-version:
	$(call terraform, version)

tf-show:
	$(call terraform, show)

tf-shell:
	@docker run --rm -it -v $(shell pwd):/opt/workspace \
			-v $(shell pwd):/opt/workspace  \
    		-v $(HOME)/.aws:/root/.aws \
    		-v $(HOME)/.ssh:/root/.ssh \
			-e AWS_ACCESS_KEY_ID=${aws_access_key_id} \
    		-e AWS_SECRET_ACCESS_KEY=${aws_secret_access_key} \
    		-e AWS_DEFAULT_REGION="ap-southeast-2" \
    		$(TF_IMAGE)\
    		bash
.PHONY: tf-shell

tf-state-list: tf-init
	$(call terraform, state list)
.PHONY: tf-state-list

tf-state-pull: tf-init
	$(call terraform, state pull)
.PHONY: tf-state-list

tf-init:
	$(call terraform, init)
.PHONY: tf-init



# DOCKER
tf-image-update:
	docker pull $(TF_IMAGE)
.PHONY: tf-image-update

tf-workspace-list:
	$(call terraform, workspace list)


# ROUTINES
define terraform
	terraform $1
endef

define create_state
	echo "Creating Terraform state file"
	echo "terraform {" > state.tf
    echo "  backend \"s3\"" {     >> state.tf
    echo "    bucket = \"$(state_bucket)\""      >> state.tf
    echo "    key    = \"$(account_id)/terraform.tfstate\"" >> state.tf
    echo "    region = \"$(state_bucket_region)\"" >> state.tf
    echo "    workspace_key_prefix = \"$(account_id)\"" >> state.tf
    echo "  }" >> state.tf
    echo "}" >> state.tf
endef

show-vars:
	cat *tf | json2hcl -reverse | jq '.variable?'
#	cat *.tf  | grep variable | grep -v\
#	perl -p -e 's/\n//g'| \
#	perl -p -e 's/variable/\nvariable/g' | \
#	perl -p -e 's/\{.*\}//g' | \
#	perl -p -e 's/variable\s+\"//g' | \
#	perl -p -e 's/\"(.*)/\=\"\"/g'

whaaat:
	@cat *.tf | \
		perl -p -e 's/\n//g'| \
	  	perl -p -e 's/variable/\nvariable/g' | \
		perl -p -e 's/\{.*\}//g' | \
		perl -p -e 's/variable\s+\"//g' | \
		perl -p -e 's/\"/\=\"\"/g'


create-team-github:
	fly -t main set-team -n $(CI_TEAM) \
        --github-auth-client-id $(GITHUB_CLIENT_ID) \
        --github-auth-client-secret $(GITHUB_SECRET) \
        --github-auth-team $(GITHUB_TEAM)

create-team-password:
	$(call create_team $(CI_TEAM) \
        --github-auth-client-id $(GITHUB_CLIENT_ID) \
        --github-auth-client-secret $(GITHUB_SECRET) \
        --github-auth-team $(GITHUB_TEAM))

define create_team
	echo fly -t main set-team -n $(CI_TEAM)
endef