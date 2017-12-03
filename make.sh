#!/usr/bin/env bash

dir=$(dirname $0)

VERSION=dev
HOMEBREW_REPO_PATH=$(pwd)/${dir}/../homebrew-repo

if [[ -f app-resource-version/version ]];then
    VERSION=$(cat app-resource-version/version)
fi

echo VERSION=${VERSION}}
echo HOMEBREW_REPO_PATH=${HOMEBREW_REPO_PATH}


make -C ${dir} ${MAKE_TARGETS} VERSION=$VERSION
