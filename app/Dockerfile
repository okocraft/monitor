# syntax=docker/dockerfile:1.23.0@sha256:2780b5c3bab67f1f76c781860de469442999ed1a0d7992a5efdf2cffc0e3d769
FROM golang:1.26.2@sha256:5f3787b7f902c07c7ec4f3aa91a301a3eda8133aa32661a3b3a3a86ab3a68a36 AS base

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
