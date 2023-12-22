#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=kaidog.sield.gateway
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}