package sqlx

import (
	"context"
	"database/sql"
)

// Tx wraps sql.Tx
type Tx struct {
	x    *sql.Tx
	cmd  *Cmd
	args []interface{}
}

// Reset ...
func (t *Tx) Reset() *Tx {
	t.cmd.Reset()
	t.args = nil

	return t
}

// Commit ...
func (t *Tx) Commit() error {
	return t.x.Commit()
}

// Rollback ...
func (t *Tx) Rollback() error {
	return t.x.Rollback()
}

// Select build select part sql.
func (t *Tx) Select(fields ...string) *Tx {
	t.cmd.Select(fields...)
	return t
}

// Into ...
func (t *Tx) Into(table string) *Tx {
	t.cmd.Into(table)
	return t
}

// From ...
func (t *Tx) From(table ...string) *Tx {
	t.cmd.From(table...)
	return t
}

// SubQuery ...
func (t *Tx) SubQuery(sub *Cmd, as ...string) *Tx {
	t.cmd.SubQuery(sub, as...)
	return t
}

// Union  ...
func (t *Tx) Union(other ...*Cmd) *Tx {
	t.cmd.Union(other...)
	return t
}

// Join ...
func (t *Tx) Join(jt JoinType, table ...string) *Tx {
	t.cmd.Join(jt, table...)
	return t
}

// InnerJoin ...
func (t *Tx) InnerJoin(table ...string) *Tx {
	t.cmd.InnerJoin(table...)
	return t
}

// LeftJoin ...
func (t *Tx) LeftJoin(table ...string) *Tx {
	t.cmd.LeftJoin(table...)
	return t
}

// RightJoin ...
func (t *Tx) RightJoin(table ...string) *Tx {
	t.cmd.RightJoin(table...)
	return t
}

// On ...
func (t *Tx) On(condition ...string) *Tx {
	t.cmd.On(condition...)
	return t
}

// JoinOn ...
func (t *Tx) JoinOn(jt JoinType, table string, condition string) *Tx {
	t.cmd.JoinOn(jt, table, condition)
	return t
}

// LeftJoinOn ...
func (t *Tx) LeftJoinOn(table string, condition string) *Tx {
	t.cmd.LeftJoinOn(table, condition)
	return t
}

// RightJoinOn ...
func (t *Tx) RightJoinOn(table string, condition string) *Tx {
	t.cmd.RightJoinOn(table, condition)
	return t
}

// InnerJoinOn ...
func (t *Tx) InnerJoinOn(table string, condition string) *Tx {
	t.cmd.InnerJoinOn(table, condition)
	return t
}

// Where ...
func (t *Tx) Where(condition string) *Tx {
	t.cmd.Where(condition)
	return t
}

// And ...
func (t *Tx) And(condition ...string) *Tx {
	t.cmd.And(condition...)
	return t
}

// Or ...
func (t *Tx) Or(condition ...string) *Tx {
	t.cmd.Or(condition...)
	return t
}

// WhereAnd ...
func (t *Tx) WhereAnd(condition string, others ...string) *Tx {
	t.cmd.WhereAnd(condition, others...)
	return t
}

// WhereOr ...
func (t *Tx) WhereOr(condition string, others ...string) *Tx {
	t.cmd.WhereOr(condition, others...)
	return t
}

// Insert ...
func (t *Tx) Insert(table string, fields ...string) *Tx {
	t.cmd.Insert(table, fields...)
	return t
}

// Values ...
func (t *Tx) Values(assignments ...string) *Tx {
	t.cmd.Values(assignments...)
	return t
}

// Duplicate ...
func (t *Tx) Duplicate(assignments ...string) *Tx {
	t.cmd.Duplicate(assignments...)
	return t
}

// DuplicateValues ...
func (t *Tx) DuplicateValues(values ...string) *Tx {
	t.cmd.DuplicateValues(values...)
	return t
}

// SetFields ...
func (t *Tx) SetFields(fields ...string) *Tx {
	t.cmd.SetFields(fields...)
	return t
}

// Set ...
func (t *Tx) Set(assignments ...string) *Tx {
	t.cmd.Set(assignments...)
	return t
}

// Update ...
func (t *Tx) Update(table string, others ...string) *Tx {
	t.cmd.Update(table, others...)
	return t
}

// Delete ...
func (t *Tx) Delete(as ...string) *Tx {
	t.cmd.Delete(as...)
	return t
}

// Replace ...
func (t *Tx) Replace(table string, fields ...string) *Tx {
	t.cmd.Replace(table, fields...)
	return t
}

// Parentheses ...
func (t *Tx) Parentheses(any string) *Tx {
	t.cmd.Parentheses(any)
	return t
}

// Limit ...
func (t *Tx) Limit(count int, offset ...int) *Tx {
	t.cmd.Limit(count, offset...)
	return t
}

// GroupBy ...
func (t *Tx) GroupBy(fields ...string) *Tx {
	t.cmd.GroupBy(fields...)
	return t
}

// OrderBy ...
func (t *Tx) OrderBy(field string, direction ...Direction) *Tx {
	t.cmd.OrderBy(field, direction...)
	return t
}

// Having ...
func (t *Tx) Having(condition string) *Tx {
	t.cmd.Having(condition)
	return t
}

// Raw appends string to command
func (t *Tx) Raw(any string) *Tx {
	t.cmd.Raw(any)
	return t
}

// SQL ...
func (t *Tx) SQL(sqlcmds ...string) string {
	return t.cmd.SQL(sqlcmds...)
}

// PrepareContext ...
func (t *Tx) PrepareContext(ctx context.Context) (*Stmt, error) {
	return WrapStmt(t.x.PrepareContext(ctx, t.SQL()))
}

// Prepare ...
func (t *Tx) Prepare(ctx context.Context) (*Stmt, error) {
	return t.PrepareContext(context.Background())
}

// StmtContext ...
func (t *Tx) StmtContext(ctx context.Context, s *Stmt) *Stmt {
	return &Stmt{x: t.x.StmtContext(ctx, s.x)}
}

// Stmt ...
func (t *Tx) Stmt(s *Stmt) *Stmt {
	return t.StmtContext(context.Background(), s)
}

// ExecContext ...
func (t *Tx) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	if len(args) > 0 {
		return t.x.ExecContext(ctx, t.SQL(), args...)
	}

	return t.x.ExecContext(ctx, t.SQL(), t.args...)
}

// Exec ...
func (t *Tx) Exec(args ...interface{}) (sql.Result, error) {
	return t.ExecContext(context.Background(), args...)
}

// QueryContext ...
func (t *Tx) QueryContext(ctx context.Context, args ...interface{}) (*Rows, error) {
	if len(args) > 0 {
		return WrapRows(t.x.QueryContext(ctx, t.SQL(), args...))
	}
	return WrapRows(t.x.QueryContext(ctx, t.SQL(), t.args...))
}

// Query ...
func (t *Tx) Query(args ...interface{}) (*Rows, error) {
	return t.QueryContext(context.Background(), args...)
}

// AllContext ...
func (t *Tx) AllContext(ctx context.Context, args ...interface{}) ([]interface{}, error) {
	rows, err := t.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	return rows.All()
}

// All ...
func (t *Tx) All(args ...interface{}) ([]interface{}, error) {
	return t.AllContext(context.Background(), args...)
}

// UnmarshalAllContext ...
func (t *Tx) UnmarshalAllContext(ctx context.Context, x interface{}, args ...interface{}) error {
	rows, err := t.QueryContext(ctx, args...)
	if err != nil {
		return err
	}

	return rows.UnmarshalAll(x)
}

// UnmarshalAll ...
func (t *Tx) UnmarshalAll(x interface{}, args ...interface{}) error {
	return t.UnmarshalAllContext(context.Background(), x, args...)
}

// QueryRowContext ...
func (t *Tx) QueryRowContext(ctx context.Context, args ...interface{}) *Row {
	return WrapRow(t.QueryContext(ctx, args...))
}

// QueryRow ...
func (t *Tx) QueryRow(args ...interface{}) *Row {
	return t.QueryRowContext(context.Background(), args...)
}

// DataContext ...
func (t *Tx) DataContext(ctx context.Context, args ...interface{}) (interface{}, error) {
	return t.QueryRowContext(ctx, args...).Data()
}

// Data ...
func (t *Tx) Data(args ...interface{}) (interface{}, error) {
	return t.DataContext(context.Background(), args...)
}

// UnmarshalContext ...
func (t *Tx) UnmarshalContext(ctx context.Context, x interface{}, args ...interface{}) error {
	return t.QueryRowContext(ctx, args...).Unmarshal(x)
}

// Unmarshal ...
func (t *Tx) Unmarshal(x interface{}, args ...interface{}) error {
	return t.UnmarshalContext(context.Background(), x, args...)
}
