#!/usr/bin/env bash

set -eo pipefail

echo "Generating gogo proto code"
cd proto

buf generate --template buf.gen.gogo.yaml $file

cd ..

# after the proto files have been generated add them to the the repo
# in the proper location. Then, remove the ephemeral tree used for generation
cp -r github.com/umee-network/umee/v4/* .
rm -rf github.com


