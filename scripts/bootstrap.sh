#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=bluebird
ARGS="server"
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName} ${ARGS}