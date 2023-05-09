package db

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/pkg/errors"

	"github.com/mal-mel/devices_api/internal/configs"
)

type Connector interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

type PGCon struct {
	Connection Connector
}

func ConnectPostgres(ctx context.Context, db *configs.Database) (PGCon, error) {
	maxConnections := strconv.Itoa(db.MaxConnections)
	connString := "postgres://" + os.Getenv(db.UserEnvKey) + ":" + os.Getenv(db.PassEnvKey) +
		"@" + db.Host + ":" + db.Port + "/" + db.DBName +
		"?sslmode=disable&pool_max_conns=" + maxConnections

	ctx, cancel := context.WithTimeout(ctx, time.Duration(db.Timeout)*time.Second)
	_ = cancel // no needed yet

	conn, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return PGCon{}, err
	}

	pgc := PGCon{Connection: conn}

	return pgc, nil
}

func (db *PGCon) StartTransaction(ctx context.Context) (Database, error) {
	tx, err := db.Connection.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &PGCon{Connection: tx}, nil
}

func (db *PGCon) CommitTransaction(ctx context.Context) error {
	tx, ok := db.Connection.(pgx.Tx)
	if !ok {
		return nil
	}

	return tx.Commit(ctx)
}

func (db *PGCon) RollbackTransaction(ctx context.Context) error {
	tx, ok := db.Connection.(pgx.Tx)
	if !ok {
		return nil
	}
	return tx.Rollback(ctx) // nolint:errcheck // rollback err
}

func (db *PGCon) Ping(ctx context.Context) error {
	if pingConn, ok := db.Connection.(interface {
		Ping(ctx context.Context) error
	}); ok {
		return pingConn.Ping(ctx)
	}
	return errors.New("can't ping dn")
}
