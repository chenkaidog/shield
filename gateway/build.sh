#!/bin/bash
RUN_NAME=kaidog.sield.gateway
mkdir -p output/bin output/conf output/log
mkdir -p output/render/html output/render/css output/render/js

cp script/* output/
cp conf/* output/conf/
cp -r render/* output/render/

chmod +x output/bootstrap.sh
go build -o output/bin/${RUN_NAME}