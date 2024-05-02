#!/bin/bash

BIN_NAME="bluebird"

if [ !-d "./bin" ]; then
   mkdir bin
fi

go build -o ./bin/${BIN_NAME}

if [ $? -eq 0 ]; then
   echo "Build success"
else
   echo "Build failed"
fi

