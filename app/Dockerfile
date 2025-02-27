# syntax=docker/dockerfile:1.14.0@sha256:0232be24407cc42c983b9b269b1534a3b98eea312aad9464dd0f1a9e547e15a7
FROM golang:1.24.0@sha256:3f7444391c51a11a039bf0359ee81cc64e663c17d787ad0e637a4de1a3f62a71 AS base

WORKDIR /tmp/app

COPY go.mod go.sum  ./

RUN go mod download

FROM base AS build-http-server

COPY . ./

RUN go build -o /bin/http-server cmd/http/main.go

FROM base AS build-cleanup-batch

COPY . ./

RUN go build -o /bin/cleanup-batch cmd/batch/cleanup/main.go

FROM build-http-server AS http-server

COPY --from=build-http-server /bin/http-server /server

ENTRYPOINT [ "/server" ]

FROM build-cleanup-batch AS cleanup-batch

COPY --from=build-cleanup-batch /bin/cleanup-batch /cleanup

ENTRYPOINT [ "/cleanup" ]
