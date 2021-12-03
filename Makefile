init:
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	go mod tidy

sqlcgen:
	sqlc generate