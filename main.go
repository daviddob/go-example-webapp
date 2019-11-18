package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

var CommitSHA string
var BuildTime string

type example struct {
	Time       string
	BuildTime  string
	CommitHash string
	Hostname   string
}

func main() {

	templates := template.Must(template.ParseFiles("templates/example-template.html"))
	hostname, _ := os.Hostname()

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		example := example{time.Now().Format(time.Stamp), BuildTime, CommitSHA, hostname}
		if err := templates.ExecuteTemplate(w, "example-template.html", example); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening on Port 8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
