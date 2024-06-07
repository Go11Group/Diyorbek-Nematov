package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type URLData struct {
	URL string `json:"url"`
}

type ParamData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/api/url-data/", URLDataHandler)
	http.HandleFunc("/api/body-data", BodyDataHandler)
	http.HandleFunc("/api/param-data", ParamDataHandler)
	http.HandleFunc("/api/url-to-body", URLToBodyHandler)
	http.HandleFunc("/api/body-to-param", BodyToParamHandler)
	http.HandleFunc("/api/param-to-url", ParamToURLHandler)
	http.HandleFunc("/api/url-to-json", URLToJSONHandler)
	http.HandleFunc("/api/body-to-json", BodyToJSONHandler)
	http.HandleFunc("/api/param-to-json", ParamToJSONHandler)
	http.HandleFunc("/api/json-to-url", JSONToURLHandler)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func URLDataHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	fmt.Fprintf(w, "Received URL: %s", url)
}

func BodyDataHandler(w http.ResponseWriter, r *http.Request) {
	var data URLData
	json.NewDecoder(r.Body).Decode(&data)
	fmt.Fprintf(w, "Received Body Data: %+v", data)
}

func ParamDataHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Fprintf(w, "Received Param ID: %s", id)
}

func URLToBodyHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	fmt.Fprintf(w, "Received URL: %s", url)
}

func BodyToParamHandler(w http.ResponseWriter, r *http.Request) {
	var data URLData
	json.NewDecoder(r.Body).Decode(&data)
	fmt.Fprintf(w, "Received Body Data: %+v", data)
}

func ParamToURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	url := fmt.Sprintf("http://example.com?id=%s&name=%s", id, name)
	fmt.Fprintf(w, "Generated URL: %s", url)
}

func URLToJSONHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	data := URLData{URL: url}
	json.NewEncoder(w).Encode(data)
}

func BodyToJSONHandler(w http.ResponseWriter, r *http.Request) {
	var data URLData
	json.NewDecoder(r.Body).Decode(&data)
	json.NewEncoder(w).Encode(data)
}

func ParamToJSONHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	data := ParamData{ID: id, Name: name}
	json.NewEncoder(w).Encode(data)
}

func JSONToURLHandler(w http.ResponseWriter, r *http.Request) {
	var data URLData
	json.NewDecoder(r.Body).Decode(&data)
	url := fmt.Sprintf("http://example.com?url=%s", data.URL)
	fmt.Fprintf(w, "Generated URL: %s", url)
}
