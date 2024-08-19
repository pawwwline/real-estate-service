package handlers

import (
	"encoding/json"
	"net/http"
	"real-estate-service/api/generated"
)

func (s *MyServer) PostFlatUpdate(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("Processing flat update request")
	user_type, ok := r.Context().Value(generated.BearerAuthScopes).(string)
	if !ok {
		http.Error(w, "Failed to get user_type from context", http.StatusInternalServerError)
		return
	}

	if user_type != "moderator" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var flat generated.Flat

	err := json.NewDecoder(r.Body).Decode(&flat)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	currentFlat, err := s.FlatRepositoryInterface.GetFlatId(flat.Id)
	if err != nil {
		http.Error(w, "Flat not found", http.StatusNotFound)
		return
	}

	if currentFlat.Status == "on moderation" && flat.Status != "approved" && flat.Status != "declined" {
		http.Error(w, "Flat is already under moderation", http.StatusConflict)
		return
	}

	err = s.FlatRepositoryInterface.UpdateFlat(&flat)
	if err != nil {
		http.Error(w, "Failed to update flat", http.StatusInternalServerError)
		return
	}

	err = s.FlatRepositoryInterface.UpdateFlat(&flat)
	if err != nil {
		http.Error(w, "Failed to create house", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"message": "Успешно обновлена квартира",
		"data":    flat,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		s.Logger.Error("Failed to encode response", "error", err)
	}
}
