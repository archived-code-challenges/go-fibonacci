package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"

	"github.com/noelruault/go-fibonacci/internal/handlers"
)

const servicePort = ":8080"

func main() {
	serverErrors := make(chan error, 1)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("main : API listening on %s", servicePort)
		serverErrors <- run()
	}()

	select {
	case err := <-serverErrors:
		log.Printf("error starting server: %s", err)

	case sig := <-shutdown:
		log.Printf("main : %v : Start shutdown", sig)
		switch {
		case sig == syscall.SIGSTOP:
			log.Printf("integrity issue caused shutdown")
		}
	}

}

func run() error {
	log := log.New(os.Stdout, "GO-FIBONACCI : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	log.Printf("main : Started")
	defer log.Println("main : Completed")

	app := &handlers.App{
		Router: mux.NewRouter(),
		API:    handlers.NewAPI(log),
	}

	app.SetupRouter()

	return http.ListenAndServe(servicePort, app.Router)
}
