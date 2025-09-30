package api

import (
	"encoding/json"
	"net/http"
	"task/game"
)

// api server structure
type APIServer struct {
	Game *game.GameEngine
}

// submit handler
func (s *APIServer) SubmitHandler(w http.ResponseWriter, r *http.Request) {
	var resp game.UserResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	s.Game.Notify <- resp
	w.WriteHeader(http.StatusOK)
}

func (s *APIServer) Start(addr string) error {
	http.HandleFunc("/submit", s.SubmitHandler)
	return http.ListenAndServe(addr, nil)
}
