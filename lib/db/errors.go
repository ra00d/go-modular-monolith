package db

import (
	"database/sql"
	"errors"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/lib/pq"
)

func IsNoRowError(err error) bool {
	return errors.Is(err, qrm.ErrNoRows) || errors.Is(err, sql.ErrNoRows)
}

// IsUniqueViolation reports whether err is a postgres unique constraint
// violation (SQLSTATE 23505). Both registered drivers are checked so callers
// do not have to know which one the connection was opened with.
func IsUniqueViolation(err error) bool {
	const uniqueViolation = "23505"

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == uniqueViolation
	}

	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code == uniqueViolation
	}

	return false
}
