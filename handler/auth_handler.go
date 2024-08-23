package handler

import (
	"encoding/json"
	"local_my_api/service"
	"local_my_api/validation"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		RespondWithJson(w, http.StatusBadRequest, ErrorResponse{
			Message: "Invalid request payload",
			Status:  "error",
		})
		return
	}

	token, err := service.LoginService(credentials.Email, credentials.Password)
	if err != nil {
		RespondWithJson(w, http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  "error",
		})
		return
	}

	response := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithJson(w, http.StatusBadRequest, ErrorResponse{
			Message: "Invalid request payload",
			Status:  "error",
		})
		return
	}

	if err := validation.ValidateRegisterInput(user.Email, user.Password, user.Name); err != nil {
		RespondWithJson(w, http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  "error",
		})
		return
	}

	err := service.RegisterService(user.Email, user.Password, user.Name)
	if err != nil {
		RespondWithJson(w, http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  "error",
		})
		return
	}

	RespondWithJson(w, http.StatusCreated, ErrorResponse{
		Message: "User registered successfully",
		Status:  "success",
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	var tokenRefreshRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&tokenRefreshRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newToken, err := service.RefreshTokenService(tokenRefreshRequest.RefreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := struct {
		Token string `json:"token"`
	}{
		Token: newToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
