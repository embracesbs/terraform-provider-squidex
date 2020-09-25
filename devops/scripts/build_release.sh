#!/bin/bash

export BINARY=terraform-provider-squidex

cd `dirname "$0"`/../../

echo "Create artifacts dir."
mkdir -p artifacts
echo "Done."
echo "Building provider for linux..."
GOOS=linux GOARCH=amd64 go build -o ./artifacts/${BINARY}_${VERSION}_linux_amd64
echo "Done."

echo "Zipping provider..."

cd artifacts

zip ${BINARY}_${VERSION}_linux_amd64.zip ./${BINARY}_${VERSION}_linux_amd64
rm ./${BINARY}_${VERSION}_linux_amd64
sha256sum -b * > SHA256SUMS

echo "Done."
