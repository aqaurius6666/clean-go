APP_NAME := clean-go

dev-recreate: observability
	@ export DEBUG=false && \
	export DEBUG_PORT=2345 && \
	docker-compose --project-name=${APP_NAME} -f deploy/dev/docker-compose.yaml up -d --force-recreate --build

gen-mock:
	@ which mockery > /dev/null || (echo \
	"mockery is not installed. Please install it first:\n \
go install github.com/vektra/mockery/v2@latest" \
	 && exit 1)
	
	@ mockery --all --outpkg mocks 

build:
	@ if [ -z "$(DEBUG)" ]; then \
		go build -buildvcs=false -o /tmp/$(APP_NAME) ./cmd/$(APP_NAME); \
    else \
       	go build -gcflags "all=-N -l" -buildvcs=false -o /tmp/$(APP_NAME) ./cmd/$(APP_NAME); \
    fi

serve: build
	@ if [ -z "$(DEBUG)" ]; then \
		/tmp/$(APP_NAME) serve; \
    else \
       	dlv --listen=0.0.0.0:$$DEBUG_PORT --accept-multiclient --headless=true --api-version=2 exec /tmp/$(APP_NAME) -- serve; \
    fi

gen-proto:
	@ which buf > /dev/null || (echo \
	"buf is not installed. Please install it first:\n \
go install github.com/bufbuild/buf/cmd/buf@latest" \
	 && exit 1)
	
	@ buf generate --template buf.gen.yaml

wire:
	@ wire internal/...

run:
	@ export $(cat deploy/dev/.env | xargs) && go run ./... serve

observability:
	@ docker-compose --project-name=${APP_NAME} -f deploy/dev/observability/docker-compose.yaml up -d --force-recreate

dev-recreate-debug: 
	@ export DEBUG=true && \
	export DEBUG_PORT=2345 && \
	docker-compose -f deploy/dev/docker-compose.yaml up -d --force-recreate --build

kafka:
	@ docker-compose -f deploy/dev/kafka.docker-compose.yaml up -d

logs:
	@ export DEBUG=false && \
	export DEBUG_PORT=2345 && \
	docker-compose -f deploy/dev/docker-compose.yaml logs -f cleango 