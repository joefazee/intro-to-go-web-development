package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"abahjoseph.com/books/pkg/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	users    *mysql.UserModel
}

func main() {

	addr := flag.String("addr", ":9007", "Pass the network address")
	dns := flag.String("dns", "root:root@tcp(localhost:8889)/ajizzy?parseTime=true", "MYSQL Connection string")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dns)
	if err != nil {
		errLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errLog,
		infoLog:  infoLog,
		users:    &mysql.UserModel{DB: db},
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting our server on %s\n", *addr)
	if err := server.ListenAndServe(); err != nil {
		errLog.Fatal(err)
	}
}

func openDB(dns string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
