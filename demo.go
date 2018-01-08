package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func getIndex() string {

	index, _ := ioutil.ReadFile("index.html")
	return string(index)
}

func handler(w http.ResponseWriter, r *http.Request) {
	index := getIndex()
	t, _ := template.New("index").Parse(index)
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
