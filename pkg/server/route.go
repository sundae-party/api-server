package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// ServeAPI list and serve all rest API route
func ServeAPI(r *mux.Router) {

	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
}
