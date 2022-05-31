package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Adventure struct {
	Title   string
	Story   []string
	Options []struct {
		Text string
		Arc  string
	}
}

func ChooseYourOwnAdventure() {
	mux := http.NewServeMux()
	mux.Handle("/", &home{})
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mux)
}

type home struct{}

func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jsonFile, _ := os.Open("../../resources/gopher.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Declared an empty interface
	var result map[string]Adventure
	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal(byteValue, &result)

	t := template.Must(template.ParseFiles("../../resources/home.html"))
	t.Execute(w, result["intro"])
}

// Tener cuidado al crear los objetos que se mapean en las templates, los campos deben sen publicos (primera mayuscula)
// inicializar arrays en ejecucion con make si no al indexar un valor lanza un outofboundsexception
