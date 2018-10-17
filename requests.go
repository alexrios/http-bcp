package main

type ExportRequest struct {
	db     string `valid:"alphanum"`
	schema string `valid:"alphanum"`
	table  string `valid:"alphanum"`
}

type ImportRequest struct {
	db     string `valid:"alphanum"`
	schema string `valid:"alphanum"`
	table  string `valid:"alphanum"`
}
