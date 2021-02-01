#!/bin/bash

set -euo pipefail

echo "Step 1. Changing to project root directory $(dirname "$0")/../../"
cd "$(dirname "$0")/../../"
echo "Done."

echo "Step 2. Creating artifacts dir."
mkdir -p artifacts
echo "Done."

echo "Step 3. Building provider for linux..."
GOOS=linux GOARCH=amd64 go build -o ./artifacts/linux/terraform-provider-squidex_v${VERSION}
echo "Done."

echo "Step 3. Building provider for windows..."
GOOS=windows GOARCH=amd64 go build -o ./artifacts/windows/terraform-provider-squidex_v${VERSION}.exe
echo "Done."

echo "Step 4. Changing current dir to $(pwd)/artifacts"
cd ./artifacts
echo "Done."

echo "Step 5. Zipping provider..."
cd linux
zip ../terraform-provider-squidex_v${VERSION}_linux_amd64.zip ./terraform-provider-squidex_v${VERSION}
cd ../windows
zip ../terraform-provider-squidex_v${VERSION}_windows_amd64.zip ./terraform-provider-squidex_v${VERSION}.exe
cd ..
rm -r linux
rm -r windows
sha256sum -b * >SHA256SUMS
echo "Done."
