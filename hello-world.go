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
		`
	const footer = `
    </body>
</html>
`

	w.Header().Add("Content Type", "text/html")
	templates, _ := template.New("header").Parse(header)
	context := Context{
		Title: "Sysdig Hello World - Home",
		Links: [2]string{"headers", "hello"},
	}
	templates.Lookup("header").Execute(w, context)

	// data in the page - output html headers
	// data in the page
	fmt.Fprintf(w, "Hi there!\n")

	templates2, _ := template.New("footer").Parse(footer)

	templates2.Lookup("footer").Execute(w, nil)
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
		<p>Please select your link.</p>
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

}

func headers(w http.ResponseWriter, req *http.Request) {

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
		`
	const footer = `
    </body>
</html>
`

	w.Header().Add("Content Type", "text/html")
	templates, _ := template.New("header").Parse(header)
	context := Context{
		Title: "Sysdig Hello World - Home",
		Links: [2]string{"headers", "hello"},
	}
	templates.Lookup("header").Execute(w, context)

	// data in the page - output html headers
	for name, headers := range req.Header {
		for _, h := range headers {

			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

	templates2, _ := template.New("footer").Parse(footer)

	templates2.Lookup("footer").Execute(w, nil)

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
