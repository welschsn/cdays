package main

import (
	"net/http"
	"log"
	"github.com/welschsn/cdays/internal/routing"
	"os"
	"github.com/welschsn/cdays/internal/version"
)

func main()  {

	log.Printf("The application starting is %s, build time is %s, commit is %v...", version.Release, version.BuildTime, version.Commit )


	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("The port wasn't set")
	}

	diagPort := os.Getenv("DIAGNOSTICS_PORT")
	if port == "" {
		log.Fatal("The diagnostics port wasn't set")
	}


	go func () {
		r := routing.NewBLRouter()
		log.Fatal(http.ListenAndServe(":" + port, r))
	}()

	{
		r := routing.NewDiagnosticRouter()
		log.Fatal(http.ListenAndServe(":" + diagPort, r))
	}


	log.Print("The application has finished")

}

func rootHandler() func(http.ResponseWriter, *http.Request ) {

	return func(w http.ResponseWriter, r * http.Request ){
		w.Write([]byte("Hello !"))
	}
}