#!/usr/bin/env bash

set -e

if [[ "$(cat ${HOME}/.bash_profile | grep "_go-template-engine")" = "" ]];then
	echo "alias go-template-engine=\"docker run --rm -it -v $(pwd):/app -w /app marcelocorreia/go-template-engine"  >> ${HOME}/.bash_profile
else
	sed -i .bk 's/alias go-template-engine.*/alias go-template-engine=\"docker run --rm -it -v $\(pwd\):\/app -w \/app marcelocorreia\/go-template-engine\"/' ${HOME}/.bash_profile
fi

echo "Alias go-template-engine created."
echo "It will be available on the next shell session"
echo "To force the load of the new alias you can try"
echo "$> source ${HOME}/.bash_profile"