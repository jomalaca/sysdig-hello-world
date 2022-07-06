package main

import (
	"fmt"
	"html/template"
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
	myVar := "From the var!"
	tmpl, err := template.ParseFiles("html/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := struct {
		MyVar string
	}{
		MyVar: myVar,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
