package routing

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
)

func NewBLRouter() http.Handler {

	r := mux.NewRouter()
	r.HandleFunc("/home", rootHandler())
	return r
}

func rootHandler() func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"Application endpoint called from %s",
			r.RemoteAddr,
		)
		w.Write([]byte("Hello ! Here is your request:\n"))
		fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
		w.Write([]byte("Header\n"))
		for k, v := range r.Header {
			fmt.Fprintf(w, "%s:\t%q\n", k, v)
		}
		fmt.Fprintf(w, "Host:\t%q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddress:\t%q\n", r.RemoteAddr)
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		}
	}
}
