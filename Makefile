include Makefile.env

EXEC_NAME = taxonomy-cli

.PHONY: build-local
build-local:
	CGO_ENABLED=0 GO111MODULE=on go build -o ${EXEC_NAME} main.go

.PHONY: clean-local
clean-local:
	rm -f ${EXEC_NAME}

DOCKER_HOSTNAME ?= ghcr.io
DOCKER_NAMESPACE ?= fybrik
DOCKER_NAME ?= taxonomy-cli
DOCKER_TAG ?= main
DOCKER_IMG := ${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/${DOCKER_NAME}:${DOCKER_TAG}

.PHONY: build-docker
build-docker:
	docker build . -t $(DOCKER_IMG)

.PHONY: push-docker
push-docker:
	docker push $(DOCKER_IMG)
