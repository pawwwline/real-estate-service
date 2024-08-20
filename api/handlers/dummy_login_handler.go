package handlers

import (
	"log/slog"
	"net/http"
	"real-estate-service/api/generated"
	"real-estate-service/internal/services/auth"
	"real-estate-service/internal/utils"
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

	var message string

	switch params.UserType {
	case "moderator":
		s.Logger.Info("Moderator login successful")
		token, err := auth.CreateToken("moderator")
		if err != nil {
			s.Logger.Error("Error creating token", "error", err)
			utils.InternalServerError(w, r, err.Error())
		}
		message = "Успешная аутентификация Token: " + token

	case "client":
		s.Logger.Info("User login successful")
		token, err := auth.CreateToken("user")
		if err != nil {
			s.Logger.Error("Error creating token", "error", err)
			utils.InternalServerError(w, r, err.Error())
		}
		message = "Успешная аутентификация Token: " + token

	default:
		s.Logger.Error("Invalid user type", "user_type", params.UserType)
		utils.BadRequest(w, r)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
