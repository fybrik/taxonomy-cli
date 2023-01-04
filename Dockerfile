# Copyright 2022 IBM Corp.
# SPDX-License-Identifier: Apache-2.0
FROM docker.io/golang:1.19.4-alpine3.17
COPY . /src/
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o ../taxonomy-cli main.go
WORKDIR /
RUN rm -r /src
ENTRYPOINT ["/taxonomy-cli"]
