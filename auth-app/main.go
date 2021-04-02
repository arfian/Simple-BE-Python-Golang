package main

import (
	"net/http"
	routes "efishery-task/auth-app/routes"
)

func main() {
	http.ListenAndServe(":2001", routes.InitRouter())
}