.PHONY: all build

build:
	goreleaser build --single-target --clean --snapshot

lint:
	golangci-lint run

generate-handler:
	rm -r ./codegen/handler/*
	oapi-codegen -package handler -generate types -o ./codegen/handler/models.go openapi.yaml
	oapi-codegen -package handler -generate chi-server -o ./codegen/handler/chi_server.go openapi.yaml

generate-storage:
	rm -r ./codegen/storage/*
	sqlc generate

db-migrate-up:
	sql-migrate up --env development
	
dev-tools:
	go install github.com/rubenv/sql-migrate/...@v1.5.1
	go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.18.0
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.13.0
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.53.3
	go install github.com/goreleaser/goreleaser@v1.19.2
