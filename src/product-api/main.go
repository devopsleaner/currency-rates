package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	f := func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s", d)
	}
	f1 := func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	}
	http.HandleFunc("/", f)
	http.HandleFunc("/GoodBye", f1)

	http.ListenAndServe(":9000", nil)
}
