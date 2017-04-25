#!/usr/bin/env bash

dir=$(dirname $0)

mkdir dist/

cd $dir

make $1

cp -Rv dist package
cp -Rv dist ../package
ls -l
pwd

