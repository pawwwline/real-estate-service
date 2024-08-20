package handlers

import (
	"encoding/json"
	"net/http"
	"real-estate-service/api/generated"
	"real-estate-service/internal/utils"
)

func (s *MyServer) PostHouseCreate(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("Processing house create request")
	user_type, ok := r.Context().Value(generated.BearerAuthScopes).(string)
	if !ok {
		utils.InternalServerError(w, r, "Invalid or missing user type")
		return
	}

	if user_type != "moderator" {
		http.Error(w, "Неавторизованный доступ", http.StatusForbidden)
		return
	}

	var house generated.House

	err := json.NewDecoder(r.Body).Decode(&house)
	if err != nil {
		http.Error(w, "Невалидные данные ввода", http.StatusBadRequest)
		return
	}

	if s.HouseRepositoryInterface == nil {
		utils.InternalServerError(w, r, "что-то пошло не так")
		return
	}

	err = s.HouseRepositoryInterface.CreateHouse(&house)
	if err != nil {
		utils.InternalServerError(w, r, "что-то пошло не так")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Дом успешно создан"})
}
