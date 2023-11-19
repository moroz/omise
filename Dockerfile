ARG GO_VERSION=1.21.4
ARG DEBIAN_VERSION=bookworm
ARG NODE_VERSION=20.9.0

FROM golang:${GO_VERSION}-${DEBIAN_VERSION} AS builder

WORKDIR /app

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN templ generate

RUN go build -o server

FROM node:$NODE_VERSION AS assets

WORKDIR /app

RUN npm i -g pnpm

COPY assets/package.json assets/pnpm-lock.yaml ./

RUN pnpm install --frozen-lockfile

FROM debian:${DEBIAN_VERSION}-slim AS runner

COPY --from=builder /app/server /app/server

CMD /app/server
