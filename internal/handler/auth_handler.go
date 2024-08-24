package handler

import (
	"encoding/json"
	"local_my_api/internal/services"
	"local_my_api/internal/validation"
	"local_my_api/pkg/utils"
	"net/http"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{authService: service}
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	token, err := h.authService.LoginService(credentials.Email, credentials.Password)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	utils.RespondWithJson(w, http.StatusOK, response)
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := validation.ValidateRegisterInput(user.Email, user.Password, user.Name); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := h.authService.RegisterService(user.Email, user.Password, user.Name)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJson(w, http.StatusCreated, utils.Response{
		Message: "User registered successfully",
		Status:  "success",
	})
}

func (h *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (h *AuthHandler) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	var tokenRefreshRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&tokenRefreshRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newToken, err := h.authService.RefreshTokenService(tokenRefreshRequest.RefreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := struct {
		Token string `json:"token"`
	}{
		Token: newToken,
	}

	utils.RespondWithJson(w, http.StatusOK, response)
}
