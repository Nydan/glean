package sqldb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	// PostgreSQL
	_ "github.com/lib/pq"
	"github.com/nydan/glean/pkg/database"

	"github.com/jmoiron/sqlx"
)

type (
	// db is wrapper for master and slave database connection
	db struct {
		driver   string
		follower *sqlx.DB
		leader   *sqlx.DB
	}

	// ConnectionOptions list of option to connect to database
	ConnectionOptions struct {
		Retry                 int
		MaxOpenConnections    int
		MaxIdleConnections    int
		ConnectionMaxLifetime time.Duration
	}
)

func connectWithRetry(ctx context.Context, driver, dataSourceName string, retry int) (*sqlx.DB, error) {
	var (
		db  *sqlx.DB
		err error
	)

	for t := 0; t <= retry; t++ {
		db, err = sqlx.ConnectContext(ctx, driver, dataSourceName)
		if err != nil {
			time.Sleep(time.Second * 3)
		} else {
			break
		}
	}

	return db, err
}

// Connect to a database
func Connect(ctx context.Context, driver, dataSourceName string, conOpts *ConnectionOptions) (*sqlx.DB, error) {
	opts := conOpts
	if opts == nil {
		opts = &ConnectionOptions{}
	}

	db, err := connectWithRetry(ctx, driver, dataSourceName, opts.Retry)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(opts.MaxOpenConnections)
	db.SetMaxIdleConns(opts.MaxIdleConnections)
	db.SetConnMaxLifetime(opts.ConnectionMaxLifetime)
	return db, nil
}

// Wrap leader and follower sqlx object to one db object
// this is for easier usage, so user doesn't have to specify leader or follower
// all exec is going to leader, all query is going to follower
func Wrap(ctx context.Context, leader, follower *sqlx.DB) (database.Database, error) {
	if leader.DriverName() != follower.DriverName() {
		return nil, fmt.Errorf("sqldb: leader and follower driver is not matched. leader = %s follower = %s", leader.DriverName(), follower.DriverName())
	}

	db := db{
		driver:   leader.DriverName(),
		leader:   leader,
		follower: follower,
	}
	return &db, nil
}

// Close all database connection
func (db *db) Close() error {
	err := db.leader.Close()
	if err != nil {
		return err
	}
	err = db.follower.Close()
	if err != nil {
		return err
	}
	return nil
}

// Leader return leader database connection
func (db *db) Leader() *sqlx.DB {
	return db.leader
}

// Follower return Follower database connection
func (db *db) Follower() *sqlx.DB {
	return db.follower
}

// Get return one value in destination interface.
// It will return error when no value returned.
func (db *db) Get(destination interface{}, query string, args ...interface{}) error {
	return db.follower.Get(destination, query, args...)
}

// Select return more than one value in destination using reflection.
func (db *db) Select(destination interface{}, query string, args ...interface{}) error {
	return db.follower.Select(destination, query, args...)
}

// SelectLeader return more than one value in destination using reflection.
func (db *db) SelectLeader(destination interface{}, query string, args ...interface{}) error {
	return db.leader.Select(destination, query, args...)
}

// Query database and return *sql.Rows
func (db *db) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.follower.Query(query, args...)
}

// Queryx queries the database and returns an *sqlx.Rows.
// Any placeholder parameters are replaced with supplied args.
func (db *db) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return db.follower.Queryx(query, args...)
}

// QueryRow expecting to return at least one *sql.Row
func (db *db) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.follower.QueryRow(query, args...)
}

// QueryRowx expecting to return at least one row
func (db *db) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return db.follower.QueryRowx(query, args...)
}

// NamedQuery using this db.
// Any named placeholder parameters are replaced with fields from arg.
func (db *db) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	return db.follower.NamedQuery(query, arg)
}

// Exec executes query without returning rows.
func (db *db) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.leader.Exec(query, args...)
}

// NamedExec uses BindStruct to get a query executable by the driver and
// then runs Exec on the result.  Returns an error from the binding
// or the query excution itself.
func (db *db) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return db.leader.NamedExec(query, arg)
}

// Begin return sql transaction object
func (db *db) Begin() (*sql.Tx, error) {
	return db.leader.Begin()
}

// Beginx return sqlx transaction object
func (db *db) Beginx() (*sqlx.Tx, error) {
	return db.leader.Beginx()
}

// Rebind a query to targeted bind type
func (db *db) Rebind(query string) string {
	return sqlx.Rebind(sqlx.BindType(db.driver), query)
}

// Named takes a query using named parameters and an argument and
// returns a new query with a list of args that can be executed by a database.
func (db *db) Named(query string, arg interface{}) (string, interface{}, error) {
	return sqlx.Named(query, arg)
}

// PrepareNamedContextLeader returns an sqlx.NamedStmt
func (db *db) PrepareNamedContextLeader(ctx context.Context, query string) (*sqlx.NamedStmt, error) {
	return db.leader.PrepareNamedContext(ctx, db.leader.Rebind(query))
}

// PrepareNamedContextFollower returns an sqlx.NamedStmt
func (db *db) PrepareNamedContextFollower(ctx context.Context, query string) (*sqlx.NamedStmt, error) {
	return db.follower.PrepareNamedContext(ctx, db.follower.Rebind(query))
}

// PreparexContextFollower returns an sqlx.Stmt
func (db *db) PreparexContextFollower(ctx context.Context, query string) (*sqlx.Stmt, error) {
	return db.follower.PreparexContext(ctx, query)
}
