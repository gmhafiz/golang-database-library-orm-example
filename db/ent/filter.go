package ent

import (
	"godb/param"
	"net/http"
	"strconv"

	"godb/db/ent/ent/gen/predicate"
	"godb/db/ent/ent/gen/user"
	"godb/filter"
)

type Filter struct {
	Base filter.Filter

	PaginateLastId int64

	Email           string
	FirstName       string
	LastName        []string
	FavouriteColour string

	PredicateUser []predicate.User
}

func filters(r *http.Request) *Filter {
	base := filter.New(r.URL.Query())

	paginateLastId, _ := strconv.ParseInt(r.URL.Query().Get("last_token"), 10, 64)

	f := &Filter{
		Base: *base,

		PaginateLastId: paginateLastId,

		Email:           r.URL.Query().Get("email"),
		FirstName:       r.URL.Query().Get("first_name"),
		LastName:        param.ToStrSlice(r, "last_name"),
		FavouriteColour: r.URL.Query().Get("favourite_colour"),
	}

	// Automatically parse at handler layer.
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
