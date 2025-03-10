CURDIR=$(shell pwd)
LOCAL_BIN:=$(CURDIR)/bin
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

get-bin:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

# Устанавливаем proto описания google/googleapis
vendor-proto/google/api:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/googleapis/googleapis vendor-proto/googleapis &&\
 	cd vendor-proto/googleapis &&\
	git sparse-checkout set --no-cone google/api &&\
	git checkout
	mkdir -p  vendor-proto/google
	mv vendor-proto/googleapis/google/api vendor-proto/google
	rm -rf vendor-proto/googleapis

# Устанавливаем proto описания google/protobuf
vendor-proto/google/protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf vendor-proto/protobuf &&\
	cd vendor-proto/protobuf &&\
	git sparse-checkout set --no-cone src/google/protobuf &&\
	git checkout
	mkdir -p  vendor-proto/google
	mv vendor-proto/protobuf/src/google/protobuf vendor-proto/google
	rm -rf vendor-proto/protobuf

vendor-proto/options:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/grpc-ecosystem/grpc-gateway vendor-proto/options &&\
	cd vendor-proto/options &&\
	git sparse-checkout set --no-cone protoc-gen-openapiv2/options &&\
	git checkout
	mkdir -p  vendor-proto/protoc-gen-openapiv2
	mv vendor-proto/options/protoc-gen-openapiv2/options vendor-proto/protoc-gen-openapiv2
	rm -rf vendor-proto/options

generate-bin: get-bin vendor-proto/google/api vendor-proto/google/protobuf vendor-proto/options
generate-rules:
	mkdir -p pkg/rules_v1
	$(PROTOC) -I api/rules/v1 -I vendor-proto \
	--go_out pkg/rules_v1 --go_opt paths=source_relative \
	--go-grpc_out pkg/rules_v1 --go-grpc_opt paths=source_relative \
	--grpc-gateway_out pkg/rules_v1 --grpc-gateway_opt paths=source_relative \
	--validate_out="lang=go,paths=source_relative:pkg/rules_v1" \
	--openapiv2_out pkg/rules_v1 \
	api/rules/v1/api.proto

generate: generate-rules

build:
	go build -o bin/ cmd/main.go
