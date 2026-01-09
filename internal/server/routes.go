package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("GET /health", s.HealthHandler)

    // Example of wrapping the mux with middleware (e.g., logging) could go here

	return mux
}

func (s *Server) HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"status": "OK",
		"db_status": s.db.Health(),
	}
    
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}