// Code generated by SQLBoiler 3.5.0-gct (https://github.com/thrasher-corp/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlite3

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
	"github.com/thrasher-corp/sqlboiler/queries/qm"
	"github.com/thrasher-corp/sqlboiler/queries/qmhelper"
	"github.com/thrasher-corp/sqlboiler/strmangle"
	"github.com/volatiletech/null"
)

// WithdrawalCrypto is an object representing the database table.
type WithdrawalCrypto struct {
	ID                  int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Address             string      `boil:"address" json:"address" toml:"address" yaml:"address"`
	AddressTag          null.String `boil:"address_tag" json:"address_tag,omitempty" toml:"address_tag" yaml:"address_tag,omitempty"`
	Fee                 float64     `boil:"fee" json:"fee" toml:"fee" yaml:"fee"`
	WithdrawalHistoryID string      `boil:"withdrawal_history_id" json:"withdrawal_history_id" toml:"withdrawal_history_id" yaml:"withdrawal_history_id"`

	R *withdrawalCryptoR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L withdrawalCryptoL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var WithdrawalCryptoColumns = struct {
	ID                  string
	Address             string
	AddressTag          string
	Fee                 string
	WithdrawalHistoryID string
}{
	ID:                  "id",
	Address:             "address",
	AddressTag:          "address_tag",
	Fee:                 "fee",
	WithdrawalHistoryID: "withdrawal_history_id",
}

// Generated where

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var WithdrawalCryptoWhere = struct {
	ID                  whereHelperint64
	Address             whereHelperstring
	AddressTag          whereHelpernull_String
	Fee                 whereHelperfloat64
	WithdrawalHistoryID whereHelperstring
}{
	ID:                  whereHelperint64{field: "\"withdrawal_crypto\".\"id\""},
	Address:             whereHelperstring{field: "\"withdrawal_crypto\".\"address\""},
	AddressTag:          whereHelpernull_String{field: "\"withdrawal_crypto\".\"address_tag\""},
	Fee:                 whereHelperfloat64{field: "\"withdrawal_crypto\".\"fee\""},
	WithdrawalHistoryID: whereHelperstring{field: "\"withdrawal_crypto\".\"withdrawal_history_id\""},
}

// WithdrawalCryptoRels is where relationship names are stored.
var WithdrawalCryptoRels = struct {
	WithdrawalHistory string
}{
	WithdrawalHistory: "WithdrawalHistory",
}

// withdrawalCryptoR is where relationships are stored.
type withdrawalCryptoR struct {
	WithdrawalHistory *WithdrawalHistory
}

// NewStruct creates a new relationship struct
func (*withdrawalCryptoR) NewStruct() *withdrawalCryptoR {
	return &withdrawalCryptoR{}
}

// withdrawalCryptoL is where Load methods for each relationship are stored.
type withdrawalCryptoL struct{}

var (
	withdrawalCryptoAllColumns            = []string{"id", "address", "address_tag", "fee", "withdrawal_history_id"}
	withdrawalCryptoColumnsWithoutDefault = []string{"address", "address_tag", "fee", "withdrawal_history_id"}
	withdrawalCryptoColumnsWithDefault    = []string{"id"}
	withdrawalCryptoPrimaryKeyColumns     = []string{"id"}
)

type (
	// WithdrawalCryptoSlice is an alias for a slice of pointers to WithdrawalCrypto.
	// This should generally be used opposed to []WithdrawalCrypto.
	WithdrawalCryptoSlice []*WithdrawalCrypto
	// WithdrawalCryptoHook is the signature for custom WithdrawalCrypto hook methods
	WithdrawalCryptoHook func(context.Context, boil.ContextExecutor, *WithdrawalCrypto) error

	withdrawalCryptoQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	withdrawalCryptoType                 = reflect.TypeOf(&WithdrawalCrypto{})
	withdrawalCryptoMapping              = queries.MakeStructMapping(withdrawalCryptoType)
	withdrawalCryptoPrimaryKeyMapping, _ = queries.BindMapping(withdrawalCryptoType, withdrawalCryptoMapping, withdrawalCryptoPrimaryKeyColumns)
	withdrawalCryptoInsertCacheMut       sync.RWMutex
	withdrawalCryptoInsertCache          = make(map[string]insertCache)
	withdrawalCryptoUpdateCacheMut       sync.RWMutex
	withdrawalCryptoUpdateCache          = make(map[string]updateCache)
	withdrawalCryptoUpsertCacheMut       sync.RWMutex
	withdrawalCryptoUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var withdrawalCryptoBeforeInsertHooks []WithdrawalCryptoHook
var withdrawalCryptoBeforeUpdateHooks []WithdrawalCryptoHook
var withdrawalCryptoBeforeDeleteHooks []WithdrawalCryptoHook
var withdrawalCryptoBeforeUpsertHooks []WithdrawalCryptoHook

var withdrawalCryptoAfterInsertHooks []WithdrawalCryptoHook
var withdrawalCryptoAfterSelectHooks []WithdrawalCryptoHook
var withdrawalCryptoAfterUpdateHooks []WithdrawalCryptoHook
var withdrawalCryptoAfterDeleteHooks []WithdrawalCryptoHook
var withdrawalCryptoAfterUpsertHooks []WithdrawalCryptoHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *WithdrawalCrypto) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalCryptoBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *WithdrawalCrypto) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalCryptoBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *WithdrawalCrypto) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalCryptoBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *WithdrawalCrypto) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalCryptoBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *WithdrawalCrypto) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalCryptoAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *WithdrawalCrypto) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalCryptoAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *WithdrawalCrypto) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalCryptoAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *WithdrawalCrypto) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalCryptoAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *WithdrawalCrypto) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalCryptoAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddWithdrawalCryptoHook registers your hook function for all future operations.
func AddWithdrawalCryptoHook(hookPoint boil.HookPoint, withdrawalCryptoHook WithdrawalCryptoHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		withdrawalCryptoBeforeInsertHooks = append(withdrawalCryptoBeforeInsertHooks, withdrawalCryptoHook)
	case boil.BeforeUpdateHook:
		withdrawalCryptoBeforeUpdateHooks = append(withdrawalCryptoBeforeUpdateHooks, withdrawalCryptoHook)
	case boil.BeforeDeleteHook:
		withdrawalCryptoBeforeDeleteHooks = append(withdrawalCryptoBeforeDeleteHooks, withdrawalCryptoHook)
	case boil.BeforeUpsertHook:
		withdrawalCryptoBeforeUpsertHooks = append(withdrawalCryptoBeforeUpsertHooks, withdrawalCryptoHook)
	case boil.AfterInsertHook:
		withdrawalCryptoAfterInsertHooks = append(withdrawalCryptoAfterInsertHooks, withdrawalCryptoHook)
	case boil.AfterSelectHook:
		withdrawalCryptoAfterSelectHooks = append(withdrawalCryptoAfterSelectHooks, withdrawalCryptoHook)
	case boil.AfterUpdateHook:
		withdrawalCryptoAfterUpdateHooks = append(withdrawalCryptoAfterUpdateHooks, withdrawalCryptoHook)
	case boil.AfterDeleteHook:
		withdrawalCryptoAfterDeleteHooks = append(withdrawalCryptoAfterDeleteHooks, withdrawalCryptoHook)
	case boil.AfterUpsertHook:
		withdrawalCryptoAfterUpsertHooks = append(withdrawalCryptoAfterUpsertHooks, withdrawalCryptoHook)
	}
}

// One returns a single withdrawalCrypto record from the query.
func (q withdrawalCryptoQuery) One(ctx context.Context, exec boil.ContextExecutor) (*WithdrawalCrypto, error) {
	o := &WithdrawalCrypto{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlite3: failed to execute a one query for withdrawal_crypto")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all WithdrawalCrypto records from the query.
func (q withdrawalCryptoQuery) All(ctx context.Context, exec boil.ContextExecutor) (WithdrawalCryptoSlice, error) {
	var o []*WithdrawalCrypto

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "sqlite3: failed to assign all query results to WithdrawalCrypto slice")
	}

	if len(withdrawalCryptoAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all WithdrawalCrypto records in the query.
func (q withdrawalCryptoQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to count withdrawal_crypto rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q withdrawalCryptoQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sqlite3: failed to check if withdrawal_crypto exists")
	}

	return count > 0, nil
}

// WithdrawalHistory pointed to by the foreign key.
func (o *WithdrawalCrypto) WithdrawalHistory(mods ...qm.QueryMod) withdrawalHistoryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.WithdrawalHistoryID),
	}

	queryMods = append(queryMods, mods...)

	query := WithdrawalHistories(queryMods...)
	queries.SetFrom(query.Query, "\"withdrawal_history\"")

	return query
}

// LoadWithdrawalHistory allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (withdrawalCryptoL) LoadWithdrawalHistory(ctx context.Context, e boil.ContextExecutor, singular bool, maybeWithdrawalCrypto interface{}, mods queries.Applicator) error {
	var slice []*WithdrawalCrypto
	var object *WithdrawalCrypto

	if singular {
		object = maybeWithdrawalCrypto.(*WithdrawalCrypto)
	} else {
		slice = *maybeWithdrawalCrypto.(*[]*WithdrawalCrypto)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &withdrawalCryptoR{}
		}
		args = append(args, object.WithdrawalHistoryID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &withdrawalCryptoR{}
			}

			for _, a := range args {
				if a == obj.WithdrawalHistoryID {
					continue Outer
				}
			}

			args = append(args, obj.WithdrawalHistoryID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`withdrawal_history`), qm.WhereIn(`withdrawal_history.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load WithdrawalHistory")
	}

	var resultSlice []*WithdrawalHistory
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice WithdrawalHistory")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for withdrawal_history")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for withdrawal_history")
	}

	if len(withdrawalCryptoAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.WithdrawalHistory = foreign
		if foreign.R == nil {
			foreign.R = &withdrawalHistoryR{}
		}
		foreign.R.WithdrawalCryptos = append(foreign.R.WithdrawalCryptos, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.WithdrawalHistoryID == foreign.ID {
				local.R.WithdrawalHistory = foreign
				if foreign.R == nil {
					foreign.R = &withdrawalHistoryR{}
				}
				foreign.R.WithdrawalCryptos = append(foreign.R.WithdrawalCryptos, local)
				break
			}
		}
	}

	return nil
}

// SetWithdrawalHistory of the withdrawalCrypto to the related item.
// Sets o.R.WithdrawalHistory to related.
// Adds o to related.R.WithdrawalCryptos.
func (o *WithdrawalCrypto) SetWithdrawalHistory(ctx context.Context, exec boil.ContextExecutor, insert bool, related *WithdrawalHistory) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"withdrawal_crypto\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"withdrawal_history_id"}),
		strmangle.WhereClause("\"", "\"", 0, withdrawalCryptoPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.WithdrawalHistoryID = related.ID
	if o.R == nil {
		o.R = &withdrawalCryptoR{
			WithdrawalHistory: related,
		}
	} else {
		o.R.WithdrawalHistory = related
	}

	if related.R == nil {
		related.R = &withdrawalHistoryR{
			WithdrawalCryptos: WithdrawalCryptoSlice{o},
		}
	} else {
		related.R.WithdrawalCryptos = append(related.R.WithdrawalCryptos, o)
	}

	return nil
}

// WithdrawalCryptos retrieves all the records using an executor.
func WithdrawalCryptos(mods ...qm.QueryMod) withdrawalCryptoQuery {
	mods = append(mods, qm.From("\"withdrawal_crypto\""))
	return withdrawalCryptoQuery{NewQuery(mods...)}
}

// FindWithdrawalCrypto retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindWithdrawalCrypto(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*WithdrawalCrypto, error) {
	withdrawalCryptoObj := &WithdrawalCrypto{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"withdrawal_crypto\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, withdrawalCryptoObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlite3: unable to select from withdrawal_crypto")
	}

	return withdrawalCryptoObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *WithdrawalCrypto) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("sqlite3: no withdrawal_crypto provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(withdrawalCryptoColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	withdrawalCryptoInsertCacheMut.RLock()
	cache, cached := withdrawalCryptoInsertCache[key]
	withdrawalCryptoInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			withdrawalCryptoAllColumns,
			withdrawalCryptoColumnsWithDefault,
			withdrawalCryptoColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(withdrawalCryptoType, withdrawalCryptoMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(withdrawalCryptoType, withdrawalCryptoMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"withdrawal_crypto\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"withdrawal_crypto\" () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"withdrawal_crypto\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, withdrawalCryptoPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "sqlite3: unable to insert into withdrawal_crypto")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == withdrawalCryptoMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "sqlite3: unable to populate default values for withdrawal_crypto")
	}

CacheNoHooks:
	if !cached {
		withdrawalCryptoInsertCacheMut.Lock()
		withdrawalCryptoInsertCache[key] = cache
		withdrawalCryptoInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the WithdrawalCrypto.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *WithdrawalCrypto) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	withdrawalCryptoUpdateCacheMut.RLock()
	cache, cached := withdrawalCryptoUpdateCache[key]
	withdrawalCryptoUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			withdrawalCryptoAllColumns,
			withdrawalCryptoPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("sqlite3: unable to update withdrawal_crypto, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"withdrawal_crypto\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, withdrawalCryptoPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(withdrawalCryptoType, withdrawalCryptoMapping, append(wl, withdrawalCryptoPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to update withdrawal_crypto row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by update for withdrawal_crypto")
	}

	if !cached {
		withdrawalCryptoUpdateCacheMut.Lock()
		withdrawalCryptoUpdateCache[key] = cache
		withdrawalCryptoUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q withdrawalCryptoQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to update all for withdrawal_crypto")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to retrieve rows affected for withdrawal_crypto")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o WithdrawalCryptoSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("sqlite3: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), withdrawalCryptoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"withdrawal_crypto\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, withdrawalCryptoPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to update all in withdrawalCrypto slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to retrieve rows affected all in update all withdrawalCrypto")
	}
	return rowsAff, nil
}

// Delete deletes a single WithdrawalCrypto record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *WithdrawalCrypto) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("sqlite3: no WithdrawalCrypto provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), withdrawalCryptoPrimaryKeyMapping)
	sql := "DELETE FROM \"withdrawal_crypto\" WHERE \"id\"=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to delete from withdrawal_crypto")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by delete for withdrawal_crypto")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q withdrawalCryptoQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("sqlite3: no withdrawalCryptoQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to delete all from withdrawal_crypto")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by deleteall for withdrawal_crypto")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o WithdrawalCryptoSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(withdrawalCryptoBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), withdrawalCryptoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"withdrawal_crypto\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, withdrawalCryptoPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to delete all from withdrawalCrypto slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by deleteall for withdrawal_crypto")
	}

	if len(withdrawalCryptoAfterDeleteHooks) != 0 {
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
func (o *WithdrawalCrypto) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindWithdrawalCrypto(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *WithdrawalCryptoSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := WithdrawalCryptoSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), withdrawalCryptoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"withdrawal_crypto\".* FROM \"withdrawal_crypto\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, withdrawalCryptoPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "sqlite3: unable to reload all in WithdrawalCryptoSlice")
	}

	*o = slice

	return nil
}

// WithdrawalCryptoExists checks if the WithdrawalCrypto row exists.
func WithdrawalCryptoExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"withdrawal_crypto\" where \"id\"=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sqlite3: unable to check if withdrawal_crypto exists")
	}

	return exists, nil
}
