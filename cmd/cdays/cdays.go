package main

import (
	"net/http"
	"log"
	"github.com/welschsn/cdays/internal/routing"
	"os"
	"github.com/welschsn/cdays/internal/version"
	"os/signal"
	"syscall"
)

func main()  {

	log.Printf(
		"The application starting is %s, build time is %s, commit is %v...",
		version.Release, version.BuildTime, version.Commit,
		)


	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("The port wasn't set")
	}

	diagPort := os.Getenv("INTERNAL_PORT")
	if port == "" {
		log.Fatal("The diagnostics port wasn't set")
	}

	log.Printf(
		"The application is listening on port %s, the internal port is %s",
		port, diagPort,
	)


	go func () {
		r := routing.NewBLRouter()
		log.Fatal(http.ListenAndServe(":" + port, r))
	}()

	{
		r := routing.NewDiagnosticRouter()
		log.Fatal(http.ListenAndServe(":" + diagPort, r))
	}


	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case killSignal := <-interrupt:
		log.Printf("Got %s. Stopping...", killSignal)
	}
}
