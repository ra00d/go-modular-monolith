// Package db provides
package db

import (
	"log"

	"modular-monolith/lib/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

var DBModule = fx.Module("Database", fx.Provide(NewDB), fx.Provide(New))

func NewDB(cfg *config.Config) *sqlx.DB {
	connStr, err := cfg.Database.DSN()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection is successful
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	//
	// postgres.SetQueryLogger(func(ctx context.Context, queryInfo postgres.QueryInfo) {
	// 	fmt.Printf("- Debug SQL: %s \n", queryInfo.Statement.DebugSql())
	//
	// 	// Depending on how the statement is executed, RowsProcessed is:
	// 	//   - Number of rows returned for Query() and QueryContext() methods
	// 	//   - RowsAffected() for Exec() and ExecContext() methods
	// 	//   - Always 0 for Rows() method.
	// 	fmt.Printf("- Rows processed: %d\n", queryInfo.RowsProcessed)
	// 	fmt.Printf("- Duration %s\n", queryInfo.Duration.String())
	// 	fmt.Printf("- Execution error: %v\n", queryInfo.Err)
	//
	// 	_, _, callerFunction := queryInfo.Caller()
	// 	fmt.Printf(
	// 		"-   function: %s\n",
	// 		callerFunction,
	// 	)
	// })
	//
	return db
}
