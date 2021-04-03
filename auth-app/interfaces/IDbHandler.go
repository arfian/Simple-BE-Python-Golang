package interfaces

import (
	"database/sql"
)

type IDbHandler interface {
	Prepare(statement string, field ...interface{}) (*sql.Stmt, error)
	Query(statement string) (*sql.Rows, error)
	CloseDb()
}