package db

import (
	"godb/filter"
	"godb/param"
	"net/http"
)

type Filter struct {
	Base filter.Filter

	Email           string
	FirstName       string
	FavouriteColour string

	LastNames []string

	Transaction bool

	Select string
}

func Filters(r *http.Request) *Filter {
	f := filter.New(r.URL.Query())

	lastNames := param.ToStrSlice(r.URL.Query(), "last_name")
	transaction := param.Bool(r.URL.Query(), "transaction")

	return &Filter{
		Base: *f,

		Email:           r.URL.Query().Get("email"),
		FirstName:       r.URL.Query().Get("first_name"),
		FavouriteColour: r.URL.Query().Get("favourite_colour"),
		LastNames:       lastNames,
		Transaction:     transaction,
		Select:          r.URL.Query().Get("select"),
	}
}
