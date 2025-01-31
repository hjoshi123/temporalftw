// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// TransactionStatus is an object representing the database table.
type TransactionStatus struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *transactionStatusR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L transactionStatusL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TransactionStatusColumns = struct {
	ID        string
	Name      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var TransactionStatusTableColumns = struct {
	ID        string
	Name      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "transaction_statuses.id",
	Name:      "transaction_statuses.name",
	CreatedAt: "transaction_statuses.created_at",
	UpdatedAt: "transaction_statuses.updated_at",
}

// Generated where

var TransactionStatusWhere = struct {
	ID        whereHelperint
	Name      whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint{field: "\"transaction_statuses\".\"id\""},
	Name:      whereHelperstring{field: "\"transaction_statuses\".\"name\""},
	CreatedAt: whereHelpertime_Time{field: "\"transaction_statuses\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"transaction_statuses\".\"updated_at\""},
}

// TransactionStatusRels is where relationship names are stored.
var TransactionStatusRels = struct {
	Transactions string
}{
	Transactions: "Transactions",
}

// transactionStatusR is where relationships are stored.
type transactionStatusR struct {
	Transactions TransactionSlice `boil:"Transactions" json:"Transactions" toml:"Transactions" yaml:"Transactions"`
}

// NewStruct creates a new relationship struct
func (*transactionStatusR) NewStruct() *transactionStatusR {
	return &transactionStatusR{}
}

func (r *transactionStatusR) GetTransactions() TransactionSlice {
	if r == nil {
		return nil
	}
	return r.Transactions
}

// transactionStatusL is where Load methods for each relationship are stored.
type transactionStatusL struct{}

var (
	transactionStatusAllColumns            = []string{"id", "name", "created_at", "updated_at"}
	transactionStatusColumnsWithoutDefault = []string{"name"}
	transactionStatusColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	transactionStatusPrimaryKeyColumns     = []string{"id"}
	transactionStatusGeneratedColumns      = []string{}
)

type (
	// TransactionStatusSlice is an alias for a slice of pointers to TransactionStatus.
	// This should almost always be used instead of []TransactionStatus.
	TransactionStatusSlice []*TransactionStatus
	// TransactionStatusHook is the signature for custom TransactionStatus hook methods
	TransactionStatusHook func(context.Context, boil.ContextExecutor, *TransactionStatus) error

	transactionStatusQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	transactionStatusType                 = reflect.TypeOf(&TransactionStatus{})
	transactionStatusMapping              = queries.MakeStructMapping(transactionStatusType)
	transactionStatusPrimaryKeyMapping, _ = queries.BindMapping(transactionStatusType, transactionStatusMapping, transactionStatusPrimaryKeyColumns)
	transactionStatusInsertCacheMut       sync.RWMutex
	transactionStatusInsertCache          = make(map[string]insertCache)
	transactionStatusUpdateCacheMut       sync.RWMutex
	transactionStatusUpdateCache          = make(map[string]updateCache)
	transactionStatusUpsertCacheMut       sync.RWMutex
	transactionStatusUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var transactionStatusAfterSelectMu sync.Mutex
var transactionStatusAfterSelectHooks []TransactionStatusHook

var transactionStatusBeforeInsertMu sync.Mutex
var transactionStatusBeforeInsertHooks []TransactionStatusHook
var transactionStatusAfterInsertMu sync.Mutex
var transactionStatusAfterInsertHooks []TransactionStatusHook

var transactionStatusBeforeUpdateMu sync.Mutex
var transactionStatusBeforeUpdateHooks []TransactionStatusHook
var transactionStatusAfterUpdateMu sync.Mutex
var transactionStatusAfterUpdateHooks []TransactionStatusHook

var transactionStatusBeforeDeleteMu sync.Mutex
var transactionStatusBeforeDeleteHooks []TransactionStatusHook
var transactionStatusAfterDeleteMu sync.Mutex
var transactionStatusAfterDeleteHooks []TransactionStatusHook

var transactionStatusBeforeUpsertMu sync.Mutex
var transactionStatusBeforeUpsertHooks []TransactionStatusHook
var transactionStatusAfterUpsertMu sync.Mutex
var transactionStatusAfterUpsertHooks []TransactionStatusHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TransactionStatus) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionStatusAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TransactionStatus) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionStatusBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TransactionStatus) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionStatusAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TransactionStatus) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionStatusBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TransactionStatus) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionStatusAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TransactionStatus) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionStatusBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TransactionStatus) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionStatusAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TransactionStatus) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionStatusBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TransactionStatus) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionStatusAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTransactionStatusHook registers your hook function for all future operations.
func AddTransactionStatusHook(hookPoint boil.HookPoint, transactionStatusHook TransactionStatusHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		transactionStatusAfterSelectMu.Lock()
		transactionStatusAfterSelectHooks = append(transactionStatusAfterSelectHooks, transactionStatusHook)
		transactionStatusAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		transactionStatusBeforeInsertMu.Lock()
		transactionStatusBeforeInsertHooks = append(transactionStatusBeforeInsertHooks, transactionStatusHook)
		transactionStatusBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		transactionStatusAfterInsertMu.Lock()
		transactionStatusAfterInsertHooks = append(transactionStatusAfterInsertHooks, transactionStatusHook)
		transactionStatusAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		transactionStatusBeforeUpdateMu.Lock()
		transactionStatusBeforeUpdateHooks = append(transactionStatusBeforeUpdateHooks, transactionStatusHook)
		transactionStatusBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		transactionStatusAfterUpdateMu.Lock()
		transactionStatusAfterUpdateHooks = append(transactionStatusAfterUpdateHooks, transactionStatusHook)
		transactionStatusAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		transactionStatusBeforeDeleteMu.Lock()
		transactionStatusBeforeDeleteHooks = append(transactionStatusBeforeDeleteHooks, transactionStatusHook)
		transactionStatusBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		transactionStatusAfterDeleteMu.Lock()
		transactionStatusAfterDeleteHooks = append(transactionStatusAfterDeleteHooks, transactionStatusHook)
		transactionStatusAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		transactionStatusBeforeUpsertMu.Lock()
		transactionStatusBeforeUpsertHooks = append(transactionStatusBeforeUpsertHooks, transactionStatusHook)
		transactionStatusBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		transactionStatusAfterUpsertMu.Lock()
		transactionStatusAfterUpsertHooks = append(transactionStatusAfterUpsertHooks, transactionStatusHook)
		transactionStatusAfterUpsertMu.Unlock()
	}
}

// OneG returns a single transactionStatus record from the query using the global executor.
func (q transactionStatusQuery) OneG(ctx context.Context) (*TransactionStatus, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single transactionStatus record from the query.
func (q transactionStatusQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TransactionStatus, error) {
	o := &TransactionStatus{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for transaction_statuses")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all TransactionStatus records from the query using the global executor.
func (q transactionStatusQuery) AllG(ctx context.Context) (TransactionStatusSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TransactionStatus records from the query.
func (q transactionStatusQuery) All(ctx context.Context, exec boil.ContextExecutor) (TransactionStatusSlice, error) {
	var o []*TransactionStatus

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TransactionStatus slice")
	}

	if len(transactionStatusAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all TransactionStatus records in the query using the global executor
func (q transactionStatusQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TransactionStatus records in the query.
func (q transactionStatusQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count transaction_statuses rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q transactionStatusQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q transactionStatusQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if transaction_statuses exists")
	}

	return count > 0, nil
}

// Transactions retrieves all the transaction's Transactions with an executor.
func (o *TransactionStatus) Transactions(mods ...qm.QueryMod) transactionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"transactions\".\"transaction_status_id\"=?", o.ID),
	)

	return Transactions(queryMods...)
}

// LoadTransactions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (transactionStatusL) LoadTransactions(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTransactionStatus interface{}, mods queries.Applicator) error {
	var slice []*TransactionStatus
	var object *TransactionStatus

	if singular {
		var ok bool
		object, ok = maybeTransactionStatus.(*TransactionStatus)
		if !ok {
			object = new(TransactionStatus)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTransactionStatus)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTransactionStatus))
			}
		}
	} else {
		s, ok := maybeTransactionStatus.(*[]*TransactionStatus)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTransactionStatus)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTransactionStatus))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &transactionStatusR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &transactionStatusR{}
			}
			args[obj.ID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`transactions`),
		qm.WhereIn(`transactions.transaction_status_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load transactions")
	}

	var resultSlice []*Transaction
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice transactions")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on transactions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for transactions")
	}

	if len(transactionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Transactions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &transactionR{}
			}
			foreign.R.TransactionStatus = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TransactionStatusID {
				local.R.Transactions = append(local.R.Transactions, foreign)
				if foreign.R == nil {
					foreign.R = &transactionR{}
				}
				foreign.R.TransactionStatus = local
				break
			}
		}
	}

	return nil
}

// AddTransactionsG adds the given related objects to the existing relationships
// of the transaction_status, optionally inserting them as new records.
// Appends related to o.R.Transactions.
// Sets related.R.TransactionStatus appropriately.
// Uses the global database handle.
func (o *TransactionStatus) AddTransactionsG(ctx context.Context, insert bool, related ...*Transaction) error {
	return o.AddTransactions(ctx, boil.GetContextDB(), insert, related...)
}

// AddTransactions adds the given related objects to the existing relationships
// of the transaction_status, optionally inserting them as new records.
// Appends related to o.R.Transactions.
// Sets related.R.TransactionStatus appropriately.
func (o *TransactionStatus) AddTransactions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Transaction) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TransactionStatusID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"transactions\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"transaction_status_id"}),
				strmangle.WhereClause("\"", "\"", 2, transactionPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.TransactionStatusID = o.ID
		}
	}

	if o.R == nil {
		o.R = &transactionStatusR{
			Transactions: related,
		}
	} else {
		o.R.Transactions = append(o.R.Transactions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &transactionR{
				TransactionStatus: o,
			}
		} else {
			rel.R.TransactionStatus = o
		}
	}
	return nil
}

// TransactionStatuses retrieves all the records using an executor.
func TransactionStatuses(mods ...qm.QueryMod) transactionStatusQuery {
	mods = append(mods, qm.From("\"transaction_statuses\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"transaction_statuses\".*"})
	}

	return transactionStatusQuery{q}
}

// FindTransactionStatusG retrieves a single record by ID.
func FindTransactionStatusG(ctx context.Context, iD int, selectCols ...string) (*TransactionStatus, error) {
	return FindTransactionStatus(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindTransactionStatus retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTransactionStatus(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*TransactionStatus, error) {
	transactionStatusObj := &TransactionStatus{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"transaction_statuses\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, transactionStatusObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from transaction_statuses")
	}

	if err = transactionStatusObj.doAfterSelectHooks(ctx, exec); err != nil {
		return transactionStatusObj, err
	}

	return transactionStatusObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TransactionStatus) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TransactionStatus) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no transaction_statuses provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(transactionStatusColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	transactionStatusInsertCacheMut.RLock()
	cache, cached := transactionStatusInsertCache[key]
	transactionStatusInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			transactionStatusAllColumns,
			transactionStatusColumnsWithDefault,
			transactionStatusColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(transactionStatusType, transactionStatusMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(transactionStatusType, transactionStatusMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"transaction_statuses\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"transaction_statuses\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into transaction_statuses")
	}

	if !cached {
		transactionStatusInsertCacheMut.Lock()
		transactionStatusInsertCache[key] = cache
		transactionStatusInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single TransactionStatus record using the global executor.
// See Update for more documentation.
func (o *TransactionStatus) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TransactionStatus.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TransactionStatus) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	transactionStatusUpdateCacheMut.RLock()
	cache, cached := transactionStatusUpdateCache[key]
	transactionStatusUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			transactionStatusAllColumns,
			transactionStatusPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update transaction_statuses, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"transaction_statuses\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, transactionStatusPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(transactionStatusType, transactionStatusMapping, append(wl, transactionStatusPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update transaction_statuses row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for transaction_statuses")
	}

	if !cached {
		transactionStatusUpdateCacheMut.Lock()
		transactionStatusUpdateCache[key] = cache
		transactionStatusUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q transactionStatusQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q transactionStatusQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for transaction_statuses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for transaction_statuses")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TransactionStatusSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TransactionStatusSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transactionStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"transaction_statuses\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, transactionStatusPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in transactionStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all transactionStatus")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TransactionStatus) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TransactionStatus) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no transaction_statuses provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(transactionStatusColumnsWithDefault, o)

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

	transactionStatusUpsertCacheMut.RLock()
	cache, cached := transactionStatusUpsertCache[key]
	transactionStatusUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			transactionStatusAllColumns,
			transactionStatusColumnsWithDefault,
			transactionStatusColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			transactionStatusAllColumns,
			transactionStatusPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert transaction_statuses, could not build update column list")
		}

		ret := strmangle.SetComplement(transactionStatusAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(transactionStatusPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert transaction_statuses, could not build conflict column list")
			}

			conflict = make([]string, len(transactionStatusPrimaryKeyColumns))
			copy(conflict, transactionStatusPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"transaction_statuses\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(transactionStatusType, transactionStatusMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(transactionStatusType, transactionStatusMapping, ret)
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
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert transaction_statuses")
	}

	if !cached {
		transactionStatusUpsertCacheMut.Lock()
		transactionStatusUpsertCache[key] = cache
		transactionStatusUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single TransactionStatus record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TransactionStatus) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TransactionStatus record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TransactionStatus) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TransactionStatus provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), transactionStatusPrimaryKeyMapping)
	sql := "DELETE FROM \"transaction_statuses\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from transaction_statuses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for transaction_statuses")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q transactionStatusQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q transactionStatusQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no transactionStatusQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from transaction_statuses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for transaction_statuses")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TransactionStatusSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TransactionStatusSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(transactionStatusBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transactionStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"transaction_statuses\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, transactionStatusPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from transactionStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for transaction_statuses")
	}

	if len(transactionStatusAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TransactionStatus) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no TransactionStatus provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TransactionStatus) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTransactionStatus(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TransactionStatusSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty TransactionStatusSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TransactionStatusSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TransactionStatusSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transactionStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"transaction_statuses\".* FROM \"transaction_statuses\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, transactionStatusPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TransactionStatusSlice")
	}

	*o = slice

	return nil
}

// TransactionStatusExistsG checks if the TransactionStatus row exists.
func TransactionStatusExistsG(ctx context.Context, iD int) (bool, error) {
	return TransactionStatusExists(ctx, boil.GetContextDB(), iD)
}

// TransactionStatusExists checks if the TransactionStatus row exists.
func TransactionStatusExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"transaction_statuses\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if transaction_statuses exists")
	}

	return exists, nil
}

// Exists checks if the TransactionStatus row exists.
func (o *TransactionStatus) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TransactionStatusExists(ctx, exec, o.ID)
}
