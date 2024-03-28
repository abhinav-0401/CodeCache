package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	var addr *string = flag.String("addr", "localhost:4321", "HTTP network address")
	flag.Parse()

	var infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	var errorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	var app = &application{
		infoLog,
		errorLog,
	}

	var mux = http.NewServeMux()
	mux.HandleFunc("/", app.home)

	var srv = &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	log.Printf("Starting server on port %s...\n", *addr)
	err := srv.ListenAndServe()
	log.Println(err)
}
