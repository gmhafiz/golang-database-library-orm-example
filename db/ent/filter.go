package ent

import (
	"net/http"
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

func filters(r *http.Request) *filter {
	paginateLastId, _ := strconv.ParseInt(r.URL.Query().Get("last_token"), 10, 64)

	f := &filter{
		db.Filters(r),
		nil,
		uint(paginateLastId),
	}

	// Optionally, automatically parse at handler layer.
	//f.UserFilter()

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
