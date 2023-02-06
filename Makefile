include Makefile.env

EXEC_PATH := bin/taxonomy-cli
export ABS_EXEC_PATH := $(shell pwd)/$(EXEC_PATH)

.PHONY: local-build
local-build:
	CGO_ENABLED=0 GO111MODULE=on go build -o ${EXEC_PATH} main.go
	chmod +x $(EXEC_PATH)

.PHONY: local-clean
local-clean:
	rm -f ${EXEC_PATH}

DOCKER_HOSTNAME ?= ghcr.io
DOCKER_NAMESPACE ?= fybrik
DOCKER_NAME ?= taxonomy-cli
DOCKER_TAG ?= main
DOCKER_IMG := ${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/${DOCKER_NAME}:${DOCKER_TAG}

.PHONY: docker-build
docker-build:
	docker build . -t $(DOCKER_IMG)

.PHONY: docker-push
docker-push:
	docker push $(DOCKER_IMG)

.PHONY: test
test: local-build
	go test -v ./...
	rm -f ./test/testdata/output_*
