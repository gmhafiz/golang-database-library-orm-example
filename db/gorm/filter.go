package gorm

import (
	"github.com/volatiletech/null/v8"
	"net/url"

	"godb/filter"
)

type Filter struct {
	Base filter.Filter

	Email           string
	FirstName       string
	FavouriteColour null.String
}

func filters(queries url.Values) *Filter {
	f := filter.New(queries)

	F := &Filter{
		Base: *f,

		Email:     queries.Get("email"),
		FirstName: queries.Get("first_name"),
	}

	favColour := queries.Get("favourite_colour")
	if favColour == "" {
		F.FavouriteColour = null.String{
			String: "",
			Valid:  false,
		}
	} else {
		F.FavouriteColour = null.String{
			String: favColour,
			Valid:  true,
		}
	}

	return F
}
