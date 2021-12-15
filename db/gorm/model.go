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
	//gorm.Model

	ID         uint   `json:"id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"-"`
}

type Country struct {
	//gorm.Model

	ID   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Address struct {
	//gorm.Model

	ID       uint           `json:"id"`
	Line1    string         `json:"line_1"`
	Line2    sql.NullString `json:"line_2"`
	Postcode sql.NullInt32  `json:"postcode"`
	City     sql.NullString `json:"city"`
	State    sql.NullString `json:"state"`
	Country  []Country      `json:"country"`
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
