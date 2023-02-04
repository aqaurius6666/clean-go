ifeq ($(APP_NAME),)
	APP_NAME := clean-go
endif
ifeq ($(DEBUG),)
	DEBUG := false
endif
export DEBUG
export DEBUG_PORT=2345

build_tags := jsoniter

up: dev-recreate logs

install: 
	@ go mod tidy
	@ go install github.com/google/wire/cmd/wire@latest
	@ go install github.com/bufbuild/buf/cmd/buf@latest
	@ go install github.com/envoyproxy/protoc-gen-validate/cmd/protoc-gen-validate-go@latest
	@ go install github.com/cosmos/gogoproto/protoc-gen-gogo@latest
	@ go install github.com/cosmos/gogoproto/protoc-gen-gogofast@latest

dev-recreate:
	@ docker-compose -f deploy/dev/docker-compose.yaml up -d
	@ docker-compose -f deploy/dev/docker-compose.yaml up -d --force-recreate --build $(APP_NAME)

gen-mock:
	@ which mockery > /dev/null || (echo \
	"mockery is not installed. Please install it first:\n \
go install github.com/vektra/mockery/v2@latest" \
	 && exit 1)
	
	@ mockery --all --outpkg mocks 

build:
ifeq ($(DEBUG),true)
	@ go build -tags=${build_tags} -gcflags "all=-N -l" -buildvcs=false -o /tmp/$(APP_NAME) ./cmd/$(APP_NAME)
else
	@ go build -tags=${build_tags} -buildvcs=false -o /tmp/$(APP_NAME) ./cmd/$(APP_NAME)
endif

serve: build
ifeq ($(DEBUG),true)
	@ dlv --listen=0.0.0.0:$$DEBUG_PORT --accept-multiclient --headless=true --api-version=2 exec /tmp/$(APP_NAME) -- serve
else
	@ /tmp/$(APP_NAME) serve
endif

down:
	@ docker-compose -f deploy/dev/docker-compose.yaml stop
	@ docker-compose -f deploy/dev/monitor.docker-compose.yaml stop

gen-proto:
	@ which buf > /dev/null || (echo \
	"buf is not installed. Please install it first:\n \
go install github.com/bufbuild/buf/cmd/buf@latest" \
	 && exit 1)
	
	@ buf generate --template buf.gen.yaml
	@ go generate ./pkg/swagger

wire:
	@ wire ./internal/...
	@ wire ./cmd/...

run:
	@ export $(grep -v '^#' deploy/dev/.env | xargs) && go run ./... serve


kafka:
	@ docker-compose -f deploy/dev/kafka.docker-compose.yaml up -d

monitor:
	@ docker-compose -f deploy/dev/monitor.docker-compose.yaml up -d

logs:
	@ docker-compose -f deploy/dev/docker-compose.yaml logs -f $(APP_NAME) 