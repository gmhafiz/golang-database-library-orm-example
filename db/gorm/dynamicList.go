package gorm

import (
	"context"
	"fmt"
	"strings"
)

func (r *repo) ListFilterByColumn(ctx context.Context, f *Filter) ([]*User, error) {
	var users []*User
	err := r.db.WithContext(ctx).
		Select([]string{"id", "first_name", "middle_name", "last_name", "email", "favourite_colour"}).
		Offset(f.Base.Offset).
		Limit(int(f.Base.Limit)).

		// Cannot use struct field when we want to use ILIKE clause.
		// Instead, we use Where() and Or() methods.
		//Find(&users, User{
		//	Email:           f.Email,
		//	FirstName:       f.FirstName, // this is not case-insensitive
		//	FavouriteColour: f.FavouriteColour,
		//}).

		Where("email = ? ", f.Email).
		Or("first_name ILIKE ?", f.FirstName).
		Or("favourite_colour = ?", f.FavouriteColour).

		// Compiler won't complain about missing Find() method!
		// Order is also important. If you put Find() in between of Limit() and
		// Where(), you will get a wrong result!
		Find(&users).
		Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repo) ListFilterSort(ctx context.Context, f *Filter) (users []*User, err error) {
	var orderClause []string
	for col, order := range f.Base.Sort {
		orderClause = append(orderClause, fmt.Sprintf("%s %s", col, order))
	}

	err = r.db.WithContext(ctx).
		Limit(int(f.Base.Limit)).
		Order(strings.Join(orderClause, ",")).
		Find(&users).
		Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repo) ListFilterPagination(ctx context.Context, f *Filter) (users []*User, err error) {
	err = r.db.Debug().WithContext(ctx).
		Limit(int(f.Base.Limit)).
		Offset(f.Base.Offset).
		Order("id").
		Find(&users). // order matters!
		Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
