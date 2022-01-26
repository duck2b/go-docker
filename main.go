package main

import (
	"fmt"
	"log"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request)  {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w,"hello world " + name)
}

func main()  {
	http.HandleFunc("/", SayHello)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}