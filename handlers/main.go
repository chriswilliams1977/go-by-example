package main

import (
	"log"
	"net/http"
)

func main() {
	//create empty mux
	//ServeMux is an HTTP request multiplexer.
	//It matches the URL of each incoming request against a list of registered patterns
	//and calls the handler for the pattern that most closely matches the URL.
	mux := http.NewServeMux()

	//use the mux.Handle function to register this with our new ServeMux,
	//so it acts as the handler for all incoming requests with the URL path /foo.
	mux.Handle("/foo", rh)

	//use the http.RedirectHandler function to create a new handler
	//This handler 307 redirects all requests it receives to http://example.org.
	rh := http.RedirectHandler("http://example.org", 307)

	//create a new server and start listening for incoming requests with the http.ListenAndServe function,
	//passing in our ServeMux for it to match requests against
	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
