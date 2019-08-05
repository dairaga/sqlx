package sqlx

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"strings"
	"time"
)

var (
	timeType    = reflect.TypeOf(time.Time{})
	stringType  = reflect.TypeOf("")
	boolType    = reflect.TypeOf(false)
	int64Type   = reflect.TypeOf(int64(0))
	float64Type = reflect.TypeOf(float64(0.0))

	rawBytesType = reflect.TypeOf(sql.RawBytes{})
	bytesType    = reflect.TypeOf([]byte{})

	nullStringType  = reflect.TypeOf(sql.NullString{})
	nullBoolType    = reflect.TypeOf(sql.NullBool{})
	nullInt64Type   = reflect.TypeOf(sql.NullInt64{})
	nullFloat64Type = reflect.TypeOf(sql.NullFloat64{})
)

type sqlobj interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
}

type sqlxobj struct {
	x    sqlobj
	cmd  *Cmd
	args []interface{}
}

func (z *sqlxobj) Args(x ...interface{}) *sqlxobj {
	if len(x) > 0 {
		z.args = append(z.args, x...)
	}
	return z
}

// Reset ...
func (z *sqlxobj) Reset() *sqlxobj {
	z.args = nil
	z.cmd.Reset()
	return z
}

// Select build select part sql.
func (z *sqlxobj) Select(fields ...string) *sqlxobj {
	z.cmd.Select(fields...)
	return z
}

// Into ...
func (z *sqlxobj) Into(table string) *sqlxobj {
	z.cmd.Into(table)
	return z
}

// From ...
func (z *sqlxobj) From(table ...string) *sqlxobj {
	z.cmd.From(table...)
	return z
}

// SubQuery ...
func (z *sqlxobj) SubQuery(sub *Cmd, as ...string) *sqlxobj {
	z.cmd.SubQuery(sub, as...)
	return z
}

// Union  ...
func (z *sqlxobj) Union(other ...*Cmd) *sqlxobj {
	z.cmd.Union(other...)
	return z
}

// Join ...
func (z *sqlxobj) Join(t JoinType, table ...string) *sqlxobj {
	z.cmd.Join(t, table...)
	return z
}

// InnerJoin ...
func (z *sqlxobj) InnerJoin(table ...string) *sqlxobj {
	z.cmd.InnerJoin(table...)
	return z
}

// LeftJoin ...
func (z *sqlxobj) LeftJoin(table ...string) *sqlxobj {
	z.cmd.LeftJoin(table...)
	return z
}

// RightJoin ...
func (z *sqlxobj) RightJoin(table ...string) *sqlxobj {
	z.cmd.RightJoin(table...)
	return z
}

// On ...
func (z *sqlxobj) On(condition ...string) *sqlxobj {
	z.cmd.On(condition...)
	return z
}

// JoinOn ...
func (z *sqlxobj) JoinOn(t JoinType, table string, condition string) *sqlxobj {
	z.cmd.JoinOn(t, table, condition)
	return z
}

// LeftJoinOn ...
func (z *sqlxobj) LeftJoinOn(table string, condition string) *sqlxobj {
	z.cmd.LeftJoinOn(table, condition)
	return z
}

// RightJoinOn ...
func (z *sqlxobj) RightJoinOn(table string, condition string) *sqlxobj {
	z.cmd.RightJoinOn(table, condition)
	return z
}

// InnerJoinOn ...
func (z *sqlxobj) InnerJoinOn(table string, condition string) *sqlxobj {
	z.cmd.InnerJoinOn(table, condition)
	return z
}

// Where ...
func (z *sqlxobj) Where(condition string) *sqlxobj {
	z.cmd.Where(condition)
	return z
}

// And ...
func (z *sqlxobj) And(condition ...string) *sqlxobj {
	z.cmd.And(condition...)
	return z
}

// Or ...
func (z *sqlxobj) Or(condition ...string) *sqlxobj {
	z.cmd.Or(condition...)
	return z
}

// WhereAnd ...
func (z *sqlxobj) WhereAnd(condition string, others ...string) *sqlxobj {
	z.cmd.WhereAnd(condition, others...)
	return z
}

// WhereOr ...
func (z *sqlxobj) WhereOr(condition string, others ...string) *sqlxobj {
	z.cmd.WhereOr(condition, others...)
	return z
}

// Insert ...
func (z *sqlxobj) Insert(table string, fields ...string) *sqlxobj {
	z.cmd.Insert(table, fields...)
	return z
}

func (z *sqlxobj) InsertValues(table string, fields ...string) *sqlxobj {
	z.cmd.InsertValues(table, fields...)
	return z
}

// Values ...
func (z *sqlxobj) Values(assignments ...string) *sqlxobj {
	z.cmd.Values(assignments...)
	return z
}

// Duplicate ...
func (z *sqlxobj) Duplicate(assignments ...string) *sqlxobj {
	z.cmd.Duplicate(assignments...)
	return z
}

// DuplicateValues ...
func (z *sqlxobj) DuplicateValues(values ...string) *sqlxobj {
	z.cmd.DuplicateValues(values...)
	return z
}

// SetFields ...
func (z *sqlxobj) SetFields(fields ...string) *sqlxobj {
	z.cmd.SetFields(fields...)
	return z
}

// Set ...
func (z *sqlxobj) Set(assignments ...string) *sqlxobj {
	z.cmd.Set(assignments...)
	return z
}

// Update ...
func (z *sqlxobj) Update(table string, others ...string) *sqlxobj {
	z.cmd.Update(table, others...)
	return z
}

// Delete ...
func (z *sqlxobj) Delete(as ...string) *sqlxobj {
	z.cmd.Delete(as...)
	return z
}

func (z *sqlxobj) DeleteFrom(table ...string) *sqlxobj {
	z.cmd.DeleteFrom(table...)
	return z
}

// Replace ...
func (z *sqlxobj) Replace(table string, fields ...string) *sqlxobj {
	z.cmd.Replace(table, fields...)
	return z
}

// Parentheses ...
func (z *sqlxobj) Parentheses(any string) *sqlxobj {
	z.cmd.Parentheses(any)
	return z
}

// Limit ...
func (z *sqlxobj) Limit(count int, offset ...int) *sqlxobj {
	z.cmd.Limit(count, offset...)
	return z
}

// GroupBy ...
func (z *sqlxobj) GroupBy(fields ...string) *sqlxobj {
	z.cmd.GroupBy(fields...)
	return z
}

// OrderBy ...
func (z *sqlxobj) OrderBy(field string, direction ...Direction) *sqlxobj {
	z.cmd.OrderBy(field, direction...)
	return z
}

// Having ...
func (z *sqlxobj) Having(condition string) *sqlxobj {
	z.cmd.Having(condition)
	return z
}

// Raw appends string to command
func (z *sqlxobj) Raw(any string) *sqlxobj {
	z.cmd.Raw(any)
	return z
}

// SQL ...
func (z *sqlxobj) SQL(sqlcmds ...string) string {
	return z.cmd.SQL(sqlcmds...)
}

// ExecContext ...
func (z *sqlxobj) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	if len(args) > 0 {
		return z.x.ExecContext(ctx, z.SQL(), args...)
	}

	return z.x.ExecContext(ctx, z.SQL(), z.args...)
}

// Exec ...
func (z *sqlxobj) Exec(args ...interface{}) (sql.Result, error) {
	return z.ExecContext(context.Background(), args...)
}

func (z *sqlxobj) CountContext(ctx context.Context, args ...interface{}) (int, error) {

	newSQL := z.cmd.SQL()
	pos := strings.Index(newSQL, "FROM")
	if pos <= 0 {
		return 0, errors.New("sql command error")
	}

	newSQL = "SELECT count(*) as cc " + newSQL[pos:]

	oldCmd := z.cmd

	defer func() {
		z.cmd = oldCmd
	}()

	newCmd := NewCmd()

	newCmd.Raw(newSQL)
	z.cmd = newCmd

	row := z.QueryRowContext(ctx, args...)
	if row.Err() != nil {
		return 0, row.Err()
	}

	return row.GetInt("cc", 0), nil
}

func (z *sqlxobj) Count(args ...interface{}) (int, error) {
	return z.CountContext(context.Background(), args...)
}

// QueryContext ...
func (z *sqlxobj) QueryContext(ctx context.Context, args ...interface{}) (*Rows, error) {
	if len(args) > 0 {
		return WrapRows(z.x.QueryContext(ctx, z.SQL(), args...))
	}
	return WrapRows(z.x.QueryContext(ctx, z.SQL(), z.args...))
}

// Query ...
func (z *sqlxobj) Query(args ...interface{}) (*Rows, error) {
	return z.QueryContext(context.Background(), args...)
}

// AllContext ...
func (z *sqlxobj) AllContext(ctx context.Context, args ...interface{}) ([]interface{}, error) {
	rows, err := z.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	return rows.All()
}

// All ...
func (z *sqlxobj) All(args ...interface{}) ([]interface{}, error) {
	return z.AllContext(context.Background(), args...)
}

// UnmarshalAllContext ...
func (z *sqlxobj) UnmarshalAllContext(ctx context.Context, x interface{}, args ...interface{}) error {
	rows, err := z.QueryContext(ctx, args...)
	if err != nil {
		return err
	}

	return rows.UnmarshalAll(x)
}

// UnmarshalAll ...
func (z *sqlxobj) UnmarshalAll(x interface{}, args ...interface{}) error {
	return z.UnmarshalAllContext(context.Background(), x, args...)
}

// QueryRowContext ...
func (z *sqlxobj) QueryRowContext(ctx context.Context, args ...interface{}) *Row {
	return WrapRow(z.QueryContext(ctx, args...))
}

// QueryRow ...
func (z *sqlxobj) QueryRow(args ...interface{}) *Row {
	return z.QueryRowContext(context.Background(), args...)
}

// DataContext ...
func (z *sqlxobj) DataContext(ctx context.Context, args ...interface{}) (interface{}, error) {
	return z.QueryRowContext(ctx, args...).Data()
}

// Data ...
func (z *sqlxobj) Data(args ...interface{}) (interface{}, error) {
	return z.DataContext(context.Background(), args...)
}

// UnmarshalContext ...
func (z *sqlxobj) UnmarshalContext(ctx context.Context, x interface{}, args ...interface{}) error {
	return z.QueryRowContext(ctx, args...).Unmarshal(x)
}

// Unmarshal ...
func (z *sqlxobj) Unmarshal(x interface{}, args ...interface{}) error {
	return z.UnmarshalContext(context.Background(), x, args...)
}

// PrepareContext ...
func (z *sqlxobj) PrepareContext(ctx context.Context) (*Stmt, error) {
	return WrapStmt(z.x.PrepareContext(ctx, z.SQL()))
}

// Prepare ...
func (z *sqlxobj) Prepare() (*Stmt, error) {
	return z.PrepareContext(context.Background())
}
