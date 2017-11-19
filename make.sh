#!/usr/bin/env bash

VERSION=$(cat resource-version/number)

VERSION=dev


if [[ -f resource-version/version ]];then
    VERSION=$(cat resource-version/version)
fi

echo VERSION=$VERSION

dir=$(dirname $0)

#make -C ${dir} ${MAKE_TARGETS} VERSION=$VERSION
