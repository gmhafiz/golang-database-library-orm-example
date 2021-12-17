package gorm

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"godb/config"
)

type User struct {
	//gorm.Model

	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"-"`

	Addresses []Address `gorm:"many2many:user_addresses;"`
}

type Country struct {
	//gorm.Model

	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`

	Address []Address `json:"address" gorm:"foreignkey:country_id"`
}

type Address struct {
	//gorm.Model

	ID       int    `json:"id,omitempty"`
	Line1    string `json:"line_1,omitempty" gorm:"Column:line_1" `
	Line2    string `json:"line_2,omitempty" gorm:"Column:line_2" `
	Postcode int32  `json:"postcode,omitempty" gorm:"default:null" `
	City     string `json:"city,omitempty" gorm:"default:null" `
	State    string `json:"state,omitempty" gorm:"default:null" `

	CountryID int `json:"countries" json:"countryID,omitempty"`
}

func New(c config.Database) *gorm.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		"db_gorm",
		c.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	err = db.AutoMigrate(&User{}, &Country{}, &Address{})
	if err != nil {
		log.Panic(err)
	}

	return db
}
