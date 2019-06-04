package sqlx

import (
	"context"
	"database/sql"
)

// DB ...
type DB struct {
	*sql.DB
}

// WrapDB ...
func WrapDB(db *sql.DB, err error) (*DB, error) {
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// XQuery ...
func (db *DB) XQuery(query string, args ...interface{}) (*Rows, error) {
	return Wrap(db.DB.Query(query, args...))
}

// XQueryContext ...
func (db *DB) XQueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return Wrap(db.DB.QueryContext(ctx, query, args...))
}

// XQueryAllContext ...
func (db *DB) XQueryAllContext(ctx context.Context, query string, args ...interface{}) ([]interface{}, error) {
	rs, err := db.XQueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// UnmarshalAllContext ...
func (db *DB) UnmarshalAllContext(ctx context.Context, x interface{}, query string, args ...interface{}) error {
	rs, err := db.XQueryContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return rs.UnmarshalAll(x)
}

// UnmarshalAll ...
func (db *DB) UnmarshalAll(x interface{}, query string, args ...interface{}) error {
	return db.UnmarshalAllContext(context.Background(), x, query, args...)
}

// XQueryAll ...
func (db *DB) XQueryAll(query string, args ...interface{}) ([]interface{}, error) {
	return db.XQueryAllContext(context.Background(), query, args...)
}

// XQueryRowContext ...
func (db *DB) XQueryRowContext(ctx context.Context, query string, args ...interface{}) *Row {
	rs, err := db.XQueryContext(ctx, query, args...)
	r := &Row{rows: rs, err: err}
	if err == nil {
		rs.Next()
		r.scan()
	}
	return r
}

// XQueryRow ...
func (db *DB) XQueryRow(query string, args ...interface{}) *Row {
	return db.XQueryRowContext(context.Background(), query, args...)
}

// FindContext ...
func (db *DB) FindContext(ctx context.Context, query string, args ...interface{}) (interface{}, error) {
	r := db.XQueryRowContext(ctx, query, args...)
	return r.Data()
}

// Find ...
func (db *DB) Find(query string, args ...interface{}) (interface{}, error) {
	return db.FindContext(context.Background(), query, args...)
}

// UnmarshalContext ...
func (db *DB) UnmarshalContext(ctx context.Context, x interface{}, query string, args ...interface{}) error {
	r := db.XQueryRowContext(ctx, query, args...)
	return r.Unmarshal(x)
}

// Unmarshal ...
func (db *DB) Unmarshal(x interface{}, query string, args ...interface{}) error {
	return db.UnmarshalContext(context.Background(), x, query, args...)
}
