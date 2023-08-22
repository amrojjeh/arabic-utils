package main

import (
	"github.com/amrojjeh/goarabic"
	"html/template"
	"log"
	"net/http"
)

func main() {
	templates := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index.html", nil)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/ar2safebw", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if !r.Form.Has("in") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		input := r.Form.Get("in")
		output, err := goarabic.ArToSafeBW(input)
		if err != nil {
			w.Write([]byte("Input is incorrect!"))
			return
		}
		w.Write([]byte(output))
	})

	log.Println("Listening on http://127.0.0.1:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
