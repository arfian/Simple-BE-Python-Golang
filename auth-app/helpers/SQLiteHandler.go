package helpers

import (
	"database/sql"
	"fmt"
	"efishery-task/auth-app/interfaces"
)

type SQLiteHandler struct {
	Conn *sql.DB
}

func getDb() (*sql.DB, error){
	sqlConn, _ := sql.Open("sqlite3", "../databases/efishery.db")
	return sqlConn
}

func (handler *SQLiteHandler) Prepare(statement string) (*sql.Stmt, error) {
	handler.Conn, err := getDb()
	if err != nil {
		fmt.Println(err)
		return new(SqliteRow),err
	}
	defer handler.Conn.Close()
	smtm, err :=  handler.Conn.Prepare(statement)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	return smtm, nil
}

func (handler *SQLiteHandler) Query(statement string) (interfaces.IRow, error) {
	handler.Conn, err := getDb()
	if err != nil {
		fmt.Println(err)
		return new(SqliteRow),err
	}
	defer handler.Conn.Close()

	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(SqliteRow),err
	}
	row := new(SqliteRow)
	row.Rows = rows

	return row, nil
}

type SqliteRow struct {
	Rows *sql.Rows
}

func (r SqliteRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return  nil
}

func (r SqliteRow) Next() bool {
	return r.Rows.Next()
}