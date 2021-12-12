package ent

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"godb/config"
	"godb/db/ent/ent/gen"
)

type database struct {
	db *gen.Client
}

func New(cfg config.Database) *gen.Client {
	db, err := sql.Open("pgx", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, "ent", cfg.Password, cfg.SSLMode))
	if err != nil {
		log.Fatal(err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)

	client := gen.NewClient(gen.Driver(drv))

	f, err := os.Create("./db/ent/migrate.sql")
	if err != nil {
		log.Fatalf("create migrate file: %v", err)
	}
	defer f.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err := client.Schema.WriteTo(context.Background(), f); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}

	return client
}

func NewRepo(db *gen.Client) *database {
	return &database{
		db: db,
	}
}
