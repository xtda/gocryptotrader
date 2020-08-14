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

// Candle is an object representing the database table.
type Candle struct {
	ID         string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ExchangeID null.String `boil:"exchange_id" json:"exchange_id,omitempty" toml:"exchange_id" yaml:"exchange_id,omitempty"`
	Base       string      `boil:"Base" json:"Base" toml:"Base" yaml:"Base"`
	Quote      string      `boil:"Quote" json:"Quote" toml:"Quote" yaml:"Quote"`
	Interval   string      `boil:"Interval" json:"Interval" toml:"Interval" yaml:"Interval"`
	Timestamp  string      `boil:"Timestamp" json:"Timestamp" toml:"Timestamp" yaml:"Timestamp"`
	Open       float64     `boil:"Open" json:"Open" toml:"Open" yaml:"Open"`
	High       float64     `boil:"High" json:"High" toml:"High" yaml:"High"`
	Low        float64     `boil:"Low" json:"Low" toml:"Low" yaml:"Low"`
	Close      float64     `boil:"Close" json:"Close" toml:"Close" yaml:"Close"`
	Volume     float64     `boil:"Volume" json:"Volume" toml:"Volume" yaml:"Volume"`
	Asset      string      `boil:"Asset" json:"Asset" toml:"Asset" yaml:"Asset"`

	R *candleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L candleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CandleColumns = struct {
	ID         string
	ExchangeID string
	Base       string
	Quote      string
	Interval   string
	Timestamp  string
	Open       string
	High       string
	Low        string
	Close      string
	Volume     string
	Asset      string
}{
	ID:         "id",
	ExchangeID: "exchange_id",
	Base:       "Base",
	Quote:      "Quote",
	Interval:   "Interval",
	Timestamp:  "Timestamp",
	Open:       "Open",
	High:       "High",
	Low:        "Low",
	Close:      "Close",
	Volume:     "Volume",
	Asset:      "Asset",
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

type whereHelperfloat64 struct{ field string }

func (w whereHelperfloat64) EQ(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat64) NEQ(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat64) LT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat64) LTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat64) GT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat64) GTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var CandleWhere = struct {
	ID         whereHelperstring
	ExchangeID whereHelpernull_String
	Base       whereHelperstring
	Quote      whereHelperstring
	Interval   whereHelperstring
	Timestamp  whereHelperstring
	Open       whereHelperfloat64
	High       whereHelperfloat64
	Low        whereHelperfloat64
	Close      whereHelperfloat64
	Volume     whereHelperfloat64
	Asset      whereHelperstring
}{
	ID:         whereHelperstring{field: "\"candle\".\"id\""},
	ExchangeID: whereHelpernull_String{field: "\"candle\".\"exchange_id\""},
	Base:       whereHelperstring{field: "\"candle\".\"Base\""},
	Quote:      whereHelperstring{field: "\"candle\".\"Quote\""},
	Interval:   whereHelperstring{field: "\"candle\".\"Interval\""},
	Timestamp:  whereHelperstring{field: "\"candle\".\"Timestamp\""},
	Open:       whereHelperfloat64{field: "\"candle\".\"Open\""},
	High:       whereHelperfloat64{field: "\"candle\".\"High\""},
	Low:        whereHelperfloat64{field: "\"candle\".\"Low\""},
	Close:      whereHelperfloat64{field: "\"candle\".\"Close\""},
	Volume:     whereHelperfloat64{field: "\"candle\".\"Volume\""},
	Asset:      whereHelperstring{field: "\"candle\".\"Asset\""},
}

// CandleRels is where relationship names are stored.
var CandleRels = struct {
	Exchange string
}{
	Exchange: "Exchange",
}

// candleR is where relationships are stored.
type candleR struct {
	Exchange *Exchange
}

// NewStruct creates a new relationship struct
func (*candleR) NewStruct() *candleR {
	return &candleR{}
}

// candleL is where Load methods for each relationship are stored.
type candleL struct{}

var (
	candleAllColumns            = []string{"id", "exchange_id", "Base", "Quote", "Interval", "Timestamp", "Open", "High", "Low", "Close", "Volume", "Asset"}
	candleColumnsWithoutDefault = []string{"id", "exchange_id", "Base", "Quote", "Interval", "Timestamp", "Open", "High", "Low", "Close", "Volume", "Asset"}
	candleColumnsWithDefault    = []string{}
	candlePrimaryKeyColumns     = []string{"id"}
)

type (
	// CandleSlice is an alias for a slice of pointers to Candle.
	// This should generally be used opposed to []Candle.
	CandleSlice []*Candle
	// CandleHook is the signature for custom Candle hook methods
	CandleHook func(context.Context, boil.ContextExecutor, *Candle) error

	candleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	candleType                 = reflect.TypeOf(&Candle{})
	candleMapping              = queries.MakeStructMapping(candleType)
	candlePrimaryKeyMapping, _ = queries.BindMapping(candleType, candleMapping, candlePrimaryKeyColumns)
	candleInsertCacheMut       sync.RWMutex
	candleInsertCache          = make(map[string]insertCache)
	candleUpdateCacheMut       sync.RWMutex
	candleUpdateCache          = make(map[string]updateCache)
	candleUpsertCacheMut       sync.RWMutex
	candleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var candleBeforeInsertHooks []CandleHook
var candleBeforeUpdateHooks []CandleHook
var candleBeforeDeleteHooks []CandleHook
var candleBeforeUpsertHooks []CandleHook

var candleAfterInsertHooks []CandleHook
var candleAfterSelectHooks []CandleHook
var candleAfterUpdateHooks []CandleHook
var candleAfterDeleteHooks []CandleHook
var candleAfterUpsertHooks []CandleHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Candle) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range candleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Candle) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range candleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Candle) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range candleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Candle) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range candleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Candle) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range candleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Candle) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range candleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Candle) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range candleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Candle) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range candleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Candle) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range candleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCandleHook registers your hook function for all future operations.
func AddCandleHook(hookPoint boil.HookPoint, candleHook CandleHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		candleBeforeInsertHooks = append(candleBeforeInsertHooks, candleHook)
	case boil.BeforeUpdateHook:
		candleBeforeUpdateHooks = append(candleBeforeUpdateHooks, candleHook)
	case boil.BeforeDeleteHook:
		candleBeforeDeleteHooks = append(candleBeforeDeleteHooks, candleHook)
	case boil.BeforeUpsertHook:
		candleBeforeUpsertHooks = append(candleBeforeUpsertHooks, candleHook)
	case boil.AfterInsertHook:
		candleAfterInsertHooks = append(candleAfterInsertHooks, candleHook)
	case boil.AfterSelectHook:
		candleAfterSelectHooks = append(candleAfterSelectHooks, candleHook)
	case boil.AfterUpdateHook:
		candleAfterUpdateHooks = append(candleAfterUpdateHooks, candleHook)
	case boil.AfterDeleteHook:
		candleAfterDeleteHooks = append(candleAfterDeleteHooks, candleHook)
	case boil.AfterUpsertHook:
		candleAfterUpsertHooks = append(candleAfterUpsertHooks, candleHook)
	}
}

// One returns a single candle record from the query.
func (q candleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Candle, error) {
	o := &Candle{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlite3: failed to execute a one query for candle")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Candle records from the query.
func (q candleQuery) All(ctx context.Context, exec boil.ContextExecutor) (CandleSlice, error) {
	var o []*Candle

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "sqlite3: failed to assign all query results to Candle slice")
	}

	if len(candleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Candle records in the query.
func (q candleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to count candle rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q candleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sqlite3: failed to check if candle exists")
	}

	return count > 0, nil
}

// Exchange pointed to by the foreign key.
func (o *Candle) Exchange(mods ...qm.QueryMod) exchangeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ExchangeID),
	}

	queryMods = append(queryMods, mods...)

	query := Exchanges(queryMods...)
	queries.SetFrom(query.Query, "\"exchange\"")

	return query
}

// LoadExchange allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (candleL) LoadExchange(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCandle interface{}, mods queries.Applicator) error {
	var slice []*Candle
	var object *Candle

	if singular {
		object = maybeCandle.(*Candle)
	} else {
		slice = *maybeCandle.(*[]*Candle)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &candleR{}
		}
		if !queries.IsNil(object.ExchangeID) {
			args = append(args, object.ExchangeID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &candleR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ExchangeID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ExchangeID) {
				args = append(args, obj.ExchangeID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`exchange`), qm.WhereIn(`exchange.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Exchange")
	}

	var resultSlice []*Exchange
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Exchange")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for exchange")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for exchange")
	}

	if len(candleAfterSelectHooks) != 0 {
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
		object.R.Exchange = foreign
		if foreign.R == nil {
			foreign.R = &exchangeR{}
		}
		foreign.R.Candle = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ExchangeID, foreign.ID) {
				local.R.Exchange = foreign
				if foreign.R == nil {
					foreign.R = &exchangeR{}
				}
				foreign.R.Candle = local
				break
			}
		}
	}

	return nil
}

// SetExchange of the candle to the related item.
// Sets o.R.Exchange to related.
// Adds o to related.R.Candle.
func (o *Candle) SetExchange(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Exchange) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"candle\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"exchange_id"}),
		strmangle.WhereClause("\"", "\"", 0, candlePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.ExchangeID, related.ID)
	if o.R == nil {
		o.R = &candleR{
			Exchange: related,
		}
	} else {
		o.R.Exchange = related
	}

	if related.R == nil {
		related.R = &exchangeR{
			Candle: o,
		}
	} else {
		related.R.Candle = o
	}

	return nil
}

// RemoveExchange relationship.
// Sets o.R.Exchange to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Candle) RemoveExchange(ctx context.Context, exec boil.ContextExecutor, related *Exchange) error {
	var err error

	queries.SetScanner(&o.ExchangeID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("exchange_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Exchange = nil
	if related == nil || related.R == nil {
		return nil
	}

	related.R.Candle = nil
	return nil
}

// Candles retrieves all the records using an executor.
func Candles(mods ...qm.QueryMod) candleQuery {
	mods = append(mods, qm.From("\"candle\""))
	return candleQuery{NewQuery(mods...)}
}

// FindCandle retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCandle(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Candle, error) {
	candleObj := &Candle{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"candle\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, candleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlite3: unable to select from candle")
	}

	return candleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Candle) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("sqlite3: no candle provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(candleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	candleInsertCacheMut.RLock()
	cache, cached := candleInsertCache[key]
	candleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			candleAllColumns,
			candleColumnsWithDefault,
			candleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(candleType, candleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(candleType, candleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"candle\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"candle\" () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"candle\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, candlePrimaryKeyColumns))
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
		return errors.Wrap(err, "sqlite3: unable to insert into candle")
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
		return errors.Wrap(err, "sqlite3: unable to populate default values for candle")
	}

CacheNoHooks:
	if !cached {
		candleInsertCacheMut.Lock()
		candleInsertCache[key] = cache
		candleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Candle.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Candle) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	candleUpdateCacheMut.RLock()
	cache, cached := candleUpdateCache[key]
	candleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			candleAllColumns,
			candlePrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("sqlite3: unable to update candle, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"candle\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, candlePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(candleType, candleMapping, append(wl, candlePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "sqlite3: unable to update candle row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by update for candle")
	}

	if !cached {
		candleUpdateCacheMut.Lock()
		candleUpdateCache[key] = cache
		candleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q candleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to update all for candle")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to retrieve rows affected for candle")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CandleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), candlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"candle\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, candlePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to update all in candle slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to retrieve rows affected all in update all candle")
	}
	return rowsAff, nil
}

// Delete deletes a single Candle record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Candle) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("sqlite3: no Candle provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), candlePrimaryKeyMapping)
	sql := "DELETE FROM \"candle\" WHERE \"id\"=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to delete from candle")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by delete for candle")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q candleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("sqlite3: no candleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to delete all from candle")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by deleteall for candle")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CandleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(candleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), candlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"candle\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, candlePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to delete all from candle slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by deleteall for candle")
	}

	if len(candleAfterDeleteHooks) != 0 {
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
func (o *Candle) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCandle(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CandleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CandleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), candlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"candle\".* FROM \"candle\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, candlePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "sqlite3: unable to reload all in CandleSlice")
	}

	*o = slice

	return nil
}

// CandleExists checks if the Candle row exists.
func CandleExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"candle\" where \"id\"=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sqlite3: unable to check if candle exists")
	}

	return exists, nil
}
