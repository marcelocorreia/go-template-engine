TARDIS_HOME?=$(shell pwd)
#
ORGANISATION?=correia
#
ANSIBLE_HOME?=$(TARDIS_HOME)/ansible

APP=tardis-dna
DIST_DIR?=$(TARDIS_HOME)/dist

ENVS_DIR?=/Users/marcelo/environments
ENVS_REPO_URL?=git@github.com:marcelocorreia/environments.git

CI_TARGET=dev
CI_TEAM=dev

DOCKER_CONTAINER=$(APP)
DOCKER_HOME=$(TARDIS_HOME)/docker
DOCKER_NAMESPACE=290901122349.dkr.ecr.ap-southeast-2.amazonaws.com/correia.io

GIT_REPO=$(APP)
GITHUB_NAMESPACE?=marcelocorreia

MODULES_REPO_URL=git@github.com:marcelocorreia/terraform-modules.git
#PIPELINE_NAME=$(APP)
STATE_BUCKET=correia-io
STATE_BUCKET_REGION=ap-southeast-2

TEMPLATES_DIR=$(TARDIS_HOME)/terraform/templates
TEST_OUTPUT_DIR=$(TARDIS_HOME)/tmp

TMP_BUILD=$(TARDIS_HOME)/tmp-build
