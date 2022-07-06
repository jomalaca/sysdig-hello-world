package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func dk(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello, dk!\n")
}

func defaultPage(w http.ResponseWriter, req *http.Request) {

	// 	const AddHtml = `

	// `

	// fmt.Fprint(w, AddHtml)
	fmt.Fprintf(w, "Welcome!\n")
	fmt.Fprintf(w, "Welcome!\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	var log = logrus.New()
	log.Out = os.Stdout

	http.HandleFunc("/", defaultPage)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/dk", dk)

	http.ListenAndServe(":8090", nil)
}
