package sqldb

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// GetContext does a QueryRow using the provided Queryer, and scans the
// resulting row to dest.  If dest is scannable, the result must only have one
// column. Otherwise, StructScan is used.  Get will return sql.ErrNoRows like
// row.Scan would. Any placeholder parameters are replaced with supplied args.
// An error is returned if the result set is empty.
func (db *db) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return db.follower.GetContext(ctx, dest, query, args...)
}

// SelectContext executes a query using the provided Queryer, and StructScans
// each row into dest, which must be a slice.  If the slice elements are
// scannable, then the result set must have only one column.  Otherwise,
// StructScan is used. The *sql.Rows are closed automatically.
// Any placeholder parameters are replaced with supplied args.
func (db *db) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return db.follower.SelectContext(ctx, dest, query, args...)
}

// QueryContext queries the database and returns an *sqlx.Rows.
// Any placeholder parameters are replaced with supplied args.
func (db *db) QueryContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return db.follower.QueryxContext(ctx, query, args...)
}

// QueryRowContext queries the database and returns an *sqlx.Row. Any placeholder parameters are replaced with supplied args.
func (db *db) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return db.follower.QueryRowxContext(ctx, query, args...)
}

// ExecContext executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (db *db) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.leader.ExecContext(ctx, query, args...)
}

// NamedExecContext using this db.
// Any named placeholder parameters are replaced with fields from arg.
func (db *db) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return db.leader.NamedExecContext(ctx, query, arg)
}
