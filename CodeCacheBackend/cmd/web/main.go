package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	db       *sql.DB
}

func main() {
	var addr *string = flag.String("addr", "localhost:4321", "HTTP network address")
	var dsn *string = flag.String("dsn", "user=zeoticus dbname=test password=sassenach04", "Data Source Name for the PostgreSQL DB")
	flag.Parse()

	var infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	var errorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := sql.Open("postgres", *dsn)
	err = db.Ping()
	if err != nil {
		errorLog.Println(err)
		return
	}

	var app = &application{
		infoLog,
		errorLog,
		db,
	}

	var mux = http.NewServeMux()
	mux.HandleFunc("/", app.home)

	var srv = &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	log.Printf("Starting server on port %s...\n", *addr)
	err = srv.ListenAndServe()
	log.Println(err)
}
