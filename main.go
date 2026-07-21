package main

import (
	"fmt"
	"net/http"
	"time"

	"richienzl/myapp/views"

	"github.com/a-h/templ"
)

func main() {
	mux := http.NewServeMux()

	// 1. Add GET prefix here to match the method constraint
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

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

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", mux)
}
