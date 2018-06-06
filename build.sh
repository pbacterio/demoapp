#!/bin/bash

if [[ $1 == "clean" ]]
then
    set -ex
    rm -r build
    exit
fi

if [[ $1 == "docker" ]]
then
    set -ex
    docker build --force-rm --tag demoapp .
    exit
fi

set -ex
dep ensure
CGO_ENABLED=0 GOOS=linux go build -o build/linux/demoapp -ldflags="-s -w"
upx -q build/linux/demoapp