package main

import (
	"fmt"
	"gophercises/cyoa/handler"
	"io/ioutil"
	"net/http"
)

func main() {
	_, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		panic(err)
	}
	handler.JSONHandler()
	// fmt.Printf(handler)

	// storyHandler := handler.JSONhandler()

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", defaultMux())
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	return mux
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, handler.JSONHandler())
}
