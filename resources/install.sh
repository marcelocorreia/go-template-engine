#!/usr/bin/env bash

if [[ "$(cat ${HOME}/.bash_profile | grep "_go-template-engine")" = "" ]];then
	echo "alias _go-template-engine=\"docker run --rm marcelocorreia/go-template-engine"\"  >> ${HOME}/.bash_profile
else
	sed -i .bk 's/alias _go-template-engine.*/alias _go-template-engine=\"docker run --rm -it -v $\(pwd\):\/app -w \/app marcelocorreia\/go-template-engine\"/' ${HOME}/.bash_profile
fi
