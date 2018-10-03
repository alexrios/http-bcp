package main

import "os"

var dbHost string
var dbUser string
var dbPw string

func main() {
	dbHost = os.Getenv("MSSQL_HOST")
	dbUser = os.Getenv("MSSQL_USER")
	dbPw = os.Getenv("MSSQL_PASSWORD")

	a := App{}
	a.MakeRoutes()
	a.Run(":8080")
}
