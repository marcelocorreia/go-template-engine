#!/usr/bin/env bash
set -e

APP_NAME=$1
API_HTTPS=https://api.github.com/repos
GITHUB_USER=marcelocorreia
VERSION=$(curl ${API_HTTPS}/${GITHUB_USER}/${APP_NAME}/tags | jq '.[0].name' | sed 's/\"//g')
OS_DETECTED=$(uname -s)
TMPDIR=/tmp/homebrew-taps-update
PKG_FILE=${TMPDIR}/${APP_NAME}-darwin-amd64-${VERSION}.zip
if [ ${OS_DETECTED} == "Darwin" ];then
    SHA_CMD="shasum -a 256"
else
    SHA_CMD="sha256sum"
fi


git clone git@github.com:marcelocorreia/homebrew-taps.git ${TMPDIR}

curl https://github.com/marcelocorreia/${APP_NAME}/releases/download/${VERSION}/${APP_NAME}-darwin-amd64-${VERSION}.zip \
		-o ${PKG_FILE} -L

SUM=$(${SHA_CMD} ${PKG_FILE}  | awk {'print $1'})

echo ${SUM}

go-template-engine -s ${TMPDIR}/go-template-engine.tpl --var version=${VERSION} --var sha256sum=${SUM} \
		-o ${TMPDIR}/${APP_NAME}.rb

cd ${TMPDIR}
git add .
git commit -m "Updating ${APP_NAME} Release: ${VERSION}"
git push