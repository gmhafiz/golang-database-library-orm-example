package ent

import (
	"context"

	"godb/db/ent/ent/gen"
	"godb/db/ent/ent/gen/predicate"
	"godb/db/ent/ent/gen/user"
	baseFilter "godb/filter"
)

// ListFilterByColumn filters using `predicate`.
func (r *database) ListFilterByColumn(ctx context.Context, f *filter) ([]*gen.User, error) {
	var predicateUser []predicate.User
	if f.Email != "" {
		predicateUser = append(predicateUser, user.EmailEQ(f.Email))
	}
	if f.FirstName != "" {
		//predicateUser = append(predicateUser, user.FirstNameContainsFold(f.FirstName))
		predicateUser = append(predicateUser, user.FirstNameEQ(f.FirstName))
	}
	if f.FavouriteColour != "" {
		predicateUser = append(predicateUser, user.FavouriteColourEQ(user.FavouriteColour(f.FavouriteColour)))
	}

	return r.db.Debug().User.Query().
		Where(predicateUser...).
		Order(gen.Asc(user.FieldID)).
		Limit(f.Base.Limit).
		Offset(f.Base.Offset).
		All(ctx)
}

func (r *database) ListFilterSort(ctx context.Context, f *filter) ([]*gen.User, error) {
	var orderFunc []gen.OrderFunc

	for col, ord := range f.Base.Sort {
		if ord == baseFilter.SqlAsc {
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

// ListFilterPagination is a simple pagination using OFFSET and LIMIT. Fine
// for small dataset but potentially slow if dataset is large, and you need the
// offset to start from a large number. See ListFilterPaginationByID below
// for pagination using by cursor method.
func (r *database) ListFilterPagination(ctx context.Context, f *filter) ([]*gen.User, error) {
	query := r.db.User.Query()
	if f.Base.Limit != 0 && !f.Base.DisablePaging {
		query = query.Limit(f.Base.Limit)
	}
	if f.Base.Offset != 0 && !f.Base.DisablePaging {
		query = query.Offset(f.Base.Offset)
	}
	resp, err := query.
		// When using LIMIT, it is important to use an ORDER BY clause that
		// constrains the result rows into a unique order
		// https://www.postgresql.org/docs/14/queries-limit.html
		Order(gen.Asc(user.FieldID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ListFilterPaginationByID Pagination by using OFFSET gives a huge performance
// penalty when dataset is large because the database does a full table scan.
// Requires 3 query parameters to be sent:
//   1. Last token
//   2. The column that was ordered
//   3. Its direction
/*
SELECT …
FROM …
WHERE id > {last_token}
ORDER by {column} {direction}
LIMIT 3
*/
func (r *database) ListFilterPaginationByID(ctx context.Context, f *filter) ([]*gen.User, error) {
	var orderFunc []gen.OrderFunc
	for col, ord := range f.Base.Sort {
		if ord == baseFilter.SqlAsc {
			orderFunc = append(orderFunc, gen.Asc(col))
		} else {
			orderFunc = append(orderFunc, gen.Desc(col))
		}
	}

	orderFunc = append(orderFunc, gen.Asc(user.FieldID))

	return r.db.User.Query().Where(user.IDGT(f.PaginateLastID)).
		Limit(f.Base.Limit).
		Order(orderFunc...).
		All(ctx)
}
