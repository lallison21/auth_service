.PHONY: gen_auth_service

gen_auth_service:
	protoc -I pkg/protos/auth_service \
		--go_out=pkg/grpc_stubs/auth_service \
		--go_opt=paths=source_relative \
		--go-grpc_out=pkg/grpc_stubs/auth_service \
		--go-grpc_opt=paths=source_relative \
		pkg/protos/auth_service/auth_service.proto