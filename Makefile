# Demands Terraform 0.13.x
TEST?=$$(go list ./... | grep -v 'vendor')
NAMESPACE=terraform.embracecloud.nl
COMPANY=embracecloud
PROVIDER=squidex
BINARY=terraform-provider-squidex
VERSION=0.4.2
OS=windows
ARCH=amd64

default: local-install

build:
	mkdir -p ./bin
	go build -o ./bin/${BINARY}_v${VERSION}

build-win:
	mkdir -p ./bin
	go build -o ./bin/${BINARY}_v${VERSION}.exe

install: build
	mkdir -p ~/.terraform.d/plugins/${NAMESPACE}/${COMPANY}/${PROVIDER}/${VERSION}/${OS}_${ARCH}
	mv ./bin/${BINARY}_v${VERSION} ~/.terraform.d/plugins/${NAMESPACE}/${COMPANY}/${PROVIDER}/${VERSION}/${OS}_${ARCH}

local-install: build
	mkdir -p ./examples/.terraform.d/plugins/${NAMESPACE}/${COMPANY}/${PROVIDER}/${VERSION}/${OS}_${ARCH}
	cp ./bin/${BINARY}_v${VERSION} ./examples/.terraform.d/plugins/${NAMESPACE}/${COMPANY}/${PROVIDER}/${VERSION}/${OS}_${ARCH}/
	mv ./bin/${BINARY}_v${VERSION} ./examples/.terraform/plugins/${NAMESPACE}/${COMPANY}/${PROVIDER}/${VERSION}/${OS}_${ARCH}/
	
local-install-win: build-win
	mkdir -p ./examples/terraform.d/plugins/${NAMESPACE}/${COMPANY}/${PROVIDER}/${VERSION}/${OS}_${ARCH}
	cp ./bin/${BINARY}_v${VERSION}.exe ./examples/terraform.d/plugins/${NAMESPACE}/${COMPANY}/${PROVIDER}/${VERSION}/${OS}_${ARCH}/
	mv ./bin/${BINARY}_v${VERSION}.exe ./examples/.terraform/plugins/${NAMESPACE}/${COMPANY}/${PROVIDER}/${VERSION}/${OS}_${ARCH}/ 
	
test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m