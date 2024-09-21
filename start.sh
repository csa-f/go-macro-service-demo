#!/bin/bash
set -x
ROOT=$(cd `dirname $0`; pwd)
cd $ROOT

if [ "$BINARY" = "" ];then
    BINARY="test-api"
fi

# while [ true ];do
#     $ROOT/bin/test-api $@ >> log/run.log 2>&1
#     sleep 10
# done

$ROOT/bin/test-api $@ >> log/run.log 2>&1
