package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/v1/personInfo", info)
	http.ListenAndServe(":8080", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is cool!")
}

func info(w http.ResponseWriter, r *http.Request) {
	p := Person{Name: "praveen", Age: 21}
	t, _ := template.ParseFiles("info.html")
	t.Execute(w, p)
}
