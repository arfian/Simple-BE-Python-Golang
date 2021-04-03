package helpers

import (
	"database/sql"
	"fmt"
	"efishery-task/auth-app/interfaces"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteHandler struct {
	Conn *sql.DB
}

func getDb() (*sql.DB, error){
	sqlConn, err := sql.Open("sqlite3", "../databases/efishery.db")

	if err != nil {
		fmt.Println(err)
	}
	return sqlConn, nil
}

func (handler *SQLiteHandler) Prepare(statement string) (*sql.Stmt, error) {
	conn, err := getDb()
	handler.Conn = conn
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	smtm, err :=  handler.Conn.Prepare(statement)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	return smtm, nil
}

func (handler *SQLiteHandler) Query(statement string) (interfaces.IRow, error) {
	conn, err := getDb()
	handler.Conn = conn
	defer handler.Conn.Close()
	if err != nil {
		fmt.Println(err)
		return new(SqliteRow),err
	}

	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(SqliteRow),err
	}
	row := new(SqliteRow)
	row.Rows = rows

	return row, nil
}

func (handler *SQLiteHandler) CloseDb() {
	defer handler.Conn.Close()
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