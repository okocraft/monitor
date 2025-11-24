# syntax=docker/dockerfile:1.20.0@sha256:26147acbda4f14c5add9946e2fd2ed543fc402884fd75146bd342a7f6271dc1d
FROM golang:1.25.4@sha256:698183780de28062f4ef46f82a79ec0ae69d2d22f7b160cf69f71ea8d98bf25d AS base

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
