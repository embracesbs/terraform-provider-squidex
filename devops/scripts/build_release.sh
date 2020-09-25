#!/bin/bash

set -euo pipefail

echo "Step 1. Changing to project root directory $(dirname "$0")/../../"
cd "$(dirname "$0")/../../"
echo "Done."

echo "Step 2. Creating artifacts dir."
mkdir -p artifacts
echo "Done."

echo "Step 3. Building provider for linux..."
GOOS=linux GOARCH=amd64 go build -o ./artifacts/terraform-provider-squidex_${VERSION}_linux_amd64
echo "Done."

echo "Step 4. Changing current dir to $(pwd)/artifacts"
cd ./artifacts
echo "Done."

echo "Step 5. Zipping provider..."
zip terraform-provider-squidex_${VERSION}_linux_amd64.zip ./terraform-provider-squidex_${VERSION}_linux_amd64
rm ./terraform-provider-squidex_${VERSION}_linux_amd64
sha256sum -b * > SHA256SUMS
echo "Done."
