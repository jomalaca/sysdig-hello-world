package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/sirupsen/logrus"
)

type Context struct {
	Title string
	Links [2]string
}

// Data for HTML templates
// type MyData struct {
// 	Title string

// 	Homelinktxt string
// 	Homelink    string

// 	Hellolinktxt string
// 	Hellolink    string

// 	Headerlinktxt string
// 	Headerlink    string
// }

func hello(w http.ResponseWriter, req *http.Request) {

	// Create data - should centralize
	// helloData := MyData{
	// 	Title:         "Sysdig Hello World - Hello",
	// 	Homelinktxt:   "Home",
	// 	Homelink:      "/",
	// 	Hellolinktxt:  "Hello",
	// 	Hellolink:     "/hello",
	// 	Headerlinktxt: "Header",
	// 	Headerlink:    "/headers",
	// }

	// Title := "Sysdig Hello World - Home"
	// // Output html header
	// parsedTemplate, _ := template.ParseFiles("template-header.html")
	// err := parsedTemplate.Execute(w, Title)
	// if err != nil {
	// 	log.Println("Error executing template :", err)
	// 	return
	// }

	// data in the page
	fmt.Fprintf(w, "Hi there!\n")

	// Output html footer
	// parsedTemplate2, _ := template.ParseFiles("template-footer.html")
	// err2 := parsedTemplate2.Execute(w, nil)
	// if err != nil {
	// 	log.Println("Error executing template :", err2)
	// 	return
	// }
}

func defaultPage(w http.ResponseWriter, req *http.Request) {

	const doc = `
<!DOCTYPE html>
<html>
    <head>
		<title>
        {{.Title}}
		</title>
    </head>
	<style>
			nav {
				background: #515151;
				text-align: center;
				font-family: Arial, Helvetica, sans-serif;
				/* THIS IS OPTIONAL*/
			}

			h1, p {
				font-family: Arial, Helvetica, sans-serif;
			}

			h1 {
				color: #515151;
			}

			nav a {
				display: inline-block;
				color: #FFF;
				padding: 18px 12px;
				text-decoration: none;
				transition: ease-in .3s;
			}

			nav a:hover {
				color: #515151;
				background: #FFF;
			}

			nav ul {
				list-style-type: none;
			}

			nav ul li {
				display: inline;
			}
		</style>
    <body>
		<img src="img/sysdig_Horz_Color_Logo_RGB_sml.jpg" alt="Sysdig Logo"/>
		<nav>
        <ul>
			<li><a href="/">Home</a></li>
            {{range .Links}}
				<li><a href="/{{.}}">{{.}}</a></li>
            {{end}}
        </ul>
		</nav>
		<h1>{{.Title}}</h1>
    </body>
</html>
`

	w.Header().Add("Content Type", "text/html")
	templates, _ := template.New("doc").Parse(doc)
	context := Context{
		Title: "Sysdig Hello World - Home",
		Links: [2]string{"headers", "hello"},
	}
	templates.Lookup("doc").Execute(w, context)
	// Create data - should centralize
	// homeData := MyData{
	// 	Title:         "Sysdig Hello World - Home",
	// 	Homelinktxt:   "Home",
	// 	Homelink:      "/",
	// 	Hellolinktxt:  "Hello",
	// 	Hellolink:     "/hello",
	// 	Headerlinktxt: "Header",
	// 	Headerlink:    "/headers",
	// }

	// Output html for home

	// Title := "Sysdig Hello World - Home"
	// parsedTemplate, _ := template.ParseFiles("home-template.html")
	// err := parsedTemplate.Execute(w, Title)
	// if err != nil {
	// 	log.Println("Error executing template :", err)
	// 	return
	// }
}

func headers(w http.ResponseWriter, req *http.Request) {

	// Create data - should centralize
	// headerData := MyData{
	// 	Title:         "Sysdig Hello World - Headers",
	// 	Homelinktxt:   "Home",
	// 	Homelink:      "/",
	// 	Hellolinktxt:  "Hello",
	// 	Hellolink:     "/hello",
	// 	Headerlinktxt: "Header",
	// 	Headerlink:    "/headers",
	// }

	// Output html header
	// parsedTemplate, _ := template.ParseFiles("template-header.html")
	// err := parsedTemplate.Execute(w, headerData)
	// if err != nil {
	// 	log.Println("Error executing template :", err)
	// 	return
	// }

	// data in the page - output html headers
	for name, headers := range req.Header {
		for _, h := range headers {

			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

	// Output html footer
	// parsedTemplate2, _ := template.ParseFiles("template-footer.html")
	// err2 := parsedTemplate2.Execute(w, nil)
	// if err != nil {
	// 	log.Println("Error executing template :", err2)
	// 	return
	// }
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
