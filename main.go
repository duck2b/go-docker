package main

import (
	"fmt"
	"log"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	fmt.Fprintf(w, "hello world "+name+": age ="+age)
}

func main() {
	http.HandleFunc("/", SayHello)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
