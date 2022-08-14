package ent

import (
	"net/url"
	"strconv"

	"godb/db"
	"godb/db/ent/ent/gen/predicate"
	"godb/db/ent/ent/gen/user"
)

type filter struct {
	*db.Filter

	PredicateUser  []predicate.User
	PaginateLastID uint
}

func filters(v url.Values) *filter {
	paginateLastId, _ := strconv.ParseInt(v.Get("last_token"), 10, 64)

	f := &filter{
		db.Filters(v),
		nil,
		uint(paginateLastId),
	}

	// Automatically parse at handler layer.
	f.UserFilter()

	return f
}

// UserFilter automatically parses query parameters into ent's predicate if they
// exist.
func (f *filter) UserFilter() {
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
