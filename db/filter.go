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

	LastName []string
}

func Filters(v url.Values) *Filter {
	f := filter.New(v)

	lastNames := param.ToStrSlice(v, "last_name")

	return &Filter{
		Base: *f,

		Email:           v.Get("email"),
		FirstName:       v.Get("first_name"),
		FavouriteColour: v.Get("favourite_colour"),
		LastName:        lastNames,
	}
}
