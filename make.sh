#!/usr/bin/env bash

dir=$(dirname $0)

make -C ${dir} ${MAKE_TARGETS}
