#!/usr/bin/env bash

VERSION=dev

if [[ -f resource_version/version ]];then
    VERSION=$(cat resource_version/version)
fi

echo VERSION=$VERSION

dir=$(dirname $0)

make -C ${dir} ${MAKE_TARGETS} VERSION=$VERSION
