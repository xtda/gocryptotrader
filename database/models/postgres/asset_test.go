// Code generated by SQLBoiler 3.5.0-gct (https://github.com/thrasher-corp/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package postgres

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
	"github.com/thrasher-corp/sqlboiler/randomize"
	"github.com/thrasher-corp/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testAssets(t *testing.T) {
	t.Parallel()

	query := Assets()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAssetsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Asset{}
	if err = randomize.Struct(seed, o, assetDBTypes, true, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Assets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAssetsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Asset{}
	if err = randomize.Struct(seed, o, assetDBTypes, true, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Assets().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Assets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAssetsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Asset{}
	if err = randomize.Struct(seed, o, assetDBTypes, true, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AssetSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Assets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAssetsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Asset{}
	if err = randomize.Struct(seed, o, assetDBTypes, true, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AssetExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Asset exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AssetExists to return true, but got false.")
	}
}

func testAssetsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Asset{}
	if err = randomize.Struct(seed, o, assetDBTypes, true, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	assetFound, err := FindAsset(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if assetFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAssetsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Asset{}
	if err = randomize.Struct(seed, o, assetDBTypes, true, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Assets().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAssetsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Asset{}
	if err = randomize.Struct(seed, o, assetDBTypes, true, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Assets().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAssetsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	assetOne := &Asset{}
	assetTwo := &Asset{}
	if err = randomize.Struct(seed, assetOne, assetDBTypes, false, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}
	if err = randomize.Struct(seed, assetTwo, assetDBTypes, false, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = assetOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = assetTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Assets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAssetsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	assetOne := &Asset{}
	assetTwo := &Asset{}
	if err = randomize.Struct(seed, assetOne, assetDBTypes, false, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}
	if err = randomize.Struct(seed, assetTwo, assetDBTypes, false, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = assetOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = assetTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Assets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func assetBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Asset) error {
	*o = Asset{}
	return nil
}

func assetAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Asset) error {
	*o = Asset{}
	return nil
}

func assetAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Asset) error {
	*o = Asset{}
	return nil
}

func assetBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Asset) error {
	*o = Asset{}
	return nil
}

func assetAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Asset) error {
	*o = Asset{}
	return nil
}

func assetBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Asset) error {
	*o = Asset{}
	return nil
}

func assetAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Asset) error {
	*o = Asset{}
	return nil
}

func assetBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Asset) error {
	*o = Asset{}
	return nil
}

func assetAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Asset) error {
	*o = Asset{}
	return nil
}

func testAssetsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Asset{}
	o := &Asset{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, assetDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Asset object: %s", err)
	}

	AddAssetHook(boil.BeforeInsertHook, assetBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	assetBeforeInsertHooks = []AssetHook{}

	AddAssetHook(boil.AfterInsertHook, assetAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	assetAfterInsertHooks = []AssetHook{}

	AddAssetHook(boil.AfterSelectHook, assetAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	assetAfterSelectHooks = []AssetHook{}

	AddAssetHook(boil.BeforeUpdateHook, assetBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	assetBeforeUpdateHooks = []AssetHook{}

	AddAssetHook(boil.AfterUpdateHook, assetAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	assetAfterUpdateHooks = []AssetHook{}

	AddAssetHook(boil.BeforeDeleteHook, assetBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	assetBeforeDeleteHooks = []AssetHook{}

	AddAssetHook(boil.AfterDeleteHook, assetAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	assetAfterDeleteHooks = []AssetHook{}

	AddAssetHook(boil.BeforeUpsertHook, assetBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	assetBeforeUpsertHooks = []AssetHook{}

	AddAssetHook(boil.AfterUpsertHook, assetAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	assetAfterUpsertHooks = []AssetHook{}
}

func testAssetsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Asset{}
	if err = randomize.Struct(seed, o, assetDBTypes, true, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Assets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAssetsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Asset{}
	if err = randomize.Struct(seed, o, assetDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(assetColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Assets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAssetToOneExchangeUsingExchange(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Asset
	var foreign Exchange

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, assetDBTypes, true, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, exchangeDBTypes, false, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ExchangeID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Exchange().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AssetSlice{&local}
	if err = local.L.LoadExchange(ctx, tx, false, (*[]*Asset)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Exchange == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Exchange = nil
	if err = local.L.LoadExchange(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Exchange == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAssetToOneSetOpExchangeUsingExchange(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Asset
	var b, c Exchange

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, assetDBTypes, false, strmangle.SetComplement(assetPrimaryKeyColumns, assetColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, exchangeDBTypes, false, strmangle.SetComplement(exchangePrimaryKeyColumns, exchangeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, exchangeDBTypes, false, strmangle.SetComplement(exchangePrimaryKeyColumns, exchangeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Exchange{&b, &c} {
		err = a.SetExchange(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Exchange != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Assets[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ExchangeID, x.ID) {
			t.Error("foreign key was wrong value", a.ExchangeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ExchangeID))
		reflect.Indirect(reflect.ValueOf(&a.ExchangeID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ExchangeID, x.ID) {
			t.Error("foreign key was wrong value", a.ExchangeID, x.ID)
		}
	}
}

func testAssetToOneRemoveOpExchangeUsingExchange(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Asset
	var b Exchange

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, assetDBTypes, false, strmangle.SetComplement(assetPrimaryKeyColumns, assetColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, exchangeDBTypes, false, strmangle.SetComplement(exchangePrimaryKeyColumns, exchangeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetExchange(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveExchange(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Exchange().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Exchange != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ExchangeID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Assets) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testAssetsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Asset{}
	if err = randomize.Struct(seed, o, assetDBTypes, true, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAssetsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Asset{}
	if err = randomize.Struct(seed, o, assetDBTypes, true, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AssetSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAssetsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Asset{}
	if err = randomize.Struct(seed, o, assetDBTypes, true, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Assets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	assetDBTypes = map[string]string{`ID`: `bigint`, `Name`: `character varying`, `ShortName`: `character varying`, `ExchangeID`: `uuid`, `Delimiter`: `character varying`}
	_            = bytes.MinRead
)

func testAssetsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(assetPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(assetAllColumns) == len(assetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Asset{}
	if err = randomize.Struct(seed, o, assetDBTypes, true, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Assets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, assetDBTypes, true, assetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAssetsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(assetAllColumns) == len(assetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Asset{}
	if err = randomize.Struct(seed, o, assetDBTypes, true, assetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Assets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, assetDBTypes, true, assetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(assetAllColumns, assetPrimaryKeyColumns) {
		fields = assetAllColumns
	} else {
		fields = strmangle.SetComplement(
			assetAllColumns,
			assetPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := AssetSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAssetsUpsert(t *testing.T) {
	t.Parallel()

	if len(assetAllColumns) == len(assetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Asset{}
	if err = randomize.Struct(seed, &o, assetDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Asset: %s", err)
	}

	count, err := Assets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, assetDBTypes, false, assetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Asset struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Asset: %s", err)
	}

	count, err = Assets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
