package db

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type IManager interface {
	Run(ctx context.Context, fn func(context.Context) error) error
}
type Beginner interface {
	Begin(ctx context.Context) (*sqlx.Tx, error)
}

type contextKey struct{}

type Manager struct {
	db *sqlx.DB
}

var _ IManager = (*Manager)(nil)

func New(db *sqlx.DB) *Manager {
	return &Manager{
		db: db,
	}
}

type DBTX interface {
	ExecContext(
		context.Context,
		string,
		...any,
	) (sql.Result, error)
	QueryContext(
		context.Context,
		string,
		...any,
	) (*sql.Rows, error)

	QueryRowxContext(
		context.Context,
		string,
		...any,
	) *sqlx.Row

	GetContext(
		context.Context,
		any,
		string,
		...any,
	) error

	SelectContext(
		context.Context,
		any,
		string,
		...any,
	) error
}

func (m *Manager) Run(
	ctx context.Context,
	fn func(context.Context) error,
) error {
	// Already inside a transaction.
	if _, ok := fromContext(ctx); ok {
		return fn(ctx)
	}

	tx, err := m.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	ctx = withTx(ctx, tx)

	if err := fn(ctx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func withTx(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, contextKey{}, tx)
}

func fromContext(ctx context.Context) (*sqlx.Tx, bool) {
	tx, ok := ctx.Value(contextKey{}).(*sqlx.Tx)
	return tx, ok
}
