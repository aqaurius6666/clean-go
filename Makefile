gen-mock:
	@ which mockery > /dev/null || (echo \
	"mockery is not installed. Please install it first:\n \
go install github.com/vektra/mockery/v2@latest" \
	 && exit 1)
	
	@ mockery --all --outpkg mocks 

build:
	@ go build -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go

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

dev-recreate:
	@ docker-compose -f deploy/dev/docker-compose.yaml up -d --force-recreate --build

kafka:
	@ docker-compose -f deploy/dev/kafka.docker-compose.yaml up -d