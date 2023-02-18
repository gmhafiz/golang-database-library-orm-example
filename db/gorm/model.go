package gorm

import (
	"github.com/lib/pq"
	"godb/db"
	"log"

	"gorm.io/driver/mysql"
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
	//Tags       []string `json:"tags" gorm:"type:[]array"`
	//Tags []string `json:"tags" gorm:"type:text[]"`
	Tags pq.StringArray `json:"tags" gorm:"type:text[]"`
	//Tags            []string `json:"tags" gorm:"type:[]text"`
	FavouriteColour string `json:"favourite_colour"`

	Addresses []Address `json:"address" gorm:"many2many:user_addresses;"`
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

	CountryID int `json:"countryID,omitempty"`
}

func New(c config.Database) *gorm.DB {
	driver := &gorm.DB{}

	switch c.Type {
	case "postgres", "postgresql", "psql", "pgsql", "pgx":
		db, err := gorm.Open(postgres.Open(db.Dsn(c)), &gorm.Config{})
		if err != nil {
			log.Panic(err)
		}
		driver = db
	case "mysql", "mariadb":
		db, err := gorm.Open(mysql.Open(db.Dsn(c)), &gorm.Config{})
		if err != nil {
			log.Panic(err)
		}
		driver = db
	}

	//err = db.AutoMigrate(&User{}, &Country{}, &Address{})
	//if err != nil {
	//	log.Panic(err)
	//}

	return driver
}
