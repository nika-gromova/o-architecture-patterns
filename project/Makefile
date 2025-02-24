CURDIR=$(shell pwd)
LOCAL_BIN:=$(CURDIR)/bin
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

get-bin:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

generate-rules:
	mkdir -p pkg/rules_v1
	$(PROTOC) -I api/rules/v1 -I vendor-proto \
	--go_out pkg/rules_v1 --go_opt paths=source_relative \
	--go-grpc_out pkg/rules_v1 --go-grpc_opt paths=source_relative \
	--grpc-gateway_out pkg/rules_v1 --grpc-gateway_opt paths=source_relative \
	--validate_out="lang=go,paths=source_relative:pkg/rules_v1" \
	--openapiv2_out pkg/rules_v1 \
	api/rules/v1/api.proto

generate: get-bin generate-rules
