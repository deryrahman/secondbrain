.PHONY: all build

build:
	rm -rf build
	mkdir build
	go build -o ./build/secondbrain ./cmd

generate-handler:
	rm -r ./codegen/handler/*
	oapi-codegen -package handler -generate types -o ./codegen/handler/models.go openapi.yaml
	oapi-codegen -package handler -generate chi-server -o ./codegen/handler/chi_server.go openapi.yaml

generate-storage:
	rm -r ./codegen/storage/*
	sqlc generate
	