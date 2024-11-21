include scripts/auth_service.mk

.PHONY: gen-proto

gen-proto:
	protoc -I pkg/grpc_stubs/auth_service \
		--go_out=pkg/grpc_stubs/auth_service \
		--go_opt=paths=source_relative \
		--go-grpc_out=pkg/grpc_stubs/auth_service \
		--go-grpc_opt=paths=source_relative \
		pkg/grpc_stubs/auth_service/auth_service.proto

.DEFAULT_GOAL: build
