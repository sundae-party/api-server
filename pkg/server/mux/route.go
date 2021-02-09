package mux

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sundae-party/api-server/pkg/apis/core/types"
	"github.com/sundae-party/api-server/pkg/storage"

	"github.com/gorilla/mux"
)

// ServeAPI list and serve all rest API route
func ServeAPI(r *mux.Router, s storage.Store) {

	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	r.HandleFunc("/api/integration", func(w http.ResponseWriter, r *http.Request) {
		var i types.Integration
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{"ok": err})
		}
		ni, err := s.PutIntegration(r.Context(), &i)
		if err != nil {
			log.Print(err)
			json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		}
		resp, err := json.Marshal(ni)
		if err != nil {
			log.Print("error convert new integration to json")
			json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, "integration": string(resp)})
	}).Methods("POST")

	r.HandleFunc("/api/light/desiredstate", func(w http.ResponseWriter, r *http.Request) {
		var i types.Light
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{"ok": err})
		}
		ni, err := s.UpdateLightStateDesiredState(r.Context(), &i)
		if err != nil {
			log.Print(err)
			json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		}
		resp, err := json.Marshal(ni)
		if err != nil {
			log.Print("error convert new integration to json")
			json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, "integration": string(resp)})
	}).Methods("POST")
}
