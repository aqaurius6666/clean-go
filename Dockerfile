
FROM golang:1.19-alpine as base

# ARG GRPC_HEALTH_PROBE_VERSION=v0.4.5

RUN apk add --no-cache ca-certificates git curl build-base

# RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
#     chmod +x /bin/grpc_health_probe 

WORKDIR /app

FROM base as dev

ENV GO111MODULE=on

# ARG GRPC_CLIENT_CLI_VERSION=v1.10.0

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && mv ./bin/air /bin/air

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