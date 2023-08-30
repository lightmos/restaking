#!/usr/bin/env bash

set -eo pipefail

echo "Generating gogo proto code"
cd proto

buf generate --template buf.gen.gogo.yaml $file

cd ..

# move proto files to the right places
cp -r github.com/lightmos/share/* ./
rm -rf github.com

go env -w GOPROXY="https://goproxy.cn"

go mod tidy

./scripts/protocgen-pulsar.sh