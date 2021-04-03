package interfaces

import (
	"database/sql"
)

type IDbHandler interface {
	Prepare(statement string) (*sql.Stmt, error)
	Query(statement string) (IRow, error)
	CloseDb()
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}