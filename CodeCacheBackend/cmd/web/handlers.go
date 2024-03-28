package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("idk")
	fmt.Fprintln(w, "here's your bruh")
}
