package sqlc

import (
	"net/url"

	"godb/filter"
)

type Filter struct {
	Base filter.Filter

	Email           string
	FirstName       string
	FavouriteColour string
}

func filters(queries url.Values) *Filter {
	f := filter.New(queries)

	return &Filter{
		Base: *f,

		Email:           queries.Get("email"),
		FirstName:       queries.Get("first_name"),
		FavouriteColour: queries.Get("favourite_colour"),
	}
}
