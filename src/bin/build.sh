#!/bin/bash
set -x
ROOT=$(cd `dirname $0`; pwd)

if [ "$CMD" = "" ];then
    CMD="gateway"
fi

# while [ true ];do
#     $ROOT/bin/test-api $@ >> log/run.log 2>&1
#     sleep 10
# done
cd ../$CMD/
go mod tidy
cd $ROOT

go build ../$CMD/cmd/$CMD.go
cp ../*/conf/*.yml ./conf