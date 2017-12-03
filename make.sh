#!/usr/bin/env bash

VERSION=dev
pwd

ls -lR

if [[ -f app-resource-version/version ]];then
    VERSION=$(cat app-resource-version/version)
fi

dir=$(dirname $0)

make -C ${dir} ${MAKE_TARGETS} VERSION=$VERSION
