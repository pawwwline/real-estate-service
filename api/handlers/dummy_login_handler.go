package handlers

import (
	"log/slog"
	"net/http"
	"real-estate-service/api/generated"
	"real-estate-service/repository"
)

type MyServer struct {
	Logger *slog.Logger
	generated.Unimplemented
	HouseRepositoryInterface repository.HouseRepositoryInterface
	FlatRepositoryInterface  repository.FlatRepositoryInterface
}

func (s *MyServer) GetDummyLogin(w http.ResponseWriter, r *http.Request, params generated.GetDummyLoginParams) {
	// Logger
	s.Logger.Info("Processing login request", "user_type", params.UserType)

	switch params.UserType {
	case "moderator":
		s.Logger.Info("Moderator login successful")
		w.Write([]byte("Moderator login successful"))

	case "user":
		s.Logger.Info("User login successful")
		w.Write([]byte("User login successful"))

	default:
		s.Logger.Error("Invalid user type", "user_type", params.UserType)
		http.Error(w, "Invalid user type", http.StatusBadRequest)
	}
}
