package sqlx

import (
	"context"
	"database/sql"
	"errors"
	"os"
)

// DBX ...
type DBX struct {
	x    *DB
	cmd  *Cmd
	args []interface{}
}

var (
	_driver string
	_dsn    string
)

// Reset ...
func (db *DBX) Reset() *DBX {
	db.args = nil
	db.cmd.Reset()
	return db
}

// Select build select part sql.
func (db *DBX) Select(fields ...string) *DBX {
	db.cmd.Select(fields...)
	return db
}

// Into ...
func (db *DBX) Into(table string) *DBX {
	db.cmd.Into(table)
	return db
}

// From ...
func (db *DBX) From(table ...string) *DBX {
	db.cmd.From(table...)
	return db
}

// SubQuery ...
func (db *DBX) SubQuery(sub *Cmd, as ...string) *DBX {
	db.cmd.SubQuery(sub, as...)
	return db
}

// Union  ...
func (db *DBX) Union(other ...*Cmd) *DBX {
	db.cmd.Union(other...)
	return db
}

// Join ...
func (db *DBX) Join(t JoinType, table ...string) *DBX {
	db.cmd.Join(t, table...)
	return db
}

// InnerJoin ...
func (db *DBX) InnerJoin(table ...string) *DBX {
	db.cmd.InnerJoin(table...)
	return db
}

// LeftJoin ...
func (db *DBX) LeftJoin(table ...string) *DBX {
	db.cmd.LeftJoin(table...)
	return db
}

// RightJoin ...
func (db *DBX) RightJoin(table ...string) *DBX {
	db.cmd.RightJoin(table...)
	return db
}

// On ...
func (db *DBX) On(condition ...string) *DBX {
	db.cmd.On(condition...)
	return db
}

// JoinOn ...
func (db *DBX) JoinOn(t JoinType, table string, condition string) *DBX {
	db.cmd.JoinOn(t, table, condition)
	return db
}

// LeftJoinOn ...
func (db *DBX) LeftJoinOn(table string, condition string) *DBX {
	db.cmd.LeftJoinOn(table, condition)
	return db
}

// RightJoinOn ...
func (db *DBX) RightJoinOn(table string, condition string) *DBX {
	db.cmd.RightJoinOn(table, condition)
	return db
}

// InnerJoinOn ...
func (db *DBX) InnerJoinOn(table string, condition string) *DBX {
	db.cmd.InnerJoinOn(table, condition)
	return db
}

// Where ...
func (db *DBX) Where(condition string) *DBX {
	db.Where(condition)
	return db
}

// And ...
func (db *DBX) And(condition ...string) *DBX {
	db.cmd.And(condition...)
	return db
}

// Or ...
func (db *DBX) Or(condition ...string) *DBX {
	db.cmd.Or(condition...)
	return db
}

// WhereAnd ...
func (db *DBX) WhereAnd(condition string, others ...string) *DBX {
	db.cmd.WhereAnd(condition, others...)
	return db
}

// WhereOr ...
func (db *DBX) WhereOr(condition string, others ...string) *DBX {
	db.cmd.WhereOr(condition, others...)
	return db
}

// Insert ...
func (db *DBX) Insert(table string, fields ...string) *DBX {
	db.cmd.Insert(table, fields...)
	return db
}

// Values ...
func (db *DBX) Values(assignments ...string) *DBX {
	db.cmd.Values(assignments...)
	return db
}

// Duplicate ...
func (db *DBX) Duplicate(assignments ...string) *DBX {
	db.cmd.Duplicate(assignments...)
	return db
}

// DuplicateValues ...
func (db *DBX) DuplicateValues(values ...string) *DBX {
	db.cmd.DuplicateValues(values...)
	return db
}

// SetFields ...
func (db *DBX) SetFields(fields ...string) *DBX {
	db.cmd.SetFields(fields...)
	return db
}

// Set ...
func (db *DBX) Set(assignments ...string) *DBX {
	db.cmd.Set(assignments...)
	return db
}

// Update ...
func (db *DBX) Update(table string, others ...string) *DBX {
	db.cmd.Update(table, others...)
	return db
}

// Delete ...
func (db *DBX) Delete(as ...string) *DBX {
	db.cmd.Delete(as...)
	return db
}

// Replace ...
func (db *DBX) Replace(table string, fields ...string) *DBX {
	db.cmd.Replace(table, fields...)
	return db
}

// Parentheses ...
func (db *DBX) Parentheses(any string) *DBX {
	db.cmd.Parentheses(any)
	return db
}

// Limit ...
func (db *DBX) Limit(count int, offset ...int) *DBX {
	db.cmd.Limit(count, offset...)
	return db
}

// GroupBy ...
func (db *DBX) GroupBy(fields ...string) *DBX {
	db.cmd.GroupBy(fields...)
	return db
}

// OrderBy ...
func (db *DBX) OrderBy(field string, direction ...Direction) *DBX {
	db.cmd.OrderBy(field, direction...)
	return db
}

// Having ...
func (db *DBX) Having(condition string) *DBX {
	db.cmd.Having(condition)
	return db
}

// Raw appends string to command
func (db *DBX) Raw(any string) *DBX {
	db.cmd.Raw(any)
	return db
}

// SQL ...
func (db *DBX) SQL() string {
	return db.cmd.String()
}

// Close ...
func (db *DBX) Close() error {
	return db.x.Close()
}

// ExecContext ...
func (db *DBX) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	if len(args) > 0 {
		return db.x.ExecContext(ctx, db.SQL(), args...)
	}

	return db.x.ExecContext(ctx, db.SQL(), db.args...)
}

// Exec ...
func (db *DBX) Exec(args ...interface{}) (sql.Result, error) {
	return db.ExecContext(context.Background(), args...)
}

// QueryContext ...
func (db *DBX) QueryContext(ctx context.Context, args ...interface{}) (*Rows, error) {
	if len(args) > 0 {
		return db.x.FindContext(ctx, db.SQL(), args...)
	}
	return db.x.FindContext(ctx, db.SQL(), db.args...)
}

// Query ...
func (db *DBX) Query(args ...interface{}) (*Rows, error) {
	return db.QueryContext(context.Background(), args...)
}

// AllContext ...
func (db *DBX) AllContext(ctx context.Context, args ...interface{}) ([]interface{}, error) {
	if len(args) > 0 {
		return db.x.AllContext(ctx, db.SQL(), args...)
	}

	return db.x.AllContext(ctx, db.SQL(), db.args...)
}

// All ...
func (db *DBX) All(args ...interface{}) ([]interface{}, error) {
	return db.AllContext(context.Background(), args...)
}

// UnmarshalAllContext ...
func (db *DBX) UnmarshalAllContext(ctx context.Context, x interface{}, args ...interface{}) error {
	if len(args) > 0 {
		return db.x.UnmarshalAllContext(ctx, x, db.SQL(), args...)
	}
	return db.x.UnmarshalAllContext(ctx, x, db.SQL(), db.args...)
}

// UnmarshalAll ...
func (db *DBX) UnmarshalAll(x interface{}, args ...interface{}) error {
	return db.UnmarshalAllContext(context.Background(), x, args...)
}

// QueryRowContext ...
func (db *DBX) QueryRowContext(ctx context.Context, args ...interface{}) *Row {
	if len(args) > 0 {
		return db.x.FindRowContext(ctx, db.SQL(), args...)
	}
	return db.x.FindRowContext(ctx, db.SQL(), db.args...)
}

// QueryRow ...
func (db *DBX) QueryRow(args ...interface{}) *Row {
	return db.QueryRowContext(context.Background(), args...)
}

// DataContext ...
func (db *DBX) DataContext(ctx context.Context, args ...interface{}) (interface{}, error) {
	if len(args) > 0 {
		return db.x.DataContext(ctx, db.SQL(), args...)
	}
	return db.x.DataContext(ctx, db.SQL(), db.args...)
}

// Data ...
func (db *DBX) Data(args ...interface{}) (interface{}, error) {
	return db.DataContext(context.Background(), args...)
}

// UnmarshalContext ...
func (db *DBX) UnmarshalContext(ctx context.Context, x interface{}, args ...interface{}) error {
	if len(args) > 0 {
		return db.x.UnmarshalContext(ctx, x, db.SQL(), args...)
	}
	return db.x.UnmarshalContext(ctx, x, db.SQL(), db.args...)
}

// Unmarshal ...
func (db *DBX) Unmarshal(x interface{}, args ...interface{}) error {
	return db.UnmarshalContext(context.Background(), x, args...)
}

// ----------------------------------------------------------------------------

// Open ....
func Open() (*DBX, error) {
	if _driver == "" || _dsn == "" {
		_driver, _ = os.LookupEnv("DB_DRIVER")
		if _driver == "" {
			return nil, errors.New("can not find Database Driver")
		}

		_dsn, _ = os.LookupEnv("DB_DSN")
		if _dsn == "" {
			return nil, errors.New("can not find Database Data Source Name")
		}
	}

	return OpenDB(_driver, _dsn)
}

// OpenDB ...
func OpenDB(driver, dsn string) (*DBX, error) {
	db, err := WrapDB(sql.Open(driver, dsn))
	if err != nil {
		return nil, err
	}
	return &DBX{db, NewCmd(), nil}, nil
}
