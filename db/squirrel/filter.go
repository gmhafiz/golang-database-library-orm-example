package squirrel

import (
	"net/http"

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

func filters(r *http.Request) *Filter {
	f := filter.New(r.URL.Query())

	lastNames := param.ToStrSlice(r, "last_name")

	return &Filter{
		Base: *f,

		Email:           r.URL.Query().Get("email"),
		FirstName:       r.URL.Query().Get("first_name"),
		FavouriteColour: r.URL.Query().Get("favourite_colour"),

		LastName: lastNames,
	}
}
