// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"fmt"
	"godb/db/ent/ent/gen/country"
	"godb/db/ent/ent/gen/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CountryDelete is the builder for deleting a Country entity.
type CountryDelete struct {
	config
	hooks    []Hook
	mutation *CountryMutation
}

// Where appends a list predicates to the CountryDelete builder.
func (cd *CountryDelete) Where(ps ...predicate.Country) *CountryDelete {
	cd.mutation.Where(ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *CountryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cd.hooks) == 0 {
		affected, err = cd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CountryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cd.mutation = mutation
			affected, err = cd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cd.hooks) - 1; i >= 0; i-- {
			if cd.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = cd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CountryDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *CountryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: country.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: country.FieldID,
			},
		},
	}
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// CountryDeleteOne is the builder for deleting a single Country entity.
type CountryDeleteOne struct {
	cd *CountryDelete
}

// Exec executes the deletion query.
func (cdo *CountryDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{country.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CountryDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}
