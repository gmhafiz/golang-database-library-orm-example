.PHONY : init sqlboiler run all

init:
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
	go get -d entgo.io/ent/cmd/ent
	go mod tidy

sqlboiler:
	sqlboiler --output db/sqlboiler/models psql

run:
	go mod tidy
	go vet ./...
	go fmt ./...
	go run main.go

all: init sqlboiler
	go mod tidy
	go generate ./...
	sqlc generate
	go vet ./...
	go fmt ./...
	go run main.go