package sqlx

import (
	"context"
	"database/sql"
	"errors"
	"os"
)

// DB ...
type DB struct {
	x    *sql.DB
	cmd  *Cmd
	args []interface{}
}

// WrapDB ...
func WrapDB(db *sql.DB, err error) (*DB, error) {
	if err != nil {
		return nil, err
	}

	return &DB{db, NewCmd(), nil}, nil
}

var (
	_driver string
	_dsn    string
)

// Reset ...
func (db *DB) Reset() *DB {
	db.args = nil
	db.cmd.Reset()
	return db
}

// Select build select part sql.
func (db *DB) Select(fields ...string) *DB {
	db.cmd.Select(fields...)
	return db
}

// Into ...
func (db *DB) Into(table string) *DB {
	db.cmd.Into(table)
	return db
}

// From ...
func (db *DB) From(table ...string) *DB {
	db.cmd.From(table...)
	return db
}

// SubQuery ...
func (db *DB) SubQuery(sub *Cmd, as ...string) *DB {
	db.cmd.SubQuery(sub, as...)
	return db
}

// Union  ...
func (db *DB) Union(other ...*Cmd) *DB {
	db.cmd.Union(other...)
	return db
}

// Join ...
func (db *DB) Join(t JoinType, table ...string) *DB {
	db.cmd.Join(t, table...)
	return db
}

// InnerJoin ...
func (db *DB) InnerJoin(table ...string) *DB {
	db.cmd.InnerJoin(table...)
	return db
}

// LeftJoin ...
func (db *DB) LeftJoin(table ...string) *DB {
	db.cmd.LeftJoin(table...)
	return db
}

// RightJoin ...
func (db *DB) RightJoin(table ...string) *DB {
	db.cmd.RightJoin(table...)
	return db
}

// On ...
func (db *DB) On(condition ...string) *DB {
	db.cmd.On(condition...)
	return db
}

// JoinOn ...
func (db *DB) JoinOn(t JoinType, table string, condition string) *DB {
	db.cmd.JoinOn(t, table, condition)
	return db
}

// LeftJoinOn ...
func (db *DB) LeftJoinOn(table string, condition string) *DB {
	db.cmd.LeftJoinOn(table, condition)
	return db
}

// RightJoinOn ...
func (db *DB) RightJoinOn(table string, condition string) *DB {
	db.cmd.RightJoinOn(table, condition)
	return db
}

// InnerJoinOn ...
func (db *DB) InnerJoinOn(table string, condition string) *DB {
	db.cmd.InnerJoinOn(table, condition)
	return db
}

// Where ...
func (db *DB) Where(condition string) *DB {
	db.cmd.Where(condition)
	return db
}

// And ...
func (db *DB) And(condition ...string) *DB {
	db.cmd.And(condition...)
	return db
}

// Or ...
func (db *DB) Or(condition ...string) *DB {
	db.cmd.Or(condition...)
	return db
}

// WhereAnd ...
func (db *DB) WhereAnd(condition string, others ...string) *DB {
	db.cmd.WhereAnd(condition, others...)
	return db
}

// WhereOr ...
func (db *DB) WhereOr(condition string, others ...string) *DB {
	db.cmd.WhereOr(condition, others...)
	return db
}

// Insert ...
func (db *DB) Insert(table string, fields ...string) *DB {
	db.cmd.Insert(table, fields...)
	return db
}

// Values ...
func (db *DB) Values(assignments ...string) *DB {
	db.cmd.Values(assignments...)
	return db
}

// Duplicate ...
func (db *DB) Duplicate(assignments ...string) *DB {
	db.cmd.Duplicate(assignments...)
	return db
}

// DuplicateValues ...
func (db *DB) DuplicateValues(values ...string) *DB {
	db.cmd.DuplicateValues(values...)
	return db
}

// SetFields ...
func (db *DB) SetFields(fields ...string) *DB {
	db.cmd.SetFields(fields...)
	return db
}

// Set ...
func (db *DB) Set(assignments ...string) *DB {
	db.cmd.Set(assignments...)
	return db
}

// Update ...
func (db *DB) Update(table string, others ...string) *DB {
	db.cmd.Update(table, others...)
	return db
}

// Delete ...
func (db *DB) Delete(as ...string) *DB {
	db.cmd.Delete(as...)
	return db
}

// Replace ...
func (db *DB) Replace(table string, fields ...string) *DB {
	db.cmd.Replace(table, fields...)
	return db
}

// Parentheses ...
func (db *DB) Parentheses(any string) *DB {
	db.cmd.Parentheses(any)
	return db
}

// Limit ...
func (db *DB) Limit(count int, offset ...int) *DB {
	db.cmd.Limit(count, offset...)
	return db
}

// GroupBy ...
func (db *DB) GroupBy(fields ...string) *DB {
	db.cmd.GroupBy(fields...)
	return db
}

// OrderBy ...
func (db *DB) OrderBy(field string, direction ...Direction) *DB {
	db.cmd.OrderBy(field, direction...)
	return db
}

// Having ...
func (db *DB) Having(condition string) *DB {
	db.cmd.Having(condition)
	return db
}

// Raw appends string to command
func (db *DB) Raw(any string) *DB {
	db.cmd.Raw(any)
	return db
}

// SQL ...
func (db *DB) SQL(sqlcmds ...string) string {
	return db.cmd.SQL(sqlcmds...)
}

// Close ...
func (db *DB) Close() error {
	return db.x.Close()
}

// ExecContext ...
func (db *DB) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	if len(args) > 0 {
		return db.x.ExecContext(ctx, db.SQL(), args...)
	}

	return db.x.ExecContext(ctx, db.SQL(), db.args...)
}

// Exec ...
func (db *DB) Exec(args ...interface{}) (sql.Result, error) {
	return db.ExecContext(context.Background(), args...)
}

// QueryContext ...
func (db *DB) QueryContext(ctx context.Context, args ...interface{}) (*Rows, error) {
	if len(args) > 0 {
		return WrapRows(db.x.QueryContext(ctx, db.SQL(), args...))
	}
	return WrapRows(db.x.QueryContext(ctx, db.SQL(), db.args...))
}

// Query ...
func (db *DB) Query(args ...interface{}) (*Rows, error) {
	return db.QueryContext(context.Background(), args...)
}

// AllContext ...
func (db *DB) AllContext(ctx context.Context, args ...interface{}) ([]interface{}, error) {
	rows, err := db.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	return rows.All()
}

// All ...
func (db *DB) All(args ...interface{}) ([]interface{}, error) {
	return db.AllContext(context.Background(), args...)
}

// UnmarshalAllContext ...
func (db *DB) UnmarshalAllContext(ctx context.Context, x interface{}, args ...interface{}) error {
	rows, err := db.QueryContext(ctx, args...)
	if err != nil {
		return err
	}

	return rows.UnmarshalAll(x)
}

// UnmarshalAll ...
func (db *DB) UnmarshalAll(x interface{}, args ...interface{}) error {
	return db.UnmarshalAllContext(context.Background(), x, args...)
}

// QueryRowContext ...
func (db *DB) QueryRowContext(ctx context.Context, args ...interface{}) *Row {
	return WrapRow(db.QueryContext(ctx, args...))
}

// QueryRow ...
func (db *DB) QueryRow(args ...interface{}) *Row {
	return db.QueryRowContext(context.Background(), args...)
}

// DataContext ...
func (db *DB) DataContext(ctx context.Context, args ...interface{}) (interface{}, error) {
	return db.QueryRowContext(ctx, args...).Data()
}

// Data ...
func (db *DB) Data(args ...interface{}) (interface{}, error) {
	return db.DataContext(context.Background(), args...)
}

// UnmarshalContext ...
func (db *DB) UnmarshalContext(ctx context.Context, x interface{}, args ...interface{}) error {
	return db.QueryRowContext(ctx, args...).Unmarshal(x)
}

// Unmarshal ...
func (db *DB) Unmarshal(x interface{}, args ...interface{}) error {
	return db.UnmarshalContext(context.Background(), x, args...)
}

// PrepareContext ...
func (db *DB) PrepareContext(ctx context.Context) (*Stmt, error) {
	return WrapStmt(db.x.PrepareContext(ctx, db.SQL()))
}

// Prepare ...
func (db *DB) Prepare() (*Stmt, error) {
	return db.PrepareContext(context.Background())
}

// ----------------------------------------------------------------------------

// Open ....
func Open() (*DB, error) {
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
func OpenDB(driver, dsn string) (*DB, error) {
	return WrapDB(sql.Open(driver, dsn))
}
