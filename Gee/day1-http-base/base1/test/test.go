package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("base1 test")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHander)
	log.Fatal(http.ListenAndServe(":9999", nil))

}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q \n", req.URL.Path)
}

func helloHander(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
