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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Expense is an object representing the database table.
type Expense struct {
	ExpenseID          null.Int64  `boil:"expense_id" json:"expense_id,omitempty" toml:"expense_id" yaml:"expense_id,omitempty"`
	ExpenseEpoch       float64     `boil:"expense_epoch" json:"expense_epoch" toml:"expense_epoch" yaml:"expense_epoch"`
	ExpenseAmount      float64     `boil:"expense_amount" json:"expense_amount" toml:"expense_amount" yaml:"expense_amount"`
	ExpenseDescription null.String `boil:"expense_description" json:"expense_description,omitempty" toml:"expense_description" yaml:"expense_description,omitempty"`
	ExpenseCategoryID  null.Int64  `boil:"expense_category_id" json:"expense_category_id,omitempty" toml:"expense_category_id" yaml:"expense_category_id,omitempty"`

	R *expenseR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L expenseL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ExpenseColumns = struct {
	ExpenseID          string
	ExpenseEpoch       string
	ExpenseAmount      string
	ExpenseDescription string
	ExpenseCategoryID  string
}{
	ExpenseID:          "expense_id",
	ExpenseEpoch:       "expense_epoch",
	ExpenseAmount:      "expense_amount",
	ExpenseDescription: "expense_description",
	ExpenseCategoryID:  "expense_category_id",
}

var ExpenseTableColumns = struct {
	ExpenseID          string
	ExpenseEpoch       string
	ExpenseAmount      string
	ExpenseDescription string
	ExpenseCategoryID  string
}{
	ExpenseID:          "expense.expense_id",
	ExpenseEpoch:       "expense.expense_epoch",
	ExpenseAmount:      "expense.expense_amount",
	ExpenseDescription: "expense.expense_description",
	ExpenseCategoryID:  "expense.expense_category_id",
}

// Generated where

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
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
func (w whereHelpernull_String) LIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" LIKE ?", x)
}
func (w whereHelpernull_String) NLIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT LIKE ?", x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var ExpenseWhere = struct {
	ExpenseID          whereHelpernull_Int64
	ExpenseEpoch       whereHelperfloat64
	ExpenseAmount      whereHelperfloat64
	ExpenseDescription whereHelpernull_String
	ExpenseCategoryID  whereHelpernull_Int64
}{
	ExpenseID:          whereHelpernull_Int64{field: "\"expense\".\"expense_id\""},
	ExpenseEpoch:       whereHelperfloat64{field: "\"expense\".\"expense_epoch\""},
	ExpenseAmount:      whereHelperfloat64{field: "\"expense\".\"expense_amount\""},
	ExpenseDescription: whereHelpernull_String{field: "\"expense\".\"expense_description\""},
	ExpenseCategoryID:  whereHelpernull_Int64{field: "\"expense\".\"expense_category_id\""},
}

// ExpenseRels is where relationship names are stored.
var ExpenseRels = struct {
	ExpenseCategory string
}{
	ExpenseCategory: "ExpenseCategory",
}

// expenseR is where relationships are stored.
type expenseR struct {
	ExpenseCategory *Category `boil:"ExpenseCategory" json:"ExpenseCategory" toml:"ExpenseCategory" yaml:"ExpenseCategory"`
}

// NewStruct creates a new relationship struct
func (*expenseR) NewStruct() *expenseR {
	return &expenseR{}
}

func (r *expenseR) GetExpenseCategory() *Category {
	if r == nil {
		return nil
	}
	return r.ExpenseCategory
}

// expenseL is where Load methods for each relationship are stored.
type expenseL struct{}

var (
	expenseAllColumns            = []string{"expense_id", "expense_epoch", "expense_amount", "expense_description", "expense_category_id"}
	expenseColumnsWithoutDefault = []string{"expense_epoch", "expense_amount"}
	expenseColumnsWithDefault    = []string{"expense_id", "expense_description", "expense_category_id"}
	expensePrimaryKeyColumns     = []string{"expense_id"}
	expenseGeneratedColumns      = []string{"expense_id"}
)

type (
	// ExpenseSlice is an alias for a slice of pointers to Expense.
	// This should almost always be used instead of []Expense.
	ExpenseSlice []*Expense
	// ExpenseHook is the signature for custom Expense hook methods
	ExpenseHook func(context.Context, boil.ContextExecutor, *Expense) error

	expenseQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	expenseType                 = reflect.TypeOf(&Expense{})
	expenseMapping              = queries.MakeStructMapping(expenseType)
	expensePrimaryKeyMapping, _ = queries.BindMapping(expenseType, expenseMapping, expensePrimaryKeyColumns)
	expenseInsertCacheMut       sync.RWMutex
	expenseInsertCache          = make(map[string]insertCache)
	expenseUpdateCacheMut       sync.RWMutex
	expenseUpdateCache          = make(map[string]updateCache)
	expenseUpsertCacheMut       sync.RWMutex
	expenseUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var expenseAfterSelectMu sync.Mutex
var expenseAfterSelectHooks []ExpenseHook

var expenseBeforeInsertMu sync.Mutex
var expenseBeforeInsertHooks []ExpenseHook
var expenseAfterInsertMu sync.Mutex
var expenseAfterInsertHooks []ExpenseHook

var expenseBeforeUpdateMu sync.Mutex
var expenseBeforeUpdateHooks []ExpenseHook
var expenseAfterUpdateMu sync.Mutex
var expenseAfterUpdateHooks []ExpenseHook

var expenseBeforeDeleteMu sync.Mutex
var expenseBeforeDeleteHooks []ExpenseHook
var expenseAfterDeleteMu sync.Mutex
var expenseAfterDeleteHooks []ExpenseHook

var expenseBeforeUpsertMu sync.Mutex
var expenseBeforeUpsertHooks []ExpenseHook
var expenseAfterUpsertMu sync.Mutex
var expenseAfterUpsertHooks []ExpenseHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Expense) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Expense) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Expense) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Expense) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Expense) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Expense) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Expense) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Expense) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Expense) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddExpenseHook registers your hook function for all future operations.
func AddExpenseHook(hookPoint boil.HookPoint, expenseHook ExpenseHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		expenseAfterSelectMu.Lock()
		expenseAfterSelectHooks = append(expenseAfterSelectHooks, expenseHook)
		expenseAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		expenseBeforeInsertMu.Lock()
		expenseBeforeInsertHooks = append(expenseBeforeInsertHooks, expenseHook)
		expenseBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		expenseAfterInsertMu.Lock()
		expenseAfterInsertHooks = append(expenseAfterInsertHooks, expenseHook)
		expenseAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		expenseBeforeUpdateMu.Lock()
		expenseBeforeUpdateHooks = append(expenseBeforeUpdateHooks, expenseHook)
		expenseBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		expenseAfterUpdateMu.Lock()
		expenseAfterUpdateHooks = append(expenseAfterUpdateHooks, expenseHook)
		expenseAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		expenseBeforeDeleteMu.Lock()
		expenseBeforeDeleteHooks = append(expenseBeforeDeleteHooks, expenseHook)
		expenseBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		expenseAfterDeleteMu.Lock()
		expenseAfterDeleteHooks = append(expenseAfterDeleteHooks, expenseHook)
		expenseAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		expenseBeforeUpsertMu.Lock()
		expenseBeforeUpsertHooks = append(expenseBeforeUpsertHooks, expenseHook)
		expenseBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		expenseAfterUpsertMu.Lock()
		expenseAfterUpsertHooks = append(expenseAfterUpsertHooks, expenseHook)
		expenseAfterUpsertMu.Unlock()
	}
}

// OneG returns a single expense record from the query using the global executor.
func (q expenseQuery) OneG(ctx context.Context) (*Expense, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single expense record from the query.
func (q expenseQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Expense, error) {
	o := &Expense{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for expense")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Expense records from the query using the global executor.
func (q expenseQuery) AllG(ctx context.Context) (ExpenseSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Expense records from the query.
func (q expenseQuery) All(ctx context.Context, exec boil.ContextExecutor) (ExpenseSlice, error) {
	var o []*Expense

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Expense slice")
	}

	if len(expenseAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Expense records in the query using the global executor
func (q expenseQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Expense records in the query.
func (q expenseQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count expense rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q expenseQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q expenseQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if expense exists")
	}

	return count > 0, nil
}

// ExpenseCategory pointed to by the foreign key.
func (o *Expense) ExpenseCategory(mods ...qm.QueryMod) categoryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"category_id\" = ?", o.ExpenseCategoryID),
	}

	queryMods = append(queryMods, mods...)

	return Categories(queryMods...)
}

// LoadExpenseCategory allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (expenseL) LoadExpenseCategory(ctx context.Context, e boil.ContextExecutor, singular bool, maybeExpense interface{}, mods queries.Applicator) error {
	var slice []*Expense
	var object *Expense

	if singular {
		var ok bool
		object, ok = maybeExpense.(*Expense)
		if !ok {
			object = new(Expense)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeExpense)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeExpense))
			}
		}
	} else {
		s, ok := maybeExpense.(*[]*Expense)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeExpense)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeExpense))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &expenseR{}
		}
		if !queries.IsNil(object.ExpenseCategoryID) {
			args[object.ExpenseCategoryID] = struct{}{}
		}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &expenseR{}
			}

			if !queries.IsNil(obj.ExpenseCategoryID) {
				args[obj.ExpenseCategoryID] = struct{}{}
			}

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
		qm.From(`category`),
		qm.WhereIn(`category.category_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Category")
	}

	var resultSlice []*Category
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Category")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for category")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for category")
	}

	if len(categoryAfterSelectHooks) != 0 {
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
		object.R.ExpenseCategory = foreign
		if foreign.R == nil {
			foreign.R = &categoryR{}
		}
		foreign.R.ExpenseCategoryExpenses = append(foreign.R.ExpenseCategoryExpenses, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ExpenseCategoryID, foreign.CategoryID) {
				local.R.ExpenseCategory = foreign
				if foreign.R == nil {
					foreign.R = &categoryR{}
				}
				foreign.R.ExpenseCategoryExpenses = append(foreign.R.ExpenseCategoryExpenses, local)
				break
			}
		}
	}

	return nil
}

// SetExpenseCategoryG of the expense to the related item.
// Sets o.R.ExpenseCategory to related.
// Adds o to related.R.ExpenseCategoryExpenses.
// Uses the global database handle.
func (o *Expense) SetExpenseCategoryG(ctx context.Context, insert bool, related *Category) error {
	return o.SetExpenseCategory(ctx, boil.GetContextDB(), insert, related)
}

// SetExpenseCategory of the expense to the related item.
// Sets o.R.ExpenseCategory to related.
// Adds o to related.R.ExpenseCategoryExpenses.
func (o *Expense) SetExpenseCategory(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Category) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"expense\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"expense_category_id"}),
		strmangle.WhereClause("\"", "\"", 0, expensePrimaryKeyColumns),
	)
	values := []interface{}{related.CategoryID, o.ExpenseID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.ExpenseCategoryID, related.CategoryID)
	if o.R == nil {
		o.R = &expenseR{
			ExpenseCategory: related,
		}
	} else {
		o.R.ExpenseCategory = related
	}

	if related.R == nil {
		related.R = &categoryR{
			ExpenseCategoryExpenses: ExpenseSlice{o},
		}
	} else {
		related.R.ExpenseCategoryExpenses = append(related.R.ExpenseCategoryExpenses, o)
	}

	return nil
}

// RemoveExpenseCategoryG relationship.
// Sets o.R.ExpenseCategory to nil.
// Removes o from all passed in related items' relationships struct.
// Uses the global database handle.
func (o *Expense) RemoveExpenseCategoryG(ctx context.Context, related *Category) error {
	return o.RemoveExpenseCategory(ctx, boil.GetContextDB(), related)
}

// RemoveExpenseCategory relationship.
// Sets o.R.ExpenseCategory to nil.
// Removes o from all passed in related items' relationships struct.
func (o *Expense) RemoveExpenseCategory(ctx context.Context, exec boil.ContextExecutor, related *Category) error {
	var err error

	queries.SetScanner(&o.ExpenseCategoryID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("expense_category_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.ExpenseCategory = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ExpenseCategoryExpenses {
		if queries.Equal(o.ExpenseCategoryID, ri.ExpenseCategoryID) {
			continue
		}

		ln := len(related.R.ExpenseCategoryExpenses)
		if ln > 1 && i < ln-1 {
			related.R.ExpenseCategoryExpenses[i] = related.R.ExpenseCategoryExpenses[ln-1]
		}
		related.R.ExpenseCategoryExpenses = related.R.ExpenseCategoryExpenses[:ln-1]
		break
	}
	return nil
}

// Expenses retrieves all the records using an executor.
func Expenses(mods ...qm.QueryMod) expenseQuery {
	mods = append(mods, qm.From("\"expense\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"expense\".*"})
	}

	return expenseQuery{q}
}

// FindExpenseG retrieves a single record by ID.
func FindExpenseG(ctx context.Context, expenseID null.Int64, selectCols ...string) (*Expense, error) {
	return FindExpense(ctx, boil.GetContextDB(), expenseID, selectCols...)
}

// FindExpense retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindExpense(ctx context.Context, exec boil.ContextExecutor, expenseID null.Int64, selectCols ...string) (*Expense, error) {
	expenseObj := &Expense{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"expense\" where \"expense_id\"=?", sel,
	)

	q := queries.Raw(query, expenseID)

	err := q.Bind(ctx, exec, expenseObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from expense")
	}

	if err = expenseObj.doAfterSelectHooks(ctx, exec); err != nil {
		return expenseObj, err
	}

	return expenseObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Expense) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Expense) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no expense provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(expenseColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	expenseInsertCacheMut.RLock()
	cache, cached := expenseInsertCache[key]
	expenseInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			expenseAllColumns,
			expenseColumnsWithDefault,
			expenseColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, expenseGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(expenseType, expenseMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(expenseType, expenseMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"expense\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"expense\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into expense")
	}

	if !cached {
		expenseInsertCacheMut.Lock()
		expenseInsertCache[key] = cache
		expenseInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Expense record using the global executor.
// See Update for more documentation.
func (o *Expense) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Expense.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Expense) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	expenseUpdateCacheMut.RLock()
	cache, cached := expenseUpdateCache[key]
	expenseUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			expenseAllColumns,
			expensePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, expenseGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update expense, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"expense\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, expensePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(expenseType, expenseMapping, append(wl, expensePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update expense row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for expense")
	}

	if !cached {
		expenseUpdateCacheMut.Lock()
		expenseUpdateCache[key] = cache
		expenseUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q expenseQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q expenseQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for expense")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for expense")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ExpenseSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ExpenseSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), expensePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"expense\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, expensePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in expense slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all expense")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Expense) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Expense) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no expense provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(expenseColumnsWithDefault, o)

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

	expenseUpsertCacheMut.RLock()
	cache, cached := expenseUpsertCache[key]
	expenseUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			expenseAllColumns,
			expenseColumnsWithDefault,
			expenseColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			expenseAllColumns,
			expensePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert expense, could not build update column list")
		}

		ret := strmangle.SetComplement(expenseAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(expensePrimaryKeyColumns))
			copy(conflict, expensePrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"expense\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(expenseType, expenseMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(expenseType, expenseMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert expense")
	}

	if !cached {
		expenseUpsertCacheMut.Lock()
		expenseUpsertCache[key] = cache
		expenseUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Expense record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Expense) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Expense record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Expense) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Expense provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), expensePrimaryKeyMapping)
	sql := "DELETE FROM \"expense\" WHERE \"expense_id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from expense")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for expense")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q expenseQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q expenseQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no expenseQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from expense")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for expense")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ExpenseSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ExpenseSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(expenseBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), expensePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"expense\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, expensePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from expense slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for expense")
	}

	if len(expenseAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Expense) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Expense provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Expense) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindExpense(ctx, exec, o.ExpenseID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ExpenseSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty ExpenseSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ExpenseSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ExpenseSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), expensePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"expense\".* FROM \"expense\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, expensePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ExpenseSlice")
	}

	*o = slice

	return nil
}

// ExpenseExistsG checks if the Expense row exists.
func ExpenseExistsG(ctx context.Context, expenseID null.Int64) (bool, error) {
	return ExpenseExists(ctx, boil.GetContextDB(), expenseID)
}

// ExpenseExists checks if the Expense row exists.
func ExpenseExists(ctx context.Context, exec boil.ContextExecutor, expenseID null.Int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"expense\" where \"expense_id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, expenseID)
	}
	row := exec.QueryRowContext(ctx, sql, expenseID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if expense exists")
	}

	return exists, nil
}

// Exists checks if the Expense row exists.
func (o *Expense) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ExpenseExists(ctx, exec, o.ExpenseID)
}
