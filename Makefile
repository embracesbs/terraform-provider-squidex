TEST?=$$(go list ./... | grep -v 'vendor')
NAME=squidex
BINARY=terraform-provider-${NAME}
VERSION=0.2.0
OS=linux
ARCH=amd64

default: install

build:
	go build -o ${BINARY}_v${VERSION}-${OS}-${ARCH}

install: build
	mkdir -p ~/.terraform.d/plugins/${OS}_${ARCH}
	mv ${BINARY}_v${VERSION}-${OS}-${ARCH} ~/.terraform.d/plugins/${OS}_${ARCH}

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m