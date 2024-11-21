include scripts/auth_service.mk

.PHONY: gen-proto up create-migration

up:
	docker compose -f deployment/development/docker-compose.yaml down -v
	docker compose -f deployment/development/docker-compose.yaml up -d postgres
	docker compose -f deployment/development/docker-compose.yaml up -d --build --force-recreate auth_service

gen-proto:
	protoc -I pkg/grpc_stubs/auth_service \
		--go_out=pkg/grpc_stubs/auth_service \
		--go_opt=paths=source_relative \
		--go-grpc_out=pkg/grpc_stubs/auth_service \
		--go-grpc_opt=paths=source_relative \
		pkg/grpc_stubs/auth_service/auth_service.proto

create-migration:
	migrate create -ext sql -dir deployment/development/migrations -seq ${MIGRATION_NAME}

.DEFAULT_GOAL=up
