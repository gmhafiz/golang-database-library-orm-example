package ent

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"godb/config"
	"godb/db/ent/ent/gen"
	_ "godb/db/ent/ent/gen/runtime"
)

type database struct {
	db *gen.Client
}

func New(cfg config.Database) *gen.Client {
	drv := &entsql.Driver{}

	switch cfg.Type {
	case "postgres", "postgresql", "psql", "pgsql", "pgx":
		db, err := sql.Open("pgx", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, "ent", cfg.Password, cfg.SSLMode))
		if err != nil {
			log.Fatal(err)
		}

		drv = entsql.OpenDB(dialect.Postgres, db)

	case "mysql", "mariadb":
		dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Name,
		)

		db, err := sql.Open(cfg.Type, dsn)
		if err != nil {
			log.Fatal(err)
		}

		drv = entsql.OpenDB(dialect.MySQL, db)

	default:
		log.Fatal(`Must choose a database driver: "postgres", "mariadb"`)
	}

	client := gen.NewClient(gen.Driver(drv))

	f, err := os.Create("./db/ent/migrate.sql")
	if err != nil {
		log.Fatalf("create migrate file: %v", err)
	}
	defer f.Close()

	//ctx := context.Background()
	//if err := client.Schema.WriteTo(ctx, f); err != nil {
	//	log.Fatalf("failed printing schema changes: %v", err)
	//}
	//
	//if err := client.Schema.Create(ctx); err != nil {
	//	log.Fatalf("failed creating schema resources: %v", err)
	//}

	return client
}

func NewRepo(db *gen.Client) *database {
	return &database{
		db: db,
	}
}
