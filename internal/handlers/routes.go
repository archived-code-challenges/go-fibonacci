package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	API    API
}

func (app *App) SetupRouter() {
	app.Router.
		Methods(http.MethodGet).
		Path("/").
		HandlerFunc(app.Health)

	app.Router.
		Methods(http.MethodGet).
		Path("/fib").
		HandlerFunc(app.GetFibonacci)
}
