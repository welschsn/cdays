package routing

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func NewDiagnosticRouter() http.Handler {

	r := mux.NewRouter()
	r.HandleFunc("/healthz", handleOK())
	r.HandleFunc("/readyz", handleOK())
	return r
}

func handleOK() func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"Diagnostic endpoint %s called from %s",
			r.URL, r.RemoteAddr,
		)
		fmt.Fprint(w, http.StatusText(http.StatusOK))
	}
}
