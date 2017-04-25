#!/usr/bin/env bash

dir=$(dirname $0)

mkdir dist/

cd $dir

make $1

ls -l
pwd

