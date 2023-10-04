
FROM golang:1.19-alpine as base

RUN apk add --no-cache ca-certificates git curl build-base

WORKDIR /app

FROM base as dev

ENV GO111MODULE=on

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go install github.com/cosmtrek/air@latest

# RUN go install github.com/vadimi/grpc-client-cli/cmd/grpc-client-cli@${GRPC_CLIENT_CLI_VERSION}

FROM base as builder

ARG BINARY_NAME=clean-go

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build
RUN /tmp/${BINARY_NAME} --help


FROM alpine as release

ARG BINARY_NAME=clean-go

WORKDIR /app
# COPY --from=builder /bin/grpc_health_probe /bin/grpc_health_probe
COPY --from=builder /tmp/${BINARY_NAME} /usr/bin/app
ENTRYPOINT ["/usr/bin/app"]