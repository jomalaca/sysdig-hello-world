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

	const AddHeader = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<title>Sysdig Hello World</title>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
	<!--===============================================================================================-->
		<link rel="icon" type="image/png" href="html/images/icons/favicon.ico"/>
	<!--===============================================================================================-->
		<link rel="stylesheet" type="text/css" href="html/vendor/bootstrap/css/bootstrap.min.css">
	<!--===============================================================================================-->
		<link rel="stylesheet" type="text/css" href="html/fonts/font-awesome-4.7.0/css/font-awesome.min.css">
	<!--===============================================================================================-->
		<link rel="stylesheet" type="text/css" href="html/vendor/animate/animate.css">
	<!--===============================================================================================-->
		<link rel="stylesheet" type="text/css" href="html/vendor/select2/select2.min.css">
	<!--===============================================================================================-->
		<link rel="stylesheet" type="text/css" href="html/css/util.css">
		<link rel="stylesheet" type="text/css" href="html/css/main.css">
	<!--===============================================================================================-->
	</head>
	<body>


		<div class="size1 bg0 where1-parent">
			<div class="flex-c-m bg-img1 size2 where1 overlay1 where2 respon2" style="background-image: url('html/images/bg01.jpg');">
			</div>

			<div class="size3 flex-col-sb flex-w p-l-75 p-r-75 p-t-45 p-b-45 respon1">
				<div class="wrap-pic1">
					<img src="html/images/icons/logo.png" alt="LOGO">
				</div>

				<div class="p-t-50 p-b-60">
					<p class="m1-txt1 p-b-36">
						<span class="m1-txt2">Output</span>
					</p>
					<!-- END HEADER -->
	`

	const AddFooter = `
	<p class="s2-txt3 p-t-18">
	This application is used for testing <a href="https://sysdig.com/">Sysdig</a> inline scanning.
</p>
</div>

<div class="flex-w">
<a href="https://twitter.com/sysdig" class="flex-c-m size5 bg4 how1 trans-04 m-r-5">
	<i class="fa fa-twitter"></i>
</a>

<a href="http://www.youtube.com/c/sysdig" class="flex-c-m size5 bg5 how1 trans-04 m-r-5">
	<i class="fa fa-youtube-play"></i>
</a>
</div>
</div>
</div>

<!--===============================================================================================-->
<script src="html/vendor/jquery/jquery-3.2.1.min.js"></script>
<!--===============================================================================================-->
<script src="html/vendor/bootstrap/js/popper.js"></script>
<script src="html/vendor/bootstrap/js/bootstrap.min.js"></script>
<!--===============================================================================================-->
<script src="html/vendor/select2/select2.min.js"></script>
<!--===============================================================================================-->
<!--===============================================================================================-->
<script src="html/vendor/tilt/tilt.jquery.min.js"></script>
<script >
$('.js-tilt').tilt({
scale: 1.1
})
</script>
<!--===============================================================================================-->
<script src="html/js/main.js"></script>

</body>
</html>
	`

	fmt.Fprint(w, AddHeader)
	fmt.Fprintf(w, "hello, dk!\n")
	fmt.Fprint(w, AddFooter)

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
