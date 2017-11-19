#!/usr/bin/env bash

VERSION=$(cat resource_version/number)

#VERSION=dev
#
#
#if [[ -f resource-version/version ]];then
#    VERSION=$(cat resource-version/version)
#    echo VERSION=$VERSION
#fi

dir=$(dirname $0)

#make -C ${dir} ${MAKE_TARGETS} VERSION=$VERSION
