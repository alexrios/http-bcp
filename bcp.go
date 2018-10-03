package main

import (
	"fmt"
	"os/exec"
)

func Export(db string, table string) string {
	outputFile := fmt.Sprintf("bcp_%v_%v.dat", db, table)
	target := fmt.Sprintf("/opt/mssql-tools/bin/bcp %v.dbo.%v out %v -c -t ';' -S %v -U %v -P %v",
		db, table, outputFile, dbHost, dbUser, dbPw)

	cmd := exec.Command("bash", "-c", target)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("bcp failed with %s\n CMD: %v", err, target)
	}
	return string(out)
}

func Import(db string, table string) string {
	outputFile := fmt.Sprintf("bcp_%v_%v.dat", db, table)
	target := fmt.Sprintf("/opt/mssql-tools/bin/bcp %v.dbo.%v in %v -c -t ';' -S %v -U %v -P %v",
		db, table, outputFile, dbHost, dbUser, dbPw)

	cmd := exec.Command("bash", "-c", target)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("bcp failed with %s\n CMD: %v", err, target)
	}
	return string(out)
}
