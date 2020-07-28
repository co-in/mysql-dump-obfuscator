package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var dbUser = flag.String("u", "root", "Database user")
var dbPassword = flag.String("s", "", "Database password")
var dbName = flag.String("d", "", "Database name")
var dbHost = flag.String("h", "127.0.0.1", "Database host")
var dbPort = flag.String("p", "3306", "Database port")

func main() {
	flag.Parse()
	dumpDir := "dumps"
	dumpFilenameFormat := fmt.Sprintf("%s-2006.01.02_15.04.05", *dbName)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", *dbUser, *dbPassword, *dbHost, *dbPort, *dbName)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println("Error opening database: ", err)
		return
	}

	dumper, err := Register(db, dumpDir, dumpFilenameFormat)

	if err != nil {
		fmt.Println("Error registering database:", err)
		return
	}

	resultFilename, err := dumper.Dump(obfuscate)

	if err != nil {
		fmt.Println("Error dumping:", err)
		return
	}

	fmt.Printf("File is saved to %s", resultFilename)

	_ = dumper.Close()

}
