#!/bin/bash

set -euo pipefail

echo "Step 1. Changing to project root directory $(dirname "$0")/../../"
cd "$(dirname "$0")/../../"
echo "Done."

echo "Step 2. Creating artifacts dir."
mkdir -p artifacts
echo "Done."

echo "Step 3. Building provider for linux..."
GOOS=linux GOARCH=amd64 go build -o ./linux/terraform-provider-squidex_v${VERSION}
echo "Done."

echo "Step 4. Building provider for windows..."
GOOS=windows GOARCH=amd64 go build -o ./windows/terraform-provider-squidex_v${VERSION}.exe
echo "Done."

echo "Step 5. Zipping provider..."
zip ./artifacts/terraform-provider-squidex_v${VERSION}_linux_amd64.zip ./linux/terraform-provider-squidex_v${VERSION}
zip ./artifacts/terraform-provider-squidex_v${VERSION}_windows_amd64.zip ./windows/terraform-provider-squidex_v${VERSION}.exe
sha256sum -b * >SHA256SUMS
echo "Done."
