package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os/exec"
)

func Export(db string, schema string, table string) (string, error) {
	outputFile := fmt.Sprintf("bcp_%v_%v_%v.txt", db, schema, table)
	target := fmt.Sprintf("/opt/mssql-tools/bin/bcp %v.%v.%v out %v -c -t ';' -S %v -U %v -P %v",
		db, schema, table, outputFile, dbHost, dbUser, dbPw)

	return execute(target)
}

func Import(db string, schema string, table string) (string, error) {
	outputFile := fmt.Sprintf("bcp_%v_%v_%v.txt", db, schema, table)
	target := fmt.Sprintf("/opt/mssql-tools/bin/bcp %v.%v.%v in %v -c -t ';' -S %v -U %v -P %v",
		db, schema, table, outputFile, dbHost, dbUser, dbPw)

	return execute(target)
}

func execute(target string) (string, error) {
	cmd := exec.Command("bash", "-c", target)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(fmt.Sprintf("bcp failed with %s\n CMD: %v", err, target))
	}
	return string(out), nil
}
