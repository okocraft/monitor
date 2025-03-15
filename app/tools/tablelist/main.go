package main

import (
	"fmt"
	"os"

	"github.com/Siroshun09/go-tablelist-codegen/database"
	"github.com/Siroshun09/go-tablelist-codegen/generator"
	"github.com/Siroshun09/serrors"
	"github.com/okocraft/monitor/lib/testutils"
	"github.com/okocraft/monitor/lib/testutils/testdb"
)

//go:generate go run main.go

func main() {
	db, err := testdb.NewTestDB(false)
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	defer func() {
		err := db.Cleanup()
		if err != nil {
			printError(err)
			os.Exit(1)
		}
	}()

	conn, err := db.Conn()
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	tables, err := database.GetTables(conn.Base(), database.QueryForMySQL)
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	root, err := testutils.GetProjectRoot()
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	file, err := os.OpenFile(root+"/internal/repositories/queries/tables.gen.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
	if err != nil {
		printError(err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			printError(err)
			os.Exit(1)
		}
	}(file)

	err = generator.GenerateCode(file, generator.TemplateParam{PackageName: "queries", Tables: tables})
	if err != nil {
		printError(err)
		os.Exit(1)
	}
}

func printError(err error) {
	fmt.Println(err.Error())
	fmt.Println(serrors.GetStackTrace(err).String())
}
