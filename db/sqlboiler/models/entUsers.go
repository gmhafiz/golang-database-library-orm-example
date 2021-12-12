// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// EntUser is an object representing the database table.
type EntUser struct {
	ID         int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	FirstName  string      `boil:"first_name" json:"first_name" toml:"first_name" yaml:"first_name"`
	MiddleName null.String `boil:"middle_name" json:"middle_name,omitempty" toml:"middle_name" yaml:"middle_name,omitempty"`
	LastName   string      `boil:"last_name" json:"last_name" toml:"last_name" yaml:"last_name"`
	Email      string      `boil:"email" json:"email" toml:"email" yaml:"email"`
	Password   string      `boil:"password" json:"password" toml:"password" yaml:"password"`

	R *entUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L entUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EntUserColumns = struct {
	ID         string
	FirstName  string
	MiddleName string
	LastName   string
	Email      string
	Password   string
}{
	ID:         "id",
	FirstName:  "first_name",
	MiddleName: "middle_name",
	LastName:   "last_name",
	Email:      "email",
	Password:   "password",
}

var EntUserTableColumns = struct {
	ID         string
	FirstName  string
	MiddleName string
	LastName   string
	Email      string
	Password   string
}{
	ID:         "entUsers.id",
	FirstName:  "entUsers.first_name",
	MiddleName: "entUsers.middle_name",
	LastName:   "entUsers.last_name",
	Email:      "entUsers.email",
	Password:   "entUsers.password",
}

// Generated where

var EntUserWhere = struct {
	ID         whereHelperint64
	FirstName  whereHelperstring
	MiddleName whereHelpernull_String
	LastName   whereHelperstring
	Email      whereHelperstring
	Password   whereHelperstring
}{
	ID:         whereHelperint64{field: "\"entUsers\".\"id\""},
	FirstName:  whereHelperstring{field: "\"entUsers\".\"first_name\""},
	MiddleName: whereHelpernull_String{field: "\"entUsers\".\"middle_name\""},
	LastName:   whereHelperstring{field: "\"entUsers\".\"last_name\""},
	Email:      whereHelperstring{field: "\"entUsers\".\"email\""},
	Password:   whereHelperstring{field: "\"entUsers\".\"password\""},
}

// EntUserRels is where relationship names are stored.
var EntUserRels = struct {
	AddressEntAddresses string
}{
	AddressEntAddresses: "AddressEntAddresses",
}

// entUserR is where relationships are stored.
type entUserR struct {
	AddressEntAddresses EntAddressSlice `boil:"AddressEntAddresses" json:"AddressEntAddresses" toml:"AddressEntAddresses" yaml:"AddressEntAddresses"`
}

// NewStruct creates a new relationship struct
func (*entUserR) NewStruct() *entUserR {
	return &entUserR{}
}

// entUserL is where Load methods for each relationship are stored.
type entUserL struct{}

var (
	entUserAllColumns            = []string{"id", "first_name", "middle_name", "last_name", "email", "password"}
	entUserColumnsWithoutDefault = []string{"first_name", "middle_name", "last_name", "email", "password"}
	entUserColumnsWithDefault    = []string{"id"}
	entUserPrimaryKeyColumns     = []string{"id"}
)

type (
	// EntUserSlice is an alias for a slice of pointers to EntUser.
	// This should almost always be used instead of []EntUser.
	EntUserSlice []*EntUser
	// EntUserHook is the signature for custom EntUser hook methods
	EntUserHook func(context.Context, boil.ContextExecutor, *EntUser) error

	entUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	entUserType                 = reflect.TypeOf(&EntUser{})
	entUserMapping              = queries.MakeStructMapping(entUserType)
	entUserPrimaryKeyMapping, _ = queries.BindMapping(entUserType, entUserMapping, entUserPrimaryKeyColumns)
	entUserInsertCacheMut       sync.RWMutex
	entUserInsertCache          = make(map[string]insertCache)
	entUserUpdateCacheMut       sync.RWMutex
	entUserUpdateCache          = make(map[string]updateCache)
	entUserUpsertCacheMut       sync.RWMutex
	entUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var entUserBeforeInsertHooks []EntUserHook
var entUserBeforeUpdateHooks []EntUserHook
var entUserBeforeDeleteHooks []EntUserHook
var entUserBeforeUpsertHooks []EntUserHook

var entUserAfterInsertHooks []EntUserHook
var entUserAfterSelectHooks []EntUserHook
var entUserAfterUpdateHooks []EntUserHook
var entUserAfterDeleteHooks []EntUserHook
var entUserAfterUpsertHooks []EntUserHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *EntUser) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entUserBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *EntUser) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entUserBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *EntUser) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entUserBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *EntUser) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entUserBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *EntUser) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entUserAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *EntUser) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entUserAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *EntUser) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entUserAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *EntUser) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entUserAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *EntUser) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entUserAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEntUserHook registers your hook function for all future operations.
func AddEntUserHook(hookPoint boil.HookPoint, entUserHook EntUserHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		entUserBeforeInsertHooks = append(entUserBeforeInsertHooks, entUserHook)
	case boil.BeforeUpdateHook:
		entUserBeforeUpdateHooks = append(entUserBeforeUpdateHooks, entUserHook)
	case boil.BeforeDeleteHook:
		entUserBeforeDeleteHooks = append(entUserBeforeDeleteHooks, entUserHook)
	case boil.BeforeUpsertHook:
		entUserBeforeUpsertHooks = append(entUserBeforeUpsertHooks, entUserHook)
	case boil.AfterInsertHook:
		entUserAfterInsertHooks = append(entUserAfterInsertHooks, entUserHook)
	case boil.AfterSelectHook:
		entUserAfterSelectHooks = append(entUserAfterSelectHooks, entUserHook)
	case boil.AfterUpdateHook:
		entUserAfterUpdateHooks = append(entUserAfterUpdateHooks, entUserHook)
	case boil.AfterDeleteHook:
		entUserAfterDeleteHooks = append(entUserAfterDeleteHooks, entUserHook)
	case boil.AfterUpsertHook:
		entUserAfterUpsertHooks = append(entUserAfterUpsertHooks, entUserHook)
	}
}

// One returns a single entUser record from the query.
func (q entUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EntUser, error) {
	o := &EntUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for entUsers")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all EntUser records from the query.
func (q entUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (EntUserSlice, error) {
	var o []*EntUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EntUser slice")
	}

	if len(entUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all EntUser records in the query.
func (q entUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count entUsers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q entUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if entUsers exists")
	}

	return count > 0, nil
}

// AddressEntAddresses retrieves all the entAddress's EntAddresses with an executor via id column.
func (o *EntUser) AddressEntAddresses(mods ...qm.QueryMod) entAddressQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"address_entUsers\" on \"entAddresses\".\"id\" = \"address_entUsers\".\"address_id\""),
		qm.Where("\"address_entUsers\".\"user_id\"=?", o.ID),
	)

	query := EntAddresses(queryMods...)
	queries.SetFrom(query.Query, "\"entAddresses\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"entAddresses\".*"})
	}

	return query
}

// LoadAddressEntAddresses allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (entUserL) LoadAddressEntAddresses(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEntUser interface{}, mods queries.Applicator) error {
	var slice []*EntUser
	var object *EntUser

	if singular {
		object = maybeEntUser.(*EntUser)
	} else {
		slice = *maybeEntUser.(*[]*EntUser)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &entUserR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &entUserR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.Select("\"entAddresses\".id, \"entAddresses\".line_1, \"entAddresses\".line_2, \"entAddresses\".postcode, \"entAddresses\".state, \"entAddresses\".country_ent_addresses, \"a\".\"user_id\""),
		qm.From("\"entAddresses\""),
		qm.InnerJoin("\"address_entUsers\" as \"a\" on \"entAddresses\".\"id\" = \"a\".\"address_id\""),
		qm.WhereIn("\"a\".\"user_id\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load entAddresses")
	}

	var resultSlice []*EntAddress

	var localJoinCols []int64
	for results.Next() {
		one := new(EntAddress)
		var localJoinCol int64

		err = results.Scan(&one.ID, &one.Line1, &one.Line2, &one.Postcode, &one.State, &one.CountryEntAddresses, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for entAddresses")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice entAddresses")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on entAddresses")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for entAddresses")
	}

	if len(entAddressAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.AddressEntAddresses = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &entAddressR{}
			}
			foreign.R.UserEntUsers = append(foreign.R.UserEntUsers, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.AddressEntAddresses = append(local.R.AddressEntAddresses, foreign)
				if foreign.R == nil {
					foreign.R = &entAddressR{}
				}
				foreign.R.UserEntUsers = append(foreign.R.UserEntUsers, local)
				break
			}
		}
	}

	return nil
}

// AddAddressEntAddresses adds the given related objects to the existing relationships
// of the entUser, optionally inserting them as new records.
// Appends related to o.R.AddressEntAddresses.
// Sets related.R.UserEntUsers appropriately.
func (o *EntUser) AddAddressEntAddresses(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*EntAddress) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"address_entUsers\" (\"user_id\", \"address_id\") values ($1, $2)"
		values := []interface{}{o.ID, rel.ID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, query)
			fmt.Fprintln(writer, values)
		}
		_, err = exec.ExecContext(ctx, query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &entUserR{
			AddressEntAddresses: related,
		}
	} else {
		o.R.AddressEntAddresses = append(o.R.AddressEntAddresses, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &entAddressR{
				UserEntUsers: EntUserSlice{o},
			}
		} else {
			rel.R.UserEntUsers = append(rel.R.UserEntUsers, o)
		}
	}
	return nil
}

// SetAddressEntAddresses removes all previously related items of the
// entUser replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.UserEntUsers's AddressEntAddresses accordingly.
// Replaces o.R.AddressEntAddresses with related.
// Sets related.R.UserEntUsers's AddressEntAddresses accordingly.
func (o *EntUser) SetAddressEntAddresses(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*EntAddress) error {
	query := "delete from \"address_entUsers\" where \"user_id\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeAddressEntAddressesFromUserEntUsersSlice(o, related)
	if o.R != nil {
		o.R.AddressEntAddresses = nil
	}
	return o.AddAddressEntAddresses(ctx, exec, insert, related...)
}

// RemoveAddressEntAddresses relationships from objects passed in.
// Removes related items from R.AddressEntAddresses (uses pointer comparison, removal does not keep order)
// Sets related.R.UserEntUsers.
func (o *EntUser) RemoveAddressEntAddresses(ctx context.Context, exec boil.ContextExecutor, related ...*EntAddress) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	query := fmt.Sprintf(
		"delete from \"address_entUsers\" where \"user_id\" = $1 and \"address_id\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeAddressEntAddressesFromUserEntUsersSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.AddressEntAddresses {
			if rel != ri {
				continue
			}

			ln := len(o.R.AddressEntAddresses)
			if ln > 1 && i < ln-1 {
				o.R.AddressEntAddresses[i] = o.R.AddressEntAddresses[ln-1]
			}
			o.R.AddressEntAddresses = o.R.AddressEntAddresses[:ln-1]
			break
		}
	}

	return nil
}

func removeAddressEntAddressesFromUserEntUsersSlice(o *EntUser, related []*EntAddress) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.UserEntUsers {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.UserEntUsers)
			if ln > 1 && i < ln-1 {
				rel.R.UserEntUsers[i] = rel.R.UserEntUsers[ln-1]
			}
			rel.R.UserEntUsers = rel.R.UserEntUsers[:ln-1]
			break
		}
	}
}

// EntUsers retrieves all the records using an executor.
func EntUsers(mods ...qm.QueryMod) entUserQuery {
	mods = append(mods, qm.From("\"entUsers\""))
	return entUserQuery{NewQuery(mods...)}
}

// FindEntUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEntUser(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*EntUser, error) {
	entUserObj := &EntUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"entUsers\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, entUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from entUsers")
	}

	if err = entUserObj.doAfterSelectHooks(ctx, exec); err != nil {
		return entUserObj, err
	}

	return entUserObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EntUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no entUsers provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(entUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	entUserInsertCacheMut.RLock()
	cache, cached := entUserInsertCache[key]
	entUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			entUserAllColumns,
			entUserColumnsWithDefault,
			entUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(entUserType, entUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(entUserType, entUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"entUsers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"entUsers\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into entUsers")
	}

	if !cached {
		entUserInsertCacheMut.Lock()
		entUserInsertCache[key] = cache
		entUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the EntUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EntUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	entUserUpdateCacheMut.RLock()
	cache, cached := entUserUpdateCache[key]
	entUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			entUserAllColumns,
			entUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update entUsers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"entUsers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, entUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(entUserType, entUserMapping, append(wl, entUserPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update entUsers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for entUsers")
	}

	if !cached {
		entUserUpdateCacheMut.Lock()
		entUserUpdateCache[key] = cache
		entUserUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q entUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for entUsers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for entUsers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EntUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), entUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"entUsers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, entUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in entUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all entUser")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EntUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no entUsers provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(entUserColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	entUserUpsertCacheMut.RLock()
	cache, cached := entUserUpsertCache[key]
	entUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			entUserAllColumns,
			entUserColumnsWithDefault,
			entUserColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			entUserAllColumns,
			entUserPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert entUsers, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(entUserPrimaryKeyColumns))
			copy(conflict, entUserPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"entUsers\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(entUserType, entUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(entUserType, entUserMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert entUsers")
	}

	if !cached {
		entUserUpsertCacheMut.Lock()
		entUserUpsertCache[key] = cache
		entUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single EntUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EntUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EntUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), entUserPrimaryKeyMapping)
	sql := "DELETE FROM \"entUsers\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from entUsers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for entUsers")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q entUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no entUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from entUsers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for entUsers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EntUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(entUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), entUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"entUsers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, entUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from entUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for entUsers")
	}

	if len(entUserAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *EntUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEntUser(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EntUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EntUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), entUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"entUsers\".* FROM \"entUsers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, entUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EntUserSlice")
	}

	*o = slice

	return nil
}

// EntUserExists checks if the EntUser row exists.
func EntUserExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"entUsers\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if entUsers exists")
	}

	return exists, nil
}
