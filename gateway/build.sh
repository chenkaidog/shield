#!/bin/bash
RUN_NAME=kaidog.sield.gateway
mkdir -p output/bin output/conf output/log
cp script/* output/
cp conf/* output/conf/
chmod +x output/bootstrap.sh
go build -o output/bin/${RUN_NAME}