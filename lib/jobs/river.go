// Package jobs provides a minimal fx-lifecycle background job runner.
// Jobs run on fixed intervals; cron expressions are deliberately not
// supported yet.
// ponytail: interval tickers cover current sweeps; swap in a cron parser
// (e.g. robfig/cron) when schedules become irregular or per-env configurable.
package jobs

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"modular-monolith/lib/config"

	"github.com/jmoiron/sqlx"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverdatabasesql"
	"github.com/riverqueue/river/rivermigrate"
	"go.uber.org/fx"
) // Module wires the runner into fx and hooks Start/Stop into the lifecycle.
var RiverModule = fx.Module("jobs",
	fx.Provide(NewRiverClient),
	fx.Invoke(MigrateRiver),
)

func NewRiverClient(cfg *config.Config, DB *sqlx.DB) *river.Client[*sql.Tx] {
	riverClient, err := river.NewClient(riverdatabasesql.New(DB.DB), &river.Config{})
	if err != nil {
		panic(err)
	}
	return riverClient
}

func MigrateRiver(DB *sqlx.DB) {
	ctx := context.Background()
	driver := riverdatabasesql.New(DB.DB)
	migrator, err := rivermigrate.New[*sql.Tx](driver, &rivermigrate.Config{
		// Test schema with no migrations for purposes of this test.
		Schema: "jobs",
	})
	if err != nil {
		panic(err)
	}

	printVersions := func(res *rivermigrate.MigrateResult) {
		for _, version := range res.Versions {
			fmt.Printf(
				"Migrated [%s] version %d\n",
				strings.ToUpper(string(res.Direction)),
				version.Version,
			)
		}
	}

	// Migrate to version 3. An actual call may want to omit all MigrateOpts,
	// which will default to applying all available up migrations.
	res, err := migrator.Migrate(ctx, rivermigrate.DirectionUp, &rivermigrate.MigrateOpts{
		TargetVersion: 3,
	})
	if err != nil {
		panic(err)
	}
	printVersions(res)

	// Migrate down by three steps. Down migrating defaults to running only one
	// step unless overridden by an option like MaxSteps or TargetVersion.
	res, err = migrator.Migrate(ctx, rivermigrate.DirectionDown, &rivermigrate.MigrateOpts{
		MaxSteps: 3,
	})
	if err != nil {
		panic(err)
	}
	printVersions(res)
}
