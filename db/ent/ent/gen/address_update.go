// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"godb/db/ent/ent/gen/address"
	"godb/db/ent/ent/gen/country"
	"godb/db/ent/ent/gen/predicate"
	"godb/db/ent/ent/gen/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AddressUpdate is the builder for updating Address entities.
type AddressUpdate struct {
	config
	hooks    []Hook
	mutation *AddressMutation
}

// Where appends a list predicates to the AddressUpdate builder.
func (au *AddressUpdate) Where(ps ...predicate.Address) *AddressUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetLine1 sets the "line_1" field.
func (au *AddressUpdate) SetLine1(s string) *AddressUpdate {
	au.mutation.SetLine1(s)
	return au
}

// SetLine2 sets the "line_2" field.
func (au *AddressUpdate) SetLine2(s string) *AddressUpdate {
	au.mutation.SetLine2(s)
	return au
}

// SetNillableLine2 sets the "line_2" field if the given value is not nil.
func (au *AddressUpdate) SetNillableLine2(s *string) *AddressUpdate {
	if s != nil {
		au.SetLine2(*s)
	}
	return au
}

// ClearLine2 clears the value of the "line_2" field.
func (au *AddressUpdate) ClearLine2() *AddressUpdate {
	au.mutation.ClearLine2()
	return au
}

// SetPostcode sets the "postcode" field.
func (au *AddressUpdate) SetPostcode(i int) *AddressUpdate {
	au.mutation.ResetPostcode()
	au.mutation.SetPostcode(i)
	return au
}

// AddPostcode adds i to the "postcode" field.
func (au *AddressUpdate) AddPostcode(i int) *AddressUpdate {
	au.mutation.AddPostcode(i)
	return au
}

// SetState sets the "state" field.
func (au *AddressUpdate) SetState(s string) *AddressUpdate {
	au.mutation.SetState(s)
	return au
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (au *AddressUpdate) SetCountryID(id uint) *AddressUpdate {
	au.mutation.SetCountryID(id)
	return au
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (au *AddressUpdate) SetNillableCountryID(id *uint) *AddressUpdate {
	if id != nil {
		au = au.SetCountryID(*id)
	}
	return au
}

// SetCountry sets the "country" edge to the Country entity.
func (au *AddressUpdate) SetCountry(c *Country) *AddressUpdate {
	return au.SetCountryID(c.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (au *AddressUpdate) AddUserIDs(ids ...uint) *AddressUpdate {
	au.mutation.AddUserIDs(ids...)
	return au
}

// AddUsers adds the "users" edges to the User entity.
func (au *AddressUpdate) AddUsers(u ...*User) *AddressUpdate {
	ids := make([]uint, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return au.AddUserIDs(ids...)
}

// Mutation returns the AddressMutation object of the builder.
func (au *AddressUpdate) Mutation() *AddressMutation {
	return au.mutation
}

// ClearCountry clears the "country" edge to the Country entity.
func (au *AddressUpdate) ClearCountry() *AddressUpdate {
	au.mutation.ClearCountry()
	return au
}

// ClearUsers clears all "users" edges to the User entity.
func (au *AddressUpdate) ClearUsers() *AddressUpdate {
	au.mutation.ClearUsers()
	return au
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (au *AddressUpdate) RemoveUserIDs(ids ...uint) *AddressUpdate {
	au.mutation.RemoveUserIDs(ids...)
	return au
}

// RemoveUsers removes "users" edges to User entities.
func (au *AddressUpdate) RemoveUsers(u ...*User) *AddressUpdate {
	ids := make([]uint, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return au.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AddressUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AddressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AddressUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AddressUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AddressUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AddressUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   address.Table,
			Columns: address.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: address.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Line1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldLine1,
		})
	}
	if value, ok := au.mutation.Line2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldLine2,
		})
	}
	if au.mutation.Line2Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: address.FieldLine2,
		})
	}
	if value, ok := au.mutation.Postcode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: address.FieldPostcode,
		})
	}
	if value, ok := au.mutation.AddedPostcode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: address.FieldPostcode,
		})
	}
	if value, ok := au.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldState,
		})
	}
	if au.mutation.CountryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   address.CountryTable,
			Columns: []string{address.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: country.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   address.CountryTable,
			Columns: []string{address.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: country.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   address.UsersTable,
			Columns: address.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedUsersIDs(); len(nodes) > 0 && !au.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   address.UsersTable,
			Columns: address.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   address.UsersTable,
			Columns: address.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{address.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AddressUpdateOne is the builder for updating a single Address entity.
type AddressUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AddressMutation
}

// SetLine1 sets the "line_1" field.
func (auo *AddressUpdateOne) SetLine1(s string) *AddressUpdateOne {
	auo.mutation.SetLine1(s)
	return auo
}

// SetLine2 sets the "line_2" field.
func (auo *AddressUpdateOne) SetLine2(s string) *AddressUpdateOne {
	auo.mutation.SetLine2(s)
	return auo
}

// SetNillableLine2 sets the "line_2" field if the given value is not nil.
func (auo *AddressUpdateOne) SetNillableLine2(s *string) *AddressUpdateOne {
	if s != nil {
		auo.SetLine2(*s)
	}
	return auo
}

// ClearLine2 clears the value of the "line_2" field.
func (auo *AddressUpdateOne) ClearLine2() *AddressUpdateOne {
	auo.mutation.ClearLine2()
	return auo
}

// SetPostcode sets the "postcode" field.
func (auo *AddressUpdateOne) SetPostcode(i int) *AddressUpdateOne {
	auo.mutation.ResetPostcode()
	auo.mutation.SetPostcode(i)
	return auo
}

// AddPostcode adds i to the "postcode" field.
func (auo *AddressUpdateOne) AddPostcode(i int) *AddressUpdateOne {
	auo.mutation.AddPostcode(i)
	return auo
}

// SetState sets the "state" field.
func (auo *AddressUpdateOne) SetState(s string) *AddressUpdateOne {
	auo.mutation.SetState(s)
	return auo
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (auo *AddressUpdateOne) SetCountryID(id uint) *AddressUpdateOne {
	auo.mutation.SetCountryID(id)
	return auo
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (auo *AddressUpdateOne) SetNillableCountryID(id *uint) *AddressUpdateOne {
	if id != nil {
		auo = auo.SetCountryID(*id)
	}
	return auo
}

// SetCountry sets the "country" edge to the Country entity.
func (auo *AddressUpdateOne) SetCountry(c *Country) *AddressUpdateOne {
	return auo.SetCountryID(c.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (auo *AddressUpdateOne) AddUserIDs(ids ...uint) *AddressUpdateOne {
	auo.mutation.AddUserIDs(ids...)
	return auo
}

// AddUsers adds the "users" edges to the User entity.
func (auo *AddressUpdateOne) AddUsers(u ...*User) *AddressUpdateOne {
	ids := make([]uint, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return auo.AddUserIDs(ids...)
}

// Mutation returns the AddressMutation object of the builder.
func (auo *AddressUpdateOne) Mutation() *AddressMutation {
	return auo.mutation
}

// ClearCountry clears the "country" edge to the Country entity.
func (auo *AddressUpdateOne) ClearCountry() *AddressUpdateOne {
	auo.mutation.ClearCountry()
	return auo
}

// ClearUsers clears all "users" edges to the User entity.
func (auo *AddressUpdateOne) ClearUsers() *AddressUpdateOne {
	auo.mutation.ClearUsers()
	return auo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (auo *AddressUpdateOne) RemoveUserIDs(ids ...uint) *AddressUpdateOne {
	auo.mutation.RemoveUserIDs(ids...)
	return auo
}

// RemoveUsers removes "users" edges to User entities.
func (auo *AddressUpdateOne) RemoveUsers(u ...*User) *AddressUpdateOne {
	ids := make([]uint, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return auo.RemoveUserIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AddressUpdateOne) Select(field string, fields ...string) *AddressUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Address entity.
func (auo *AddressUpdateOne) Save(ctx context.Context) (*Address, error) {
	var (
		err  error
		node *Address
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AddressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Address)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AddressMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AddressUpdateOne) SaveX(ctx context.Context) *Address {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AddressUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AddressUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AddressUpdateOne) sqlSave(ctx context.Context) (_node *Address, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   address.Table,
			Columns: address.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: address.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`gen: missing "Address.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, address.FieldID)
		for _, f := range fields {
			if !address.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
			}
			if f != address.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Line1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldLine1,
		})
	}
	if value, ok := auo.mutation.Line2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldLine2,
		})
	}
	if auo.mutation.Line2Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: address.FieldLine2,
		})
	}
	if value, ok := auo.mutation.Postcode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: address.FieldPostcode,
		})
	}
	if value, ok := auo.mutation.AddedPostcode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: address.FieldPostcode,
		})
	}
	if value, ok := auo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldState,
		})
	}
	if auo.mutation.CountryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   address.CountryTable,
			Columns: []string{address.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: country.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   address.CountryTable,
			Columns: []string{address.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: country.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   address.UsersTable,
			Columns: address.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !auo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   address.UsersTable,
			Columns: address.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   address.UsersTable,
			Columns: address.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Address{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{address.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
