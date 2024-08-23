package handler

import (
	"encoding/json"
	"local_my_api/model"
	"local_my_api/service"
	"local_my_api/utils"
	"local_my_api/validation"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUserListHandler(w http.ResponseWriter, r *http.Request) {
	var userList = []model.User{}

	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	pagination := utils.ParsePagination(pageStr, sizeStr)

	userList, totalUsers, status, err := service.GetUserListService(pagination)
	if err != nil {
		http.Error(w, err.Error(), status)
	}

	response := struct {
		Users    []model.User `json:"users"`
		Total    int64        `json:"total"`
		Page     int          `json:"page"`
		PageSize int          `json:"pageSize"`
	}{
		Users:    userList,
		Total:    totalUsers,
		Page:     pagination.Page,
		PageSize: pagination.Size,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, status, err := service.GetUserService(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	userResponse := toUserResponse(user)
	response, err := json.Marshal(userResponse)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		RespondWithJson(w, http.StatusInternalServerError, ErrorResponse{
			Message: "Error marshaling response",
			Status:  "error",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paramsId := vars["id"]

	userID, ok := r.Context().Value(utils.UserIDKey).(string)
	if !ok {
		RespondWithJson(w, http.StatusInternalServerError, ErrorResponse{
			Message: "Error getting user ID from context",
			Status:  "error",
		})
		return
	}

	if userID != paramsId {
		RespondWithJson(w, http.StatusForbidden, ErrorResponse{
			Message: "You are not authorized to update this user",
			Status:  "error",
		})
		return
	}

	var userInput struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		RespondWithJson(w, http.StatusBadRequest, ErrorResponse{
			Message: "Invalid request payload",
			Status:  "error",
		})
		return
	}

	if err := validation.ValidateUpdateUserInput(userInput.Name); err != nil {
		RespondWithJson(w, http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  "error",
		})
		return
	}

	user, status, err := service.UpdateUserService(paramsId, userInput.Name)
	if err != nil {
		RespondWithJson(w, status, ErrorResponse{
			Message: err.Error(),
			Status:  "error",
		})
		return
	}

	userResponse := toUserResponse(user)
	response, err := json.Marshal(userResponse)
	if err != nil {
		RespondWithJson(w, http.StatusInternalServerError, ErrorResponse{
			Message: "Error marshaling response",
			Status:  "error",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paramsId := vars["id"]

	userID, ok := r.Context().Value(utils.UserIDKey).(string)
	if !ok {
		RespondWithJson(w, http.StatusInternalServerError, ErrorResponse{
			Message: "Error getting user ID from context",
			Status:  "error",
		})
		return
	}

	if userID != paramsId {
		RespondWithJson(w, http.StatusForbidden, ErrorResponse{
			Message: "You are not authorized to delete this user",
			Status:  "error",
		})
		return
	}

	status, err := service.DeleteUserService(paramsId)
	if err != nil {
		RespondWithJson(w, status, ErrorResponse{
			Message: err.Error(),
			Status:  "error",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func toUserResponse(user *model.User) *UserResponse {
	return &UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
