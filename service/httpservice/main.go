package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

var alive = true

func main() {
	name := flag.String("name", "backend-1", "Name of this service")
	listen := flag.String("listen", ":9000", "Http listen addr")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !alive {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal Server Error")
			return
		}
		strA := r.FormValue("a")
		strB := r.FormValue("b")

		intA, err := strconv.Atoi(strA)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Wrong input")
			return
		}

		intB, err := strconv.Atoi(strB)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Wrong input")
			return
		}

		fmt.Fprintf(w, "Result from %s: %d\n", *name, intA+intB)
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
