package db

import (
	"net/url"

	"godb/filter"
	"godb/param"
)

type Filter struct {
	Base filter.Filter

	Email           string
	FirstName       string
	FavouriteColour string

	LastNames []string

	Transaction bool
}

func Filters(v url.Values) *Filter {
	f := filter.New(v)

	lastNames := param.ToStrSlice(v, "last_name")
	transaction := param.Bool(v, "transaction")

	return &Filter{
		Base: *f,

		Email:           v.Get("email"),
		FirstName:       v.Get("first_name"),
		FavouriteColour: v.Get("favourite_colour"),
		LastNames:       lastNames,
		Transaction:     transaction,
	}
}
