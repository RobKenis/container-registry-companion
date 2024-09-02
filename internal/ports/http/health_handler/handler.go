package health_handler

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Status string `json:"status"`
}

type Handler struct {
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	health := Health{
		Status: "UP",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(health)
}
