package interfaces

import (
	"database/sql"
	"github.com/irahardianto/service-pattern-go/interfaces"
)

type IDbHandler interface {
	Prepare(statement string) (*sql.Stmt, error)
	Query(statement string) (IRow, error)
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}