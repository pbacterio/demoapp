#!/bin/bash

function build_image {
    docker build --force-rm --build-arg OS=${1} --build-arg ARCH=${2} -t demoapp:${1}-${2} .
}

pushd "$(dirname "$0")"

if [[ $1 == "clean" ]]
then
    set -ex
    rm -rf build
    exit
fi

if [[ $1 == "docker" ]]
then
    set -ex
    build_image linux amd64
    build_image linux arm
    build_image linux arm64
    build_image darwin amd64
    exit
fi

# function in-docker {
#     docker run -v $PWD:/go/src/github.com/pbacterio/demoapp \
#       --rm -it gobuilder "/go/src/github.com/pbacterio/demoapp/build.sh"
# }

function build {
    CGO_ENABLED=0 GOOS=$1 GOARCH=$2 go build -o build/${1}_${2}/demoapp -ldflags="-s -w"
}

set -ex
dep ensure
build linux amd64
build linux arm
build linux arm64
build darwin amd64
upx -q build/*/*
