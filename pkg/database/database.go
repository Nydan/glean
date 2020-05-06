package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// A Database interface provides connectivity to DB
type Database interface {
	// sqldb
	Leader() *sqlx.DB
	Follower() *sqlx.DB

	// sqlx only
	Get(destination interface{}, query string, args ...interface{}) error
	Select(destination interface{}, query string, args ...interface{}) error
	Rebind(query string) string
	Named(query string, arg interface{}) (string, interface{}, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Begin() (*sql.Tx, error)
	Beginx() (*sqlx.Tx, error)

	// context query
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)

	PrepareNamedContextLeader(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	PrepareNamedContextFollower(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	PreparexContextFollower(ctx context.Context, query string) (*sqlx.Stmt, error)

	// SelectLeader using leader db to select data. Use this method with cautious
	SelectLeader(destination interface{}, query string, args ...interface{}) error
}
