package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func Init() {
	dsn := "phpmyadmin:root@tcp(127.0.0.1:3306)/local"

	var err error
	Conn, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := Conn.Ping(); err != nil {
		panic(err)
	}

	files, err := filepath.Glob(filepath.Join("migrations", "*.sql"))
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		err = run(Conn, file)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Setup complete")
}

func run(db *sql.DB, filePath string) error {
	sqlBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	db.Exec(string(sqlBytes))

	return nil
}
