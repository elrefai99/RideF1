.PHONY: api build migrate fmt vet test

api:
	cd app/api && go run ./cmd/backend

build:
	cd app/api && go build ./...

migrate:
	cd app/api && go run ./cmd/database

fmt:
	cd app/api && gofmt -w $$(go list -f '{{range .GoFiles}}{{$.Dir}}/{{.}} {{end}}' ./...)

vet:
	cd app/api && go vet ./...

test:
	cd app/api && go test ./...
