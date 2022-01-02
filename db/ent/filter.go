package ent

import (
	"net/url"

	"godb/db/ent/ent/gen/predicate"
	"godb/db/ent/ent/gen/user"
	"godb/filter"
)

type Filter struct {
	Base filter.Filter

	Email           string
	FirstName       string
	FavouriteColour string

	PredicateUser []predicate.User
}

func filters(queries url.Values) *Filter {
	base := filter.New(queries)

	f := &Filter{
		Base: *base,

		Email:           queries.Get("email"),
		FirstName:       queries.Get("first_name"),
		FavouriteColour: queries.Get("favourite_colour"),
	}

	f.UserFilter()

	return f
}

// UserFilter automatically parses query parameters into ent's predicate if they
// exist.
func (f *Filter) UserFilter() {
	var predicateUser []predicate.User

	if f.Email != "" {
		predicateUser = append(predicateUser, user.EmailEQ(f.Email))
	}
	if f.FirstName != "" {
		// use `...ContainsFold` for case-insensitive search.
		predicateUser = append(predicateUser, user.FirstNameContainsFold(f.FirstName))
	}

	if f.FavouriteColour != "" {
		predicateUser = append(predicateUser, user.FavouriteColourEQ(user.FavouriteColour(f.FavouriteColour)))
	}

	f.PredicateUser = predicateUser
}
