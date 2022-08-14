package sqlx

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	"godb/config"
)

func New(c config.Database) *sqlx.DB {
	db, err := sqlx.Open(c.Type, Dsn(c))
	if err != nil {
		log.Fatal(err)
	}

	Alive(db.DB)

	return db
}

func Dsn(c config.Database) string {
	var dsn string
	switch c.Type {
	case "postgres", "postgresql", "psql", "pgsql", "pgx":
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
			c.User,
			c.Password,
			c.Host,
			c.Port,
			c.Name,
			c.SSLMode,
		)
	case "mysql", "mariadb":
		dsn = fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
			c.User,
			c.Password,
			c.Host,
			c.Port,
			c.Name,
		)
	default:
		log.Fatal(`Must choose a database driver: "postgres", "mariadb"`)
	}

	return dsn
}

func Alive(db *sql.DB) {
	log.Println("Connecting to repository... ")
	for {
		// Ping by itself is un-reliable, the connections are cached. This
		// ensures that the repository is still running by executing a harmless
		// dummy query against it.
		//
		// Also, we perform an exponential backoff when querying the repository
		// to spread our retries.
		_, err := db.Exec("SELECT true")
		if err == nil {
			log.Println("Database connected")
			return
		}

		base, capacity := time.Second, time.Minute
		for backoff := base; err != nil; backoff <<= 1 {
			if backoff > capacity {
				backoff = capacity
			}

			// A pseudo-random number generator here is fine. No need to be
			// cryptographically secure. Ignore with the following comment:
			/* #nosec */
			jitter := rand.Int63n(int64(backoff * 3))
			sleep := base + time.Duration(jitter)
			time.Sleep(sleep)
			_, err := db.Exec("SELECT true")
			if err == nil {
				log.Println("Database connected")
				return
			}
		}
	}
}
