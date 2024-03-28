package main

import "net/http"

func (app *application) routes() http.Handler {
	var mux = http.NewServeMux()

	mux.HandleFunc("GET /bruh", app.home)

	return mux
}
