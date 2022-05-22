package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

type YamlStruct struct {
	Path string
	Url  string
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := pathsToUrls[r.URL.Path]
		if url != "" {
			http.Redirect(w, r, url, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(w, r)
		}
	})
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var a []YamlStruct

	err := yaml.Unmarshal([]byte(yml), &a)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
		return nil, err
	}

	return MapHandler(buildMap(a), fallback), nil
}

func buildMap(yml []YamlStruct) map[string]string {
	yamlMap := make(map[string]string)

	for _, element := range yml {
		yamlMap[element.Path] = element.Url
	}

	return yamlMap
}

func UrlShortener() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
