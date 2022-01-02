package ent

import (
	"context"

	"godb/db/ent/ent/gen"
	"godb/db/ent/ent/gen/user"
	"godb/filter"
)

func (r *database) ListFilterByColumn(ctx context.Context, f *Filter) ([]*gen.User, error) {
	// We can put the logic to parse query here, or we make it as a method to
	// Filter struct (see db/ent/filter.go).
	//var predicateUser []predicate.User
	//
	//if f.Email != "" {
	//	predicateUser = append(predicateUser, user.EmailEQ(f.Email))
	//}
	//if f.FirstName != "" {
	//	predicateUser = append(predicateUser, user.FirstNameContainsFold(f.FirstName))
	//}
	//if f.FavouriteColour != "" {
	//	predicateUser = append(predicateUser, user.FavouriteColourEQ(user.FavouriteColour(f.FavouriteColour)))
	//}

	return r.db.User.Query().
		Where(f.PredicateUser...).
		Limit(int(f.Base.Limit)).
		Offset(f.Base.Offset).
		Order(gen.Asc(user.FieldID)).
		All(ctx)
}

func (r *database) ListFilterSort(ctx context.Context, f *Filter) ([]*gen.User, error) {
	var orderFunc []gen.OrderFunc

	for col, ord := range f.Base.Sort {
		if ord == filter.SqlAsc {
			orderFunc = append(orderFunc, gen.Asc(col))
		} else {
			orderFunc = append(orderFunc, gen.Desc(col))
		}
	}

	orderFunc = append(orderFunc, gen.Asc(user.FieldID))

	return r.db.User.Query().
		Order(orderFunc...).
		All(ctx)
}

func (r *database) ListFilterPagination(ctx context.Context, f *Filter) ([]*gen.User, error) {
	return r.db.User.Query().
		Limit(int(f.Base.Limit)).
		Offset(f.Base.Offset).
		// When using LIMIT, it is important to use an ORDER BY clause that
		// constrains the result rows into a unique order
		// https://www.postgresql.org/docs/14/queries-limit.html
		Order(gen.Asc(user.FieldID)).
		All(ctx)
}
