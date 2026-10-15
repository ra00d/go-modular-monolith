package db

import (
	"context"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/jmoiron/sqlx"
)

type BaseRepository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) BaseRepository {
	return BaseRepository{
		db: db,
	}
}

func (r *BaseRepository) DB(ctx context.Context) DBTX {
	if tx, ok := fromContext(ctx); ok {
		return tx
	}

	return r.db
}

func (r BaseRepository) QueryOne(
	ctx context.Context,
	stmt pg.SelectStatement,
	row any,
) error {
	if err := stmt.QueryContext(ctx, r.DB(ctx), row); err != nil {
		if IsNoRowError(err) {
			return err
		}
		return err
	}
	return nil
}

func (r BaseRepository) QueryMany(
	ctx context.Context,
	stmt pg.SelectStatement,
	row any,
) error {
	if err := stmt.QueryContext(ctx, r.DB(ctx), row); err != nil {
		if IsNoRowError(err) {
			return err
		}
		return err
	}
	return nil
}

// Paginate fetches one page of stmt and the total number of rows stmt matches,
// ignoring limit and offset. stmt must carry its own filters, they are reused
// for the count so both queries always agree.
func (r BaseRepository) Paginate[T any](
	ctx context.Context,
	stmt pg.SelectStatement,
	limit, offset int,
) (T, int, error) {
	var page T
	var totals []struct{ Count int }
	if err := pg.SELECT(pg.COUNT(pg.STAR)).FROM(stmt.AsTable("_count")).
		QueryContext(ctx, r.DB(ctx), &totals); err != nil {
		return page, 0, err
	}
	if err := r.QueryMany(ctx, stmt.LIMIT(int64(limit)).OFFSET(int64(offset)), &page); err != nil {
		return page, 0, err
	}
	if len(totals) == 0 {
		return page, 0, nil
	}
	return page, totals[0].Count, nil
}
