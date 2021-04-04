package main

import (
	"net/http"
	"os"
	routes "efishery-task/auth-app/routers"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func createTable() {
	database, _ := sql.Open("sqlite3", "./databases/efishery.db")
    database.Exec("DROP TABLE users")
    statementCreate, _ := database.Prepare("CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(100), phone VARCHAR(30), role VARCHAR(20), password VARCHAR(4), username VARCHAR(100), created_at DATETIME DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME DEFAULT CURRENT_TIMESTAMP)")
    statementCreate.Exec()
    database.Close()
}

func main() {
	os.MkdirAll("./databases", 0755)
    os.Create("./databases/efishery.db")
    createTable()
	http.ListenAndServe(":2001", routes.InitRouter())
}