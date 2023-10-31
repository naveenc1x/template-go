package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

//go:embed templates/*
var templates embed.FS
var t = template.Must(template.ParseFS(templates, "templates/*"))

//go:embed static/*
var assets embed.FS

func main() {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello " + r.RemoteAddr + "the time is\n " + time.Now().Format(time.RFC1123)))
	})

	fmt.Println("Listening on port " + port)
	http.ListenAndServe(":"+port, nil)
}
