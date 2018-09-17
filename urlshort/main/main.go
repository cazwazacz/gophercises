package main

import (
	"flag"
	"fmt"
	"gophercises/urlshort"
	"io/ioutil"
	"net/http"
)

func main() {
	yamlPath := flag.String("yaml", "", "Specify a yml file to load paths from")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlString := `
    - path: /urlshort
      url: https://github.com/gophercises/urlshort
    - path: /urlshort-final
      url: https://github.com/gophercises/urlshort/tree/solution
    `

	yaml := []byte(yamlString)

	if *yamlPath != "" {
		yamlFile, err := ioutil.ReadFile(*yamlPath)
		if err != nil {
			panic(err)
		}

		yaml = yamlFile
	}

	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}

	json := []byte(`[{"path": "/json", "url": "https://gobyexample.com/json"}]`)

	jsonHandler, err := urlshort.JSONHandler(json, yamlHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
