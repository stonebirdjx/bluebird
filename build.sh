#!/bin/bash
RUN_NAME=bluebird
DOCKER_NAME="1245863260/bluebird"
DOCKER_TAG=`date "+%s"`
mkdir -p output/bin
cp script/* output 2>/dev/null
chmod +x output/bootstrap.sh
swag init
go build -o output/bin/${RUN_NAME}

nerdctl build -t ${DOCKER_NAME}:${DOCKER_TAG} .

if [[ $? -ne 0 ]] ;then
    echo "${DOCKER_NAME}:${DOCKER_TAG} build failed"
    exit 0
fi

echo "${DOCKER_NAME}:${DOCKER_TAG} build success"