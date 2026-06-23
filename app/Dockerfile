# syntax=docker/dockerfile:1.24.0@sha256:87999aa3d42bdc6bea60565083ee17e86d1f3339802f543c0d03998580f9cb89
FROM golang:1.26.4@sha256:792443b89f65105abba56b9bd5e97f680a80074ac62fc844a584212f8c8102c3 AS base

WORKDIR /tmp/app

COPY go.mod go.sum ./

RUN go mod download

FROM base AS build-http-server

COPY . ./

RUN go build -o /bin/http-server cmd/http/main.go

FROM base AS build-cleanup-batch

COPY . ./

RUN go build -o /bin/cleanup-batch cmd/batch/cleanup/main.go

FROM base AS build-setup-cmd

COPY . ./

RUN go build -o /bin/setup cmd/setup/main.go

FROM build-http-server AS http-server

COPY --from=build-http-server /bin/http-server /server

ENTRYPOINT [ "/server" ]

FROM build-cleanup-batch AS cleanup-batch

COPY --from=build-cleanup-batch /bin/cleanup-batch /cleanup

ENTRYPOINT [ "/cleanup" ]

FROM build-setup-cmd AS setup-cmd

COPY --from=build-setup-cmd /bin/setup /setup

ENTRYPOINT [ "/setup" ]
