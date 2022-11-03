package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(responseWritter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWritter, "Wellcome to my go page %s", request.URL.Path)
}
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":7777", nil)
}
