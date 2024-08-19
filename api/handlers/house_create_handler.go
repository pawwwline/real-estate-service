package handlers

import (
	"encoding/json"
	"net/http"
	"real-estate-service/api/generated"
)

func (s *MyServer) PostHouseCreate(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("Processing house create request")
	user_type, ok := r.Context().Value(generated.BearerAuthScopes).(string)
	if !ok {
		http.Error(w, "Failed to get role from context", http.StatusInternalServerError)
		return
	}

	if user_type != "moderator" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var house generated.House

	err := json.NewDecoder(r.Body).Decode(&house)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = s.HouseRepositoryInterface.CreateHouse(&house)
	if err != nil {
		http.Error(w, "Failed to create house", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Дом успешно создан"})
}
