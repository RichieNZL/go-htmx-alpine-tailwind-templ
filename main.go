package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"richienzl/myapp/views"

	"github.com/a-h/templ"
)

func main() {
	mux := http.NewServeMux()

	// 1. Add GET prefix here to match the method constraint
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Direct route for /favicon.ico -> static/favicon.ico
	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.svg")
	})

	// 2. Main Page Route
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		// Render layout with Counter child component
		views.Layout("Go Stack").Render(
			templ.WithChildren(r.Context(), views.Counter()),
			w,
		)
	})

	// 3. HTMX Endpoint
	mux.HandleFunc("GET /time", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format("15:04:05 PM")
		views.TimeResponse(currentTime).Render(r.Context(), w)
	})

	fmt.Println("Server running on https://localhost:8443")
	err := http.ListenAndServeTLS(":8443", "certs/localhost+2.pem", "certs/localhost+2-key.pem", mux)

	if err != nil {
		log.Fatal(err)
	}
}
