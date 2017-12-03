#!/usr/bin/env bash

dir=$(dirname $0)

VERSION=dev
HOMEBREW_REPO_PATH=${dir}/../homebrew-repo

if [[ -f app-resource-version/version ]];then
    VERSION=$(cat app-resource-version/version)
fi


make -C ${dir} ${MAKE_TARGETS} VERSION=$VERSION
