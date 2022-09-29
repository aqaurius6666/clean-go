gen-mock:
	@ which mockery > /dev/null || (echo \
	"mockery is not installed. Please install it first:\n \
go install github.com/vektra/mockery/v2@latest" \
	 && exit 1)
	
	@ mockery --all --outpkg mocks 

build:
	@ go build -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go

PROTO_DIR := proto
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto $(PROTO_DIR)/*/*.proto $(PROTO_DIR)/*/*/*.proto)
OUT_DIR := pkg/proto
# gen-proto:
# 	@ mkdir -p $(OUT_DIR)
# 	@ protoc \
# --gogo_opt paths=source_relative --gogo_out=plugins=grpc,\
# Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/types,\
# Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
# Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
# Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
# Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
# Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:\
# ${OUT_DIR} \
# -I=${PROTO_DIR} \
# ${PROTO_FILES}



dev-recreate:
	@ docker-compose -f deploy/dev/docker-compose.yaml up -d --force-recreate --build
# --govalidators_out=$OUT_DIR/authdto --govalidators_opt paths=source_relative \
protoc \
--openapiv2_out ./docs --openapiv2_opt logtostderr=true --openapiv2_opt allow_delete_body=true \
--validate_out="lang=go:$OUT_DIR" \
--gogo_opt paths=source_relative --gogo_out=plugins=grpc,\
Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgithub.com/mwitkow/go-proto-validators/validator.proto=github.com/mwitkow/go-proto-validators,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:\
$OUT_DIR/authdto \
-I=$PROTO_DIR -I=$PROTO_DIR/auth -I=$GOPATH/src -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway\
        api.proto

