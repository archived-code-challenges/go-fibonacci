package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/noelruault/go-fibonacci/internal/models"
	"github.com/noelruault/go-fibonacci/internal/web"
)

type API struct {
	fibonaccisvc models.FibonacciService

	log *log.Logger
}

func NewAPI(log *log.Logger) API {
	fs := models.NewFibonacciService()

	return API{
		fibonaccisvc: fs,
		log:          log,
	}
}

// Health endpoint
func (app *App) Health(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var health struct {
		Status string `json:"status"`
	}
	health.Status = "ok"
	web.Respond(ctx, w, health, http.StatusOK)
}

// GetFibonacci will return up to the 94th sequence of a N-th Fibonacci number.
func (app *App) GetFibonacci(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	number, err := strconv.ParseUint(r.URL.Query().Get("n"), 10, 8)
	// The bitSize=8 argument specifies the integer type that the result must fit into. More info at https://golang.org/pkg/strconv/#ParseInt
	if err != nil || number < 1 {
		web.Respond(ctx, w, map[string]interface{}{"error": "error on decoding URL"}, http.StatusNotAcceptable)
		return
	}

	if number > 93 {
		// We can return correct value of 93th Fibonacci number, which is 12,200,160,415,121,876,738.
		// The 94th sequence number (19,740,274,219,868,223,167) would require 65-bit long variable.
		web.Respond(ctx, w, map[string]interface{}{"error": "highest 'n' suported is 93"}, http.StatusNotAcceptable)
		return
	}

	fibonacci := app.API.fibonaccisvc.Iterative(uint8(number))
	if err != nil {
		web.Respond(ctx, w, map[string]interface{}{"error": "server error"}, http.StatusInternalServerError)
	}

	web.Respond(ctx, w, fibonacci, http.StatusOK)
}
