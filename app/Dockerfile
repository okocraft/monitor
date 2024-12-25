# syntax=docker/dockerfile:1.12.1@sha256:93bfd3b68c109427185cd78b4779fc82b484b0b7618e36d0f104d4d801e66d25
FROM golang:1.23.4@sha256:f06d2bb355a67ccc6c23f3699766323a09ed0a4b724a6b25a300d4b30e01f02c AS base

WORKDIR /tmp/app

COPY go.mod go.sum  ./

RUN go mod download

FROM base AS build-http-server

COPY . ./

RUN go build -o /bin/http-server cmd/http/main.go

FROM build-http-server AS http-server

COPY --from=build-http-server /bin/http-server /server

ENTRYPOINT [ "/server" ]
