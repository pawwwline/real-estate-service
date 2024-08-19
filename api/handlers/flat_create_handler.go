package handlers

import (
	"encoding/json"
	"net/http"
	"real-estate-service/api/generated"
)

func (s *MyServer) PostFlatCreate(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("Processing flat create request")

	user_type, ok := r.Context().Value(generated.BearerAuthScopes).(string)
	if !ok {
		http.Error(w, "Failed to get user_type from context", http.StatusInternalServerError)
		return
	}

	if user_type != "moderator" && user_type != "client" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var flat generated.Flat

	err := json.NewDecoder(r.Body).Decode(&flat)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = s.FlatRepositoryInterface.CreateFlat(&flat)
	if err != nil {
		http.Error(w, "Failed to create flat", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"message": "Успешно создана квартира",
		"data":    flat,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		s.Logger.Error("Failed to encode response", "error", err)
	}
}
