#!/bin/bash

WORK_DIR=$(cd `dirname $0`; pwd)
echo $WORK_DIR
GOMOD=$WORK_DIR/go.mod
export GO111MODULE=auto
export GOPATH=/home/ubuntu/go
export GOROOT=/usr/local/go
cd $WORK_DIR && go generate && go build -o main.sh && ./main.sh