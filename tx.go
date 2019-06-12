package sqlx

import (
	"context"
	"database/sql"
)

// Tx wraps sql.Tx
type Tx struct {
	*sqlxobj
}

// WrapTx ...
func WrapTx(tx *sql.Tx, err error) (*Tx, error) {
	if err != nil {
		return nil, err
	}

	return &Tx{&sqlxobj{tx, NewCmd(), nil}}, nil
}

// Commit ...
func (t *Tx) Commit() error {
	return t.sqlxobj.x.(*sql.Tx).Commit()
}

// Rollback ...
func (t *Tx) Rollback() error {
	return t.sqlxobj.x.(*sql.Tx).Rollback()
}

// StmtContext ...
func (t *Tx) StmtContext(ctx context.Context, s *Stmt) *Stmt {
	return &Stmt{x: t.sqlxobj.x.(*sql.Tx).StmtContext(ctx, s.x)}
}

// Stmt ...
func (t *Tx) Stmt(s *Stmt) *Stmt {
	return t.StmtContext(context.Background(), s)
}
