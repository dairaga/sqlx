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

// Find ...
func (db *DB) Find(query string, args ...interface{}) (*Rows, error) {
	return WrapRows(db.DB.Query(query, args...))
}

// FindContext ...
func (db *DB) FindContext(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	return WrapRows(db.DB.QueryContext(ctx, query, args...))
}

// AllContext ...
func (db *DB) AllContext(ctx context.Context, query string, args ...interface{}) ([]interface{}, error) {
	rs, err := db.FindContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// All ...
func (db *DB) All(query string, args ...interface{}) ([]interface{}, error) {
	return db.AllContext(context.Background(), query, args...)
}

// UnmarshalAllContext ...
func (db *DB) UnmarshalAllContext(ctx context.Context, x interface{}, query string, args ...interface{}) error {
	rs, err := db.FindContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return rs.UnmarshalAll(x)
}

// UnmarshalAll ...
func (db *DB) UnmarshalAll(x interface{}, query string, args ...interface{}) error {
	return db.UnmarshalAllContext(context.Background(), x, query, args...)
}

// FindRowContext ...
func (db *DB) FindRowContext(ctx context.Context, query string, args ...interface{}) *Row {
	rs, err := db.FindContext(ctx, query, args...)
	r := &Row{rows: rs, err: err}
	if err == nil {
		rs.Next()
		r.scan()
		rs.Close()
	}
	return r
}

// FindRow ...
func (db *DB) FindRow(query string, args ...interface{}) *Row {
	return db.FindRowContext(context.Background(), query, args...)
}

// DataContext ...
func (db *DB) DataContext(ctx context.Context, query string, args ...interface{}) (interface{}, error) {
	r := db.FindRowContext(ctx, query, args...)
	return r.Data()
}

// Data ...
func (db *DB) Data(query string, args ...interface{}) (interface{}, error) {
	return db.DataContext(context.Background(), query, args...)
}

// UnmarshalContext ...
func (db *DB) UnmarshalContext(ctx context.Context, x interface{}, query string, args ...interface{}) error {
	r := db.FindRowContext(ctx, query, args...)
	return r.Unmarshal(x)
}

// Unmarshal ...
func (db *DB) Unmarshal(x interface{}, query string, args ...interface{}) error {
	return db.UnmarshalContext(context.Background(), x, query, args...)
}
