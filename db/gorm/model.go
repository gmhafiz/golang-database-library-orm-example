package gorm

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"godb/config"
)

type User struct {
	gorm.Model

	FirstName  string
	MiddleName sql.NullString
	LastName   string
	Email      string
	Password   string
}

type Country struct {
	gorm.Model

	Code string
	Name string
}

type Address struct {
	gorm.Model

	ID       uint
	Line1    string
	Line2    sql.NullString
	Postcode sql.NullInt32
	City     sql.NullString
	State    sql.NullString
	Country  []Country
	//Country  []Country `gorm:"foreignKey:ID"`
}

func New(c config.Database) *gorm.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
		c.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	db.Debug()

	//err = db.AutoMigrate(&User{}, &Country{}, &Address{})
	//if err != nil {
	//	log.Panic(err)
	//}

	return db
}
