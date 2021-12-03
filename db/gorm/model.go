package gorm

import (
	"fmt"
	"godb/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type UserGorm struct {
	FirstName  string `gorm:"first_name"`
	MiddleName string `gorm:"middle_name"`
	LastName   string `gorm:"last_name"`
	Email      string `gorm:"email"`
	Password   string `gorm:"password"`

	gorm.Model
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

	err = db.AutoMigrate(&UserGorm{})
	if err != nil {
		log.Panic(err)
	}

	return db
}
