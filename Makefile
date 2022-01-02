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

check: sqlboiler
	go generate ./...
	sqlc generate
	go mod tidy
	go vet ./...
	go fmt ./...

run: check
	go run main.go

all: init check sqlboiler run
