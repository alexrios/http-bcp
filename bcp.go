package main

import (
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func Export(db string, schema string, table string) (string, error) {
	exportLogger := log.WithFields(log.Fields{
		"phase":  "runtime",
		"event":  "Exporting using BCP",
		"db":     db,
		"schema": schema,
		"table":  table,
		"path":   bcpPath,
	})
	outputFile := fmt.Sprintf("%v/%v_%v_%v.txt", bcpPath, db, schema, table)
	target := fmt.Sprintf("/opt/mssql-tools/bin/bcp %v.%v.%v out %v -c -t ';' -S %v -U %v -P %v",
		db, schema, table, outputFile, dbHost, dbUser, dbPw)
	exportLogger.Info("Calling BCP routine")
	return execute(target)
}

func Import(db string, schema string, table string) (string, error) {
	importLogger := log.WithFields(log.Fields{
		"phase":  "runtime",
		"event":  "Exporting using BCP",
		"db":     db,
		"schema": schema,
		"table":  table,
		"path":   bcpPath,
	})

	outputFile := fmt.Sprintf("%v/%v_%v_%v.txt", bcpPath, db, schema, table)
	target := fmt.Sprintf("/opt/mssql-tools/bin/bcp %v.%v.%v in %v -c -t ';' -S %v -U %v -P %v",
		db, schema, table, outputFile, dbHost, dbUser, dbPw)
	importLogger.Info("Calling BCP routine")
	return execute(target)
}

func execute(target string) (string, error) {
	bcpLogger := log.WithFields(log.Fields{
		"phase":  "runtime",
		"event":  "Calling BCP",
		"target": target,
	})
	cmd := exec.Command("bash", "-c", target)
	out, err := cmd.CombinedOutput()
	if err != nil {
		bcpLogger.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error calling BCP")
		return "", errors.New(fmt.Sprintf("BCP failed with %s\n CMD: %v", err, target))
	}
	msg := string(out)
	bcpLogger.WithFields(log.Fields{
		"message": msg,
	}).Info("BCP call completed with no errors")

	return msg, nil
}
