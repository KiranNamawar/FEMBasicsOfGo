package main

import (
	"fmt"
	"html/template"
	"museum/api"
	"museum/data"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go server!!"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	html.Execute(w, data.GetAll())

}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api", api.Get)
	server.HandleFunc("/api/new", api.Post)
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)
	err := http.ListenAndServe(":3001", server)
	if err != nil {
		fmt.Println("Error While Starting the server", err)
	}
}
