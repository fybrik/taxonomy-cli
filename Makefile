include Makefile.env

EXEC_NAME = taxonomy-cli

.PHONY: local-build
local-build:
	CGO_ENABLED=0 GO111MODULE=on go build -o ${EXEC_NAME} main.go

.PHONY: local-clean
local-clean:
	rm -f ${EXEC_NAME}

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
