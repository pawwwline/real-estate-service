package handlers

import (
	"encoding/json"
	"net/http"
	"real-estate-service/api/generated"
	"real-estate-service/internal/utils"
)

func (s *MyServer) PostFlatCreate(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("Processing flat create request")

	user_type, ok := r.Context().Value(generated.BearerAuthScopes).(string)
	s.Logger.Info("User type:", user_type)
	if !ok {
		http.Error(w, "Не авторизированный доступ", http.StatusForbidden)
		return
	}

	var flat generated.Flat

	err := json.NewDecoder(r.Body).Decode(&flat)
	if err != nil {
		http.Error(w, "Невалидные данные", http.StatusBadRequest)
		return
	}

	err = s.FlatRepositoryInterface.CreateFlat(&flat)
	if err != nil {
		//http.Error(w, "Failed to create flat", http.StatusInternalServerError)
		utils.InternalServerError(w, r, "Failed to create flat")
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
