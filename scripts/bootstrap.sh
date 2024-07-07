#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=bluebird
ARGS="server -c config/server.yaml"
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName} ${ARGS}