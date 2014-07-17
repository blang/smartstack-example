package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var alive = true

func main() {
	name := flag.String("name", "frontend-1", "Name of this service")
	listen := flag.String("listen", ":80", "Http listen addr")
	service := flag.String("service", "http://127.0.0.1:9000/", "Address to service backend")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "My Name is: %s\n", *name)
		if alive {
			fmt.Fprint(w, "Everything is ok")
		} else {
			fmt.Fprint(w, "I'm dead...")
		}
	})

	http.HandleFunc("/calc", func(w http.ResponseWriter, r *http.Request) {
		if !alive {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal Server Error")
			return
		}
		strA := r.FormValue("a")
		strB := r.FormValue("b")

		_, err := strconv.Atoi(strA)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Wrong input")
			return
		}

		_, err = strconv.Atoi(strB)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Wrong input")
			return
		}

		resp, err := http.Get(*service + "?a=" + strA + "&b=" + strB)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			fmt.Fprintf(w, "Error on requesting Backend: ", err)
			return
		}

		fmt.Fprintf(w, "My Name is: %s\n", *name)
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error on backend response: %s", err)
			return
		}
		fmt.Fprintf(w, "Backend: %s\n", string(b))

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
