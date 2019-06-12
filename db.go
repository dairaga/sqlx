package sqlx

import (
	"context"
	"database/sql"
	"errors"
	"os"
)

// DB ...
type DB struct {
	*sqlxobj
}

// WrapDB ...
func WrapDB(db *sql.DB, err error) (*DB, error) {
	if err != nil {
		return nil, err
	}

	return &DB{&sqlxobj{db, NewCmd(), nil}}, nil
}

var (
	_driver string
	_dsn    string
)

// Close ...
func (db *DB) Close() error {
	return db.sqlxobj.x.(*sql.DB).Close()
}

// Begin ...
func (db *DB) Begin() (*Tx, error) {
	return WrapTx(db.sqlxobj.x.(*sql.DB).Begin())
}

// BeginTx ...
func (db *DB) BeginTx(ctx context.Context, ops *sql.TxOptions) (*Tx, error) {
	return WrapTx(db.sqlxobj.x.(*sql.DB).BeginTx(ctx, ops))
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
