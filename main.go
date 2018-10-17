package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var dbHost string
var dbUser string
var dbPw string

var callBackExportUrl string
var callBackImportUrl string

var bcpPath string

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	bootPhaseLogger := log.WithFields(log.Fields{
		"phase": "boot",
	})
	bootPhaseLogger.Info("Starting App...")
	readEnvVars()
	a := App{}
	a.MakeRoutes()
	a.Run(":8080")
}

func readEnvVars() {
	bootPhaseLogger := log.WithFields(log.Fields{
		"phase": "boot",
	})
	dbHost = os.Getenv("MSSQL_HOST")
	bootPhaseLogger.WithFields(log.Fields{
		"ENV_VAR":   "MSSQL_HOST",
		"ENV_VALUE": dbHost,
	}).Info("Gathering ENV")

	dbUser = os.Getenv("MSSQL_USER")
	bootPhaseLogger.WithFields(log.Fields{
		"ENV_VAR":   "MSSQL_USER",
		"ENV_VALUE": dbHost,
	}).Info("Gathering ENV")

	dbPw = os.Getenv("MSSQL_PASSWORD")
	var passwordForLog string
	if len(dbPw) == 0 {
		passwordForLog = "Not Shown but filled"
	} else {
		passwordForLog = "Empty string"
	}
	bootPhaseLogger.WithFields(log.Fields{
		"ENV_VAR":   dbHost,
		"ENV_VALUE": passwordForLog,
	}).Info("Gathering ENV")

	callBackExportUrl = os.Getenv("EXPORT_CALLBACK_URL")
	bootPhaseLogger.WithFields(log.Fields{
		"ENV_VAR":   "EXPORT_CALLBACK_URL",
		"ENV_VALUE": callBackExportUrl,
	}).Info("Gathering ENV")

	callBackImportUrl = os.Getenv("IMPORT_CALLBACK_URL")
	bootPhaseLogger.WithFields(log.Fields{
		"ENV_VAR":   "IMPORT_CALLBACK_URL",
		"ENV_VALUE": callBackImportUrl,
	}).Info("Gathering ENV")

	bcpPath = os.Getenv("BCP_PATH")
	bootPhaseLogger.WithFields(log.Fields{
		"ENV_VAR":   "BCP_PATH",
		"ENV_VALUE": bcpPath,
	}).Info("Gathering ENV")
}
