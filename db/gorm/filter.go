package gorm

import (
	"github.com/volatiletech/null/v8"
	"godb/filter"
	"godb/param"
	"net/http"
)

type Filter struct {
	Base filter.Filter

	Email           string
	FirstName       string
	FavouriteColour null.String

	LastName []string
}

func filters(r *http.Request) *Filter {
	f := filter.New(r.URL.Query())

	F := &Filter{
		Base: *f,

		Email:     r.URL.Query().Get("email"),
		FirstName: r.URL.Query().Get("first_name"),
		LastName:  param.ToStrSlice(r, "last_name"),
	}

	favColour := r.URL.Query().Get("favourite_colour")
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
