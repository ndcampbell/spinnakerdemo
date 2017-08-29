package main

import (
	"html/template"
	"log"
	"net/http"
)

const index = `
<html>
	<head>
	</head>
	<body style="background-color:blue;">
		<center>
		<img src="https://secure.meetupstatic.com/photos/event/9/a/highres_436080154.jpeg">
		<h1>Hello!</h1>
		</center>
	</body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("index").Parse(index)
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
