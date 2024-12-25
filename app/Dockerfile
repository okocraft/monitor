# syntax=docker/dockerfile:1.12.1@sha256:93bfd3b68c109427185cd78b4779fc82b484b0b7618e36d0f104d4d801e66d25
FROM golang:1.23.4@sha256:cb00c135e0e60590614a68f1361fcafdf0f4d2bec2efa7d11b4eb634d1cbe4ed AS base

WORKDIR /tmp/app

COPY go.mod go.sum  ./

RUN go mod download

FROM base AS build-http-server

COPY . ./

RUN go build -o /bin/http-server cmd/http/main.go

FROM build-http-server AS http-server

COPY --from=build-http-server /bin/http-server /server

ENTRYPOINT [ "/server" ]
