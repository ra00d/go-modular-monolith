package db

import (
	"context"
	"os"
	"testing"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// paginateTable mirrors a table with a grouping column so the test can prove
// filters reach both the count and the page query.
type paginateRow struct {
	ID  int
	Grp string
}

var (
	paginateID  = pg.IntegerColumn("id")
	paginateGrp = pg.StringColumn("grp")
	// jet matches scanned columns by "<table alias>.<field>", so the alias has
	// to line up with the destination type name.
	paginateTbl = pg.NewTable("public", "db_paginate_test", "paginateRow", paginateID, paginateGrp)
)

// openTestDB connects to the local postgres, skipping when it is unreachable so
// CI without infra stays green.
func openTestDB(t *testing.T) *sqlx.DB {
	t.Helper()
	dsn := os.Getenv("TEST_PG_DSN")
	if dsn == "" {
		dsn = "postgres://" + env("TEST_PG_USER", "postgres") + ":" + env("TEST_PG_PASSWORD", "secret") +
			"@" + env("TEST_PG_HOST", "localhost") + ":5433/" + env("TEST_PG_DB", "erp") + "?sslmode=disable"
	}
	database, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		t.Skipf("postgres unreachable: %v", err)
	}
	t.Cleanup(func() { _ = database.Close() })

	if _, err := database.Exec("DROP TABLE IF EXISTS db_paginate_test"); err != nil {
		t.Fatalf("drop table: %v", err)
	}
	if _, err := database.Exec("CREATE TABLE db_paginate_test (id int, grp text)"); err != nil {
		t.Fatalf("create table: %v", err)
	}
	t.Cleanup(func() { _, _ = database.Exec("DROP TABLE db_paginate_test") })
	return database
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// seed replaces the table contents, numbering rows 1..n in insertion order.
func seed(t *testing.T, database *sqlx.DB, groups ...string) {
	t.Helper()
	if _, err := database.Exec("TRUNCATE db_paginate_test"); err != nil {
		t.Fatalf("truncate: %v", err)
	}
	for i, grp := range groups {
		if _, err := database.Exec(
			"INSERT INTO db_paginate_test (id, grp) VALUES ($1, $2)", i+1, grp,
		); err != nil {
			t.Fatalf("seed: %v", err)
		}
	}
}

func TestPaginate(t *testing.T) {
	database := openTestDB(t)
	repo := NewRepo(database)
	ctx := context.Background()

	tests := []struct {
		name      string
		seed      []string
		limit     int
		offset    int
		filter    pg.BoolExpression
		wantIDs   []int
		wantCount int
	}{
		{
			name:      "count is the total, not the page size",
			seed:      []string{"a", "a", "a", "a", "a"},
			limit:     2,
			wantIDs:   []int{1, 2},
			wantCount: 5,
		},
		{
			name:      "offset walks the pages",
			seed:      []string{"a", "a", "a", "a", "a"},
			limit:     2,
			offset:    4,
			wantIDs:   []int{5},
			wantCount: 5,
		},
		{
			name:      "filters narrow both the count and the rows",
			seed:      []string{"a", "b", "b"},
			limit:     10,
			filter:    paginateGrp.EQ(pg.String("b")),
			wantIDs:   []int{2, 3},
			wantCount: 2,
		},
		{
			name:      "offset past the end returns no rows but keeps the total",
			seed:      []string{"a", "a"},
			limit:     10,
			offset:    50,
			wantIDs:   nil,
			wantCount: 2,
		},
		{
			name:      "empty table",
			seed:      nil,
			limit:     10,
			wantIDs:   nil,
			wantCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seed(t, database, tt.seed...)

			stmt := paginateTbl.SELECT(paginateID)
			if tt.filter != nil {
				stmt = stmt.WHERE(tt.filter)
			}

			rows, count, err := repo.Paginate[[]paginateRow](ctx, stmt, tt.limit, tt.offset)
			if err != nil {
				t.Fatalf("Paginate: %v", err)
			}
			if count != tt.wantCount {
				t.Errorf("count = %d, want %d", count, tt.wantCount)
			}
			if len(rows) != len(tt.wantIDs) {
				t.Fatalf("len(rows) = %d, want %d (%v)", len(rows), len(tt.wantIDs), rows)
			}
			for i, want := range tt.wantIDs {
				if rows[i].ID != want {
					t.Errorf("row %d id = %d, want %d", i, rows[i].ID, want)
				}
			}
		})
	}
}
