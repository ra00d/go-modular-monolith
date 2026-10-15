package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"modular-monolith/lib/config"
	"modular-monolith/lib/logger"
	"modular-monolith/lib/utils"

	"github.com/google/uuid"
	"github.com/pressly/goose/v3"
	"github.com/shopspring/decimal"

	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/generator/postgres"
	"github.com/go-jet/jet/v2/generator/template"
	postgres2 "github.com/go-jet/jet/v2/postgres"
)

func MigrateAndGenerateTypes(migrations embed.FS) any {
	genDir, err := utils.CurrentDirPath(2)
	if err != nil {
		log.Fatal(err)
	}

	schema := filepath.Base(genDir)
	return func(cfg *config.Config) {
		dsn, err := cfg.Database.DSNWithSchema(schema)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}

		migrateAndGenerateTypes(schema, genDir, dsn, migrations)
	}
}

// MigrateAndGenerateTypes function    migrate the database and generate models and sql-builders using jet
// This operation is working per schema
func migrateAndGenerateTypes(module string, genPath string, dsn string, FS embed.FS) {
	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = migrateModule(DB, module, FS)
	if err != nil {
		log.Fatal(err)
	}

	err = generateTypes(DB, module, genPath)
	if err != nil {
		log.Fatal(err)
	}
}

func migrateModule(db *sql.DB, module string, FS embed.FS) error {
	// Create the schema first.
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if _, err := tx.Exec(fmt.Sprintf(
		`CREATE SCHEMA IF NOT EXISTS "%s"`,
		module,
	)); err != nil {
		return err
	}
	// Use a dedicated connection.
	// conn, err := db.Conn(context.Background())
	// if err != nil {
	// 	return err
	// }
	// defer conn.Close()
	//
	// if _, err := conn.ExecContext(
	// 	context.Background(),
	// 	fmt.Sprintf(`SET search_path TO "%s"`, module),
	// ); err != nil {
	// 	return err
	// }
	ctx := context.Background()
	if _, err := db.ExecContext(
		ctx,
		fmt.Sprintf(`SET search_path TO "%s"`, module),
	); err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	provider, err := goose.NewProvider(
		goose.DialectPostgres,
		db,
		FS,
		goose.WithTableName(module+".goose_db_version"),
	)
	if err != nil {
		return err
	}

	res, err := provider.Up(
		ctx,
	)
	if err != nil {
		return err
	}
	for _, v := range res {
		msg := fmt.Sprintf("Migration %s has been applied", v.Source.Path)
		logger.Debug(msg)
	}
	sRes, err := provider.Status(ctx)
	for _, v := range sRes {
		msg := fmt.Sprintf(
			"Migration %s has been %s at %s",
			v.Source.Path,
			v.State,
			utils.HumanTime(v.AppliedAt),
		)
		logger.Debug(msg)
	}
	return err
}

func generateTypes(db *sql.DB, schema, path string) error {
	err := postgres.GenerateDB(
		db, schema, path+"/infra/.gen",
		template.Default(postgres2.Dialect).
			UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
				return template.DefaultSchema(schemaMetaData).
					UseModel(
						template.DefaultModel().
							UsePath("../models").
							UseTable(func(table metadata.Table) template.TableModel {
								return template.DefaultTableModel(table).
									UseField(func(columnMetaData metadata.Column) template.TableModelField {
										defaultTableModelField := template.DefaultTableModelField(
											columnMetaData,
										)

										switch columnMetaData.DataType.Name {
										case "jsonb":
											defaultTableModelField.Type = template.NewType(
												&JSONType{},
											)
										}
										return defaultTableModelField
									})
							}),
					).
					UseSQLBuilder(
						template.DefaultSQLBuilder().
							UsePath(".."),
					)
			}),
	) // or GenerateDSN(...)
	if err != nil {
		return err
	}
	return nil
}

func NewNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func PtrToNull(p *string) sql.NullString {
	if p == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *p, Valid: true}
}

func NullToPtr(n sql.NullString) *string {
	if !n.Valid {
		return nil
	}
	v := n.String
	return &v
}

func NullTimeToPtr(n sql.NullTime) *time.Time {
	if !n.Valid {
		return nil
	}
	v := n.Time
	return &v
}

func PtrTimeToNull(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

func PtrUUIDToNull(p *uuid.UUID) uuid.NullUUID {
	if p == nil {
		return uuid.NullUUID{}
	}
	return uuid.NullUUID{UUID: *p, Valid: true}
}

func NullUUIDToPtr(n uuid.NullUUID) *uuid.UUID {
	if !n.Valid {
		return nil
	}
	v := n.UUID
	return &v
}

// NullStringFromPtr is an alias for PtrToNull. Exists so callers in
// controller mappers and infra repositories can use the same name when
// they're already importing lib/db.
func NullStringFromPtr(p *string) sql.NullString { return PtrToNull(p) }

// NullTimeFromPtr is an alias for PtrTimeToNull.
func NullTimeFromPtr(t *time.Time) sql.NullTime { return PtrTimeToNull(t) }

// NullUUIDFromPtr is an alias for PtrUUIDToNull.
func NullUUIDFromPtr(p *uuid.UUID) uuid.NullUUID { return PtrUUIDToNull(p) }

// NullStringFromDecimalPtr converts a *decimal.Decimal to a sql.NullString
// holding its String() form. Returns the zero NullString when p is nil.
func NullStringFromDecimalPtr(p *decimal.Decimal) sql.NullString {
	if p == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: p.String(), Valid: true}
}

// MustDecimal parses a numeric string into a decimal.Decimal. Returns
// zero on parse error so a malformed column does not crash a row scan.
func MustDecimal(s string) decimal.Decimal {
	if s == "" {
		return decimal.Zero
	}
	d, err := decimal.NewFromString(s)
	if err != nil {
		return decimal.Zero
	}
	return d
}

// NullTimeFromTime is a convenience constructor: returns the zero
// NullTime when t is the zero value, otherwise wraps it.
func NullTimeFromTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: t, Valid: true}
}
