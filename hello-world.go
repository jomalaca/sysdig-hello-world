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

func writeHtmlHeader(w http.ResponseWriter, req *http.Request, context Context) {
	const header = `
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
			<p>
			`

	w.Header().Add("Content Type", "text/html")
	templates, _ := template.New("header").Parse(header)
	templates.Lookup("header").Execute(w, context)

}

func writeHtmlFooter(w http.ResponseWriter, req *http.Request, context Context) {
	const footer = `
	</p>
    </body>
</html>
`

	w.Header().Add("Content Type", "text/html")
	templates, _ := template.New("footer").Parse(footer)
	templates.Lookup("footer").Execute(w, context)

}

// =========

func hello(w http.ResponseWriter, req *http.Request) {

	c := Context{
		Title: "Sysdig Hello World - Hello",
		Links: [2]string{"headers", "hello"},
	}
	writeHtmlHeader(w, req, c)
	// data in the page - output html headers
	// data in the page
	fmt.Fprintf(w, "Hi there!\n")

	writeHtmlFooter(w, req, c)
}

func defaultPage(w http.ResponseWriter, req *http.Request) {

	c := Context{
		Title: "Sysdig Hello World - Home",
		Links: [2]string{"headers", "hello"},
	}
	writeHtmlHeader(w, req, c)
	// data in the page - output html headers
	// data in the page
	fmt.Fprintf(w, "Please select your link.\n")

	writeHtmlFooter(w, req, c)

}

func headers(w http.ResponseWriter, req *http.Request) {

	c := Context{
		Title: "Sysdig Hello World - Headers",
		Links: [2]string{"headers", "hello"},
	}
	writeHtmlHeader(w, req, c)

	// data in the page - output html headers
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

	writeHtmlFooter(w, req, c)
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
