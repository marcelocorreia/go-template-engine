#!/usr/bin/env bash

dir=$(dirname $0)
cd $dir
echo make $1
pwd
ls -l
make $1
