package handlers

import (
	"encoding/json"
	"net/http"
	"real-estate-service/api/generated"
)

func (s *MyServer) GetHouseId(w http.ResponseWriter, r *http.Request, id generated.HouseId) {
	s.Logger.Info("Processing get house by ID request")

	user_type, ok := r.Context().Value(generated.BearerAuthScopes).(string)

	if !ok {
		http.Error(w, "Failed to get role from context", http.StatusInternalServerError)
		return
	}

	var err error
	var flats []generated.Flat

	switch user_type {
	case "client":
		flats, err = s.FlatRepositoryInterface.GetApprovedFlatsByHouseId(id)
	case "moderator":
		flats, err = s.FlatRepositoryInterface.GetFlatsByHouseId(id)
	default:
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	if err != nil {
		http.Error(w, "Failed to retrieve flats", http.StatusInternalServerError)
		return
	}

	response := flats

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

}
