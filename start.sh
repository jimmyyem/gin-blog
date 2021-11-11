#!/bin/bash

WORK_DIR=$(cd `dirname $0`; pwd)
GO=/usr/local/go/bin/go
echo $WORK_DIR
source /etc/profile
cd $WORK_DIR/pkg/e && $GO generate && cd ../../ && $GO build -o main.sh && ./main.sh