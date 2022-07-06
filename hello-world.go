package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/sirupsen/logrus"
)

type MyData struct {
	Title string

	Homelinktxt string
	Homelink    string

	Hellolinktxt string
	Hellolink    string

	Headerlinktxt string
	Headerlink    string
}

func hello(w http.ResponseWriter, req *http.Request) {

	helloData := MyData{
		Title:         "Sysdig Hello World - Hello",
		Homelinktxt:   "Home",
		Homelink:      "/",
		Hellolinktxt:  "Hello",
		Hellolink:     "/hello",
		Headerlinktxt: "Header",
		Headerlink:    "/headers",
	}
	parsedTemplate, _ := template.ParseFiles("template-header.html")
	err := parsedTemplate.Execute(w, helloData)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}

	fmt.Fprintf(w, "Hi there!\n")

	parsedTemplate2, _ := template.ParseFiles("template-footer.html")
	err2 := parsedTemplate2.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template :", err2)
		return
	}
}

func defaultPage(w http.ResponseWriter, req *http.Request) {
	homeData := MyData{
		Title:         "Sysdig Hello World - Home",
		Homelinktxt:   "Home",
		Homelink:      "/",
		Hellolinktxt:  "Hello",
		Hellolink:     "/hello",
		Headerlinktxt: "Header",
		Headerlink:    "/headers",
	}
	parsedTemplate, _ := template.ParseFiles("home-template.html")
	err := parsedTemplate.Execute(w, homeData)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	headerData := MyData{
		Title:         "Sysdig Hello World - Headers",
		Homelinktxt:   "Home",
		Homelink:      "/",
		Hellolinktxt:  "Hello",
		Hellolink:     "/hello",
		Headerlinktxt: "Header",
		Headerlink:    "/headers",
	}
	parsedTemplate, _ := template.ParseFiles("template-header.html")
	err := parsedTemplate.Execute(w, headerData)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}

	for name, headers := range req.Header {
		for _, h := range headers {

			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

	parsedTemplate2, _ := template.ParseFiles("template-footer.html")
	err2 := parsedTemplate2.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template :", err2)
		return
	}
}

func main() {

	var log = logrus.New()
	log.Out = os.Stdout

	http.HandleFunc("/", defaultPage)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
