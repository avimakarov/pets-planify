package internal

import (
	"context"
	"database/sql"
)

//go:generate mockgen --source=deps.go --destination=deps_test.go --package=${GOPACKAGE}_test

type Tx interface {
	Commit() error
	Rollback() error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}
