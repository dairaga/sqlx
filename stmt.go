package sqlx

import (
	"context"
	"database/sql"
)

// Stmt wraps sql.Stmt
type Stmt struct {
	x *sql.Stmt
}

// WrapStmt ...
func WrapStmt(stmt *sql.Stmt, err error) (*Stmt, error) {
	if err != nil {
		return nil, err
	}

	return &Stmt{x: stmt}, nil
}

// Close ...
func (stmt *Stmt) Close() error {
	return stmt.x.Close()
}

// ExecContext ...
func (stmt *Stmt) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	return stmt.x.ExecContext(ctx, args...)
}

// Exec ...
func (stmt *Stmt) Exec(args ...interface{}) (sql.Result, error) {
	return stmt.ExecContext(context.Background(), args...)
}

// QueryContext ...
func (stmt *Stmt) QueryContext(ctx context.Context, args ...interface{}) (*Rows, error) {
	return WrapRows(stmt.x.QueryContext(ctx, args...))
}

// Query ...
func (stmt *Stmt) Query(args ...interface{}) (*Rows, error) {
	return stmt.QueryContext(context.Background(), args...)
}

// AllContext ...
func (stmt *Stmt) AllContext(ctx context.Context, args ...interface{}) ([]interface{}, error) {
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	return rows.All()
}

// All ...
func (stmt *Stmt) All(args ...interface{}) ([]interface{}, error) {
	return stmt.AllContext(context.Background(), args...)
}

// UnmarshalAllContext ...
func (stmt *Stmt) UnmarshalAllContext(ctx context.Context, x interface{}, args ...interface{}) error {
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return err
	}

	return rows.UnmarshalAll(x)
}

// UnmarshalAll ...
func (stmt *Stmt) UnmarshalAll(x interface{}, args ...interface{}) error {
	return stmt.UnmarshalAllContext(context.Background(), x, args...)
}

// QueryRowContext ...
func (stmt *Stmt) QueryRowContext(ctx context.Context, args ...interface{}) *Row {
	return WrapRow(stmt.QueryContext(ctx, args...))
}

// QueryRow ...
func (stmt *Stmt) QueryRow(args ...interface{}) *Row {
	return stmt.QueryRowContext(context.Background(), args...)
}

// DataContext ...
func (stmt *Stmt) DataContext(ctx context.Context, args ...interface{}) (interface{}, error) {
	return stmt.QueryRowContext(ctx, args...).Data()
}

// Data ...
func (stmt *Stmt) Data(args ...interface{}) (interface{}, error) {
	return stmt.DataContext(context.Background(), args...)
}

// UnmarshalContext ...
func (stmt *Stmt) UnmarshalContext(ctx context.Context, x interface{}, args ...interface{}) error {
	return stmt.QueryRowContext(ctx, args...).Unmarshal(x)
}

// Unmarshal ...
func (stmt *Stmt) Unmarshal(x interface{}, args ...interface{}) error {
	return stmt.UnmarshalContext(context.Background(), x, args...)
}
