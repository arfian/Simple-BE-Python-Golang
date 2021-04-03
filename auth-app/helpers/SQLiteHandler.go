package helpers

import (
	"database/sql"
	"fmt"
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

func (handler *SQLiteHandler) Prepare(statement string, field ...interface{}) (*sql.Stmt, error) {
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
	smtm.Exec(field...)
	defer handler.Conn.Close()
	return smtm, nil
}

func (handler *SQLiteHandler) Query(statement string) (*sql.Rows, error) {
	conn, err := getDb()
	handler.Conn = conn
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rows, err := handler.Conn.Query(statement)
	defer handler.Conn.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return rows, nil
}

func (handler *SQLiteHandler) CloseDb() {
	defer handler.Conn.Close()
}