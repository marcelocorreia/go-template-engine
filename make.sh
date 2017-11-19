#!/usr/bin/env bash

ls -l ./
ls -l ../

VERSION=dev


if [[ -f resource-version/version ]];then
    VERSION=$(cat resource-version/version)
fi

echo VERSION=$VERSION

dir=$(dirname $0)

#make -C ${dir} ${MAKE_TARGETS} VERSION=$VERSION
