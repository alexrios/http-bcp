package main

import "os"

var dbHost string
var dbUser string
var dbPw string

var callBackExportUrl string
var callBackImportUrl string

func main() {
	readEnvVars()
	a := App{}
	a.MakeRoutes()
	a.Run(":8080")
}

func readEnvVars() {
	dbHost = os.Getenv("MSSQL_HOST")
	dbUser = os.Getenv("MSSQL_USER")
	dbPw = os.Getenv("MSSQL_PASSWORD")
	callBackExportUrl = os.Getenv("EXPORT_CALLBACK_URL")
	callBackImportUrl = os.Getenv("IMPORT_CALLBACK_URL")
}
