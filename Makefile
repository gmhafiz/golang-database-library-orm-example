init:
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres
	go mod tidy

sqlcgen:
	sqlc generate