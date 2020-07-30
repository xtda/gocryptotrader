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
)

// Exchange is an object representing the database table.
type Exchange struct {
	ID   string `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *exchangeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L exchangeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ExchangeColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}

// Generated where

var ExchangeWhere = struct {
	ID   whereHelperstring
	Name whereHelperstring
}{
	ID:   whereHelperstring{field: "\"exchange\".\"id\""},
	Name: whereHelperstring{field: "\"exchange\".\"name\""},
}

// ExchangeRels is where relationship names are stored.
var ExchangeRels = struct {
	Candle                          string
	ExchangeNameWithdrawalHistories string
}{
	Candle:                          "Candle",
	ExchangeNameWithdrawalHistories: "ExchangeNameWithdrawalHistories",
}

// exchangeR is where relationships are stored.
type exchangeR struct {
	Candle                          *Candle
	ExchangeNameWithdrawalHistories WithdrawalHistorySlice
}

// NewStruct creates a new relationship struct
func (*exchangeR) NewStruct() *exchangeR {
	return &exchangeR{}
}

// exchangeL is where Load methods for each relationship are stored.
type exchangeL struct{}

var (
	exchangeAllColumns            = []string{"id", "name"}
	exchangeColumnsWithoutDefault = []string{"id", "name"}
	exchangeColumnsWithDefault    = []string{}
	exchangePrimaryKeyColumns     = []string{"id"}
)

type (
	// ExchangeSlice is an alias for a slice of pointers to Exchange.
	// This should generally be used opposed to []Exchange.
	ExchangeSlice []*Exchange
	// ExchangeHook is the signature for custom Exchange hook methods
	ExchangeHook func(context.Context, boil.ContextExecutor, *Exchange) error

	exchangeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	exchangeType                 = reflect.TypeOf(&Exchange{})
	exchangeMapping              = queries.MakeStructMapping(exchangeType)
	exchangePrimaryKeyMapping, _ = queries.BindMapping(exchangeType, exchangeMapping, exchangePrimaryKeyColumns)
	exchangeInsertCacheMut       sync.RWMutex
	exchangeInsertCache          = make(map[string]insertCache)
	exchangeUpdateCacheMut       sync.RWMutex
	exchangeUpdateCache          = make(map[string]updateCache)
	exchangeUpsertCacheMut       sync.RWMutex
	exchangeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var exchangeBeforeInsertHooks []ExchangeHook
var exchangeBeforeUpdateHooks []ExchangeHook
var exchangeBeforeDeleteHooks []ExchangeHook
var exchangeBeforeUpsertHooks []ExchangeHook

var exchangeAfterInsertHooks []ExchangeHook
var exchangeAfterSelectHooks []ExchangeHook
var exchangeAfterUpdateHooks []ExchangeHook
var exchangeAfterDeleteHooks []ExchangeHook
var exchangeAfterUpsertHooks []ExchangeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Exchange) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exchangeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Exchange) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exchangeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Exchange) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exchangeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Exchange) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exchangeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Exchange) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exchangeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Exchange) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exchangeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Exchange) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exchangeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Exchange) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exchangeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Exchange) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exchangeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddExchangeHook registers your hook function for all future operations.
func AddExchangeHook(hookPoint boil.HookPoint, exchangeHook ExchangeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		exchangeBeforeInsertHooks = append(exchangeBeforeInsertHooks, exchangeHook)
	case boil.BeforeUpdateHook:
		exchangeBeforeUpdateHooks = append(exchangeBeforeUpdateHooks, exchangeHook)
	case boil.BeforeDeleteHook:
		exchangeBeforeDeleteHooks = append(exchangeBeforeDeleteHooks, exchangeHook)
	case boil.BeforeUpsertHook:
		exchangeBeforeUpsertHooks = append(exchangeBeforeUpsertHooks, exchangeHook)
	case boil.AfterInsertHook:
		exchangeAfterInsertHooks = append(exchangeAfterInsertHooks, exchangeHook)
	case boil.AfterSelectHook:
		exchangeAfterSelectHooks = append(exchangeAfterSelectHooks, exchangeHook)
	case boil.AfterUpdateHook:
		exchangeAfterUpdateHooks = append(exchangeAfterUpdateHooks, exchangeHook)
	case boil.AfterDeleteHook:
		exchangeAfterDeleteHooks = append(exchangeAfterDeleteHooks, exchangeHook)
	case boil.AfterUpsertHook:
		exchangeAfterUpsertHooks = append(exchangeAfterUpsertHooks, exchangeHook)
	}
}

// One returns a single exchange record from the query.
func (q exchangeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Exchange, error) {
	o := &Exchange{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlite3: failed to execute a one query for exchange")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Exchange records from the query.
func (q exchangeQuery) All(ctx context.Context, exec boil.ContextExecutor) (ExchangeSlice, error) {
	var o []*Exchange

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "sqlite3: failed to assign all query results to Exchange slice")
	}

	if len(exchangeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Exchange records in the query.
func (q exchangeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to count exchange rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q exchangeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sqlite3: failed to check if exchange exists")
	}

	return count > 0, nil
}

// Candle pointed to by the foreign key.
func (o *Exchange) Candle(mods ...qm.QueryMod) candleQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"exchange_id\" = ?", o.ID),
	}

	queryMods = append(queryMods, mods...)

	query := Candles(queryMods...)
	queries.SetFrom(query.Query, "\"candle\"")

	return query
}

// ExchangeNameWithdrawalHistories retrieves all the withdrawal_history's WithdrawalHistories with an executor via exchange_name_id column.
func (o *Exchange) ExchangeNameWithdrawalHistories(mods ...qm.QueryMod) withdrawalHistoryQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"withdrawal_history\".\"exchange_name_id\"=?", o.ID),
	)

	query := WithdrawalHistories(queryMods...)
	queries.SetFrom(query.Query, "\"withdrawal_history\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"withdrawal_history\".*"})
	}

	return query
}

// LoadCandle allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-1 relationship.
func (exchangeL) LoadCandle(ctx context.Context, e boil.ContextExecutor, singular bool, maybeExchange interface{}, mods queries.Applicator) error {
	var slice []*Exchange
	var object *Exchange

	if singular {
		object = maybeExchange.(*Exchange)
	} else {
		slice = *maybeExchange.(*[]*Exchange)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &exchangeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &exchangeR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`candle`), qm.WhereIn(`candle.exchange_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Candle")
	}

	var resultSlice []*Candle
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Candle")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for candle")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for candle")
	}

	if len(exchangeAfterSelectHooks) != 0 {
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
		object.R.Candle = foreign
		if foreign.R == nil {
			foreign.R = &candleR{}
		}
		foreign.R.Exchange = object
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ID, foreign.ExchangeID) {
				local.R.Candle = foreign
				if foreign.R == nil {
					foreign.R = &candleR{}
				}
				foreign.R.Exchange = local
				break
			}
		}
	}

	return nil
}

// LoadExchangeNameWithdrawalHistories allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (exchangeL) LoadExchangeNameWithdrawalHistories(ctx context.Context, e boil.ContextExecutor, singular bool, maybeExchange interface{}, mods queries.Applicator) error {
	var slice []*Exchange
	var object *Exchange

	if singular {
		object = maybeExchange.(*Exchange)
	} else {
		slice = *maybeExchange.(*[]*Exchange)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &exchangeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &exchangeR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`withdrawal_history`), qm.WhereIn(`withdrawal_history.exchange_name_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load withdrawal_history")
	}

	var resultSlice []*WithdrawalHistory
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice withdrawal_history")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on withdrawal_history")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for withdrawal_history")
	}

	if len(withdrawalHistoryAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ExchangeNameWithdrawalHistories = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &withdrawalHistoryR{}
			}
			foreign.R.ExchangeName = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.ExchangeNameID) {
				local.R.ExchangeNameWithdrawalHistories = append(local.R.ExchangeNameWithdrawalHistories, foreign)
				if foreign.R == nil {
					foreign.R = &withdrawalHistoryR{}
				}
				foreign.R.ExchangeName = local
				break
			}
		}
	}

	return nil
}

// SetCandle of the exchange to the related item.
// Sets o.R.Candle to related.
// Adds o to related.R.Exchange.
func (o *Exchange) SetCandle(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Candle) error {
	var err error

	if insert {
		queries.Assign(&related.ExchangeID, o.ID)

		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"candle\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, []string{"exchange_id"}),
			strmangle.WhereClause("\"", "\"", 0, candlePrimaryKeyColumns),
		)
		values := []interface{}{o.ID, related.ID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		queries.Assign(&related.ExchangeID, o.ID)
	}

	if o.R == nil {
		o.R = &exchangeR{
			Candle: related,
		}
	} else {
		o.R.Candle = related
	}

	if related.R == nil {
		related.R = &candleR{
			Exchange: o,
		}
	} else {
		related.R.Exchange = o
	}
	return nil
}

// RemoveCandle relationship.
// Sets o.R.Candle to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Exchange) RemoveCandle(ctx context.Context, exec boil.ContextExecutor, related *Candle) error {
	var err error

	queries.SetScanner(&related.ExchangeID, nil)
	if _, err = related.Update(ctx, exec, boil.Whitelist("exchange_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Candle = nil
	if related == nil || related.R == nil {
		return nil
	}

	related.R.Exchange = nil
	return nil
}

// AddExchangeNameWithdrawalHistories adds the given related objects to the existing relationships
// of the exchange, optionally inserting them as new records.
// Appends related to o.R.ExchangeNameWithdrawalHistories.
// Sets related.R.ExchangeName appropriately.
func (o *Exchange) AddExchangeNameWithdrawalHistories(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*WithdrawalHistory) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.ExchangeNameID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"withdrawal_history\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"exchange_name_id"}),
				strmangle.WhereClause("\"", "\"", 0, withdrawalHistoryPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.ExchangeNameID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &exchangeR{
			ExchangeNameWithdrawalHistories: related,
		}
	} else {
		o.R.ExchangeNameWithdrawalHistories = append(o.R.ExchangeNameWithdrawalHistories, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &withdrawalHistoryR{
				ExchangeName: o,
			}
		} else {
			rel.R.ExchangeName = o
		}
	}
	return nil
}

// SetExchangeNameWithdrawalHistories removes all previously related items of the
// exchange replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.ExchangeName's ExchangeNameWithdrawalHistories accordingly.
// Replaces o.R.ExchangeNameWithdrawalHistories with related.
// Sets related.R.ExchangeName's ExchangeNameWithdrawalHistories accordingly.
func (o *Exchange) SetExchangeNameWithdrawalHistories(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*WithdrawalHistory) error {
	query := "update \"withdrawal_history\" set \"exchange_name_id\" = null where \"exchange_name_id\" = ?"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.ExchangeNameWithdrawalHistories {
			queries.SetScanner(&rel.ExchangeNameID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.ExchangeName = nil
		}

		o.R.ExchangeNameWithdrawalHistories = nil
	}
	return o.AddExchangeNameWithdrawalHistories(ctx, exec, insert, related...)
}

// RemoveExchangeNameWithdrawalHistories relationships from objects passed in.
// Removes related items from R.ExchangeNameWithdrawalHistories (uses pointer comparison, removal does not keep order)
// Sets related.R.ExchangeName.
func (o *Exchange) RemoveExchangeNameWithdrawalHistories(ctx context.Context, exec boil.ContextExecutor, related ...*WithdrawalHistory) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.ExchangeNameID, nil)
		if rel.R != nil {
			rel.R.ExchangeName = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("exchange_name_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.ExchangeNameWithdrawalHistories {
			if rel != ri {
				continue
			}

			ln := len(o.R.ExchangeNameWithdrawalHistories)
			if ln > 1 && i < ln-1 {
				o.R.ExchangeNameWithdrawalHistories[i] = o.R.ExchangeNameWithdrawalHistories[ln-1]
			}
			o.R.ExchangeNameWithdrawalHistories = o.R.ExchangeNameWithdrawalHistories[:ln-1]
			break
		}
	}

	return nil
}

// Exchanges retrieves all the records using an executor.
func Exchanges(mods ...qm.QueryMod) exchangeQuery {
	mods = append(mods, qm.From("\"exchange\""))
	return exchangeQuery{NewQuery(mods...)}
}

// FindExchange retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindExchange(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Exchange, error) {
	exchangeObj := &Exchange{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"exchange\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, exchangeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlite3: unable to select from exchange")
	}

	return exchangeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Exchange) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("sqlite3: no exchange provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(exchangeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	exchangeInsertCacheMut.RLock()
	cache, cached := exchangeInsertCache[key]
	exchangeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			exchangeAllColumns,
			exchangeColumnsWithDefault,
			exchangeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(exchangeType, exchangeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(exchangeType, exchangeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"exchange\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"exchange\" () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"exchange\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, exchangePrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "sqlite3: unable to insert into exchange")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
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
		return errors.Wrap(err, "sqlite3: unable to populate default values for exchange")
	}

CacheNoHooks:
	if !cached {
		exchangeInsertCacheMut.Lock()
		exchangeInsertCache[key] = cache
		exchangeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Exchange.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Exchange) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	exchangeUpdateCacheMut.RLock()
	cache, cached := exchangeUpdateCache[key]
	exchangeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			exchangeAllColumns,
			exchangePrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("sqlite3: unable to update exchange, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"exchange\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, exchangePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(exchangeType, exchangeMapping, append(wl, exchangePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "sqlite3: unable to update exchange row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by update for exchange")
	}

	if !cached {
		exchangeUpdateCacheMut.Lock()
		exchangeUpdateCache[key] = cache
		exchangeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q exchangeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to update all for exchange")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to retrieve rows affected for exchange")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ExchangeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), exchangePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"exchange\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, exchangePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to update all in exchange slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to retrieve rows affected all in update all exchange")
	}
	return rowsAff, nil
}

// Delete deletes a single Exchange record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Exchange) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("sqlite3: no Exchange provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), exchangePrimaryKeyMapping)
	sql := "DELETE FROM \"exchange\" WHERE \"id\"=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to delete from exchange")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by delete for exchange")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q exchangeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("sqlite3: no exchangeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to delete all from exchange")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by deleteall for exchange")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ExchangeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(exchangeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), exchangePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"exchange\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, exchangePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to delete all from exchange slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by deleteall for exchange")
	}

	if len(exchangeAfterDeleteHooks) != 0 {
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
func (o *Exchange) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindExchange(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ExchangeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ExchangeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), exchangePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"exchange\".* FROM \"exchange\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, exchangePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "sqlite3: unable to reload all in ExchangeSlice")
	}

	*o = slice

	return nil
}

// ExchangeExists checks if the Exchange row exists.
func ExchangeExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"exchange\" where \"id\"=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sqlite3: unable to check if exchange exists")
	}

	return exists, nil
}
