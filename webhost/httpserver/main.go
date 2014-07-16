package main

import (
	"flag"
	"fmt"
	"net/http"
)

var alive = true

func main() {
	name := flag.String("name", "example", "Name of this service")
	listen := flag.String("listen", ":80", "Http listen addr")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "My Name is: %s\n", *name)
		if alive {
			fmt.Fprint(w, "Everything is ok")
		} else {
			fmt.Fprint(w, "I'm dead...")
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if alive {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "OK")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Not OK")
		}
	})

	http.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) {
		alive = false
		fmt.Fprint(w, "You killed me...")
	})

	http.HandleFunc("/reanimate", func(w http.ResponseWriter, r *http.Request) {
		alive = true
		fmt.Fprint(w, "You saved my life...")
	})

	fmt.Printf("HTTP Server started, listening on %s\n", *listen)
	if err := http.ListenAndServe(*listen, nil); err != nil {
		fmt.Printf("HTTP Server returned with error: %s\n", err)
	}
}
