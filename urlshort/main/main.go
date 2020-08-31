package main

import (
	"fmt"
	"net/http"

	"urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/fb": "https://www.facebook.com",
		"/pexa":     "https://www.peaxa.com.au",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
  - path: /fb
    url: https://www.facebook.com
  - path: /pexa
    url: https://www.peaxa.com.au
  `
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", notFound)
	return mux
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not Found")
}
