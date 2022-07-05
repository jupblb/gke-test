package main

import (
	"bytes"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/google/uuid"
)

var (
	tpl      = template.Must(template.ParseFiles("index.html"))
	workerId = uuid.New()
)

type Website struct {
	Date     string
	WorkerId string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}

	err := tpl.Execute(buf, Website{
		Date:     time.Now().Format(time.RFC3339),
		WorkerId: workerId.String(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
