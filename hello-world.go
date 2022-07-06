package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Hi there!\n")

}

func defaultPage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Valid endpoints are:\n")
	fmt.Fprintf(w, "/\n")
	fmt.Fprintf(w, "/headers\n")
	fmt.Fprintf(w, "/hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// data in the page - output html headers
	for name, headers := range req.Header {
		for _, h := range headers {

			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

}

func main() {

	// Setup Logging
	var log = logrus.New()
	log.Out = os.Stdout

	// Needed to serve static assets
	fs := http.FileServer(http.Dir("./img"))
	http.Handle("/img/", http.StripPrefix("/img/", fs))

	// Menu Items
	http.HandleFunc("/", defaultPage)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Serve it on a platter
	http.ListenAndServe(":8090", nil)
}
