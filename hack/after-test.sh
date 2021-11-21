#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

PLATFORM=linux/amd64
OUTPUT=./output

mkdir -p $OUTPUT/$PLATFORM
for service_name in `ls $(pwd)/cmd`; do
    if [ -z `pidof $service_name` ];then
        echo "[WARNING] $service_name process not found"
    else
        kill -9 `pidof $service_name`
    fi
done
