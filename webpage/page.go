package main

import (
	"net/http"
)

func homeHandlerFunc(responseWritter http.ResponseWriter, request *http.Request) {
	responseWritter.Write([]byte("Hello go lovers"))
}
func aboutHandlerFunc(responseWritter http.ResponseWriter, request *http.Request) {
	responseWritter.Write([]byte("About go"))
}
func main() {
	http.HandleFunc("/", homeHandlerFunc)
	http.HandleFunc("/home", homeHandlerFunc)
	http.HandleFunc("/about", aboutHandlerFunc)
	http.ListenAndServe(":7777", nil)
}
