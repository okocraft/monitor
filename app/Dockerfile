# syntax=docker/dockerfile:1.25.0@sha256:0adf442eae370b6087e08edc7c50b552d80ddf261576f4ebd6421006b2461f12
FROM golang:1.26.4@sha256:f96cc555eb8db430159a3aa6797cd5bae561945b7b0fe7d0e284c63a3b291609 AS base

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
