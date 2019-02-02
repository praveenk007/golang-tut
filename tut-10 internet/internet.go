package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	resp, _ := http.Get("https://reqres.in/api/users?page=2")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	resp.Body.Close()
	fmt.Println(string_body)
}
