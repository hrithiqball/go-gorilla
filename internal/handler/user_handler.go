package handler

import (
	"encoding/json"
	"local_my_api/internal/models"
	"local_my_api/internal/services"
	"local_my_api/internal/validation"
	"local_my_api/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type UserHandler interface {
	GetUserListHandler(w http.ResponseWriter, r *http.Request)
	GetUserHandler(w http.ResponseWriter, r *http.Request)
	GetUserBusinessHandler(w http.ResponseWriter, r *http.Request)
	UpdateUserHandler(w http.ResponseWriter, r *http.Request)
	DeleteUserHandler(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService services.UserService
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserHandler(service services.UserService) UserHandler {
	return &userHandler{userService: service}
}

func (h *userHandler) GetUserListHandler(w http.ResponseWriter, r *http.Request) {
	var userList = []models.User{}

	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	pagination := utils.ParsePagination(pageStr, sizeStr)

	userList, totalUsers, err := h.userService.GetUserListService(pagination)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	userListResponse := []UserResponse{}
	for _, user := range userList {
		userListResponse = append(userListResponse, *toUserResponse(&user))
	}

	response := struct {
		UserList []UserResponse `json:"userList"`
		Total    int64          `json:"total"`
		Page     int            `json:"page"`
		PageSize int            `json:"pageSize"`
	}{
		UserList: userListResponse,
		Total:    totalUsers,
		Page:     pagination.Page,
		PageSize: pagination.Size,
	}

	utils.RespondWithJson(w, http.StatusOK, response)
}

func (h *userHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, err := h.userService.GetUserService(vars["id"])
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ResponseWithError(w, http.StatusNotFound, "User not found")
			return
		}

		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	userResponse := toUserResponse(user)
	utils.RespondWithJson(w, http.StatusOK, userResponse)
}

func (h *userHandler) GetUserBusinessHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	businessList, err := h.userService.GetUserBusinessService(ID)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJson(w, http.StatusOK, businessList)
}

func (h *userHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	userID, ok := r.Context().Value(utils.UserIDKey).(string)
	if !ok {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error getting user ID from context")
		return
	}

	if userID != ID {
		utils.ResponseWithError(w, http.StatusForbidden, "You are not authorized to update this user")
		return
	}

	var userInput struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := validation.ValidateUpdateUserInput(userInput.Name); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.userService.UpdateUserService(ID, &models.UserUpdate{
		Name: userInput.Name,
	})
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	userResponse := toUserResponse(user)
	utils.RespondWithJson(w, http.StatusOK, userResponse)
}

func (h *userHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paramsId := vars["id"]

	userID, ok := r.Context().Value(utils.UserIDKey).(string)
	if !ok {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error getting user ID from context")
		return
	}

	if userID != paramsId {
		utils.ResponseWithError(w, http.StatusForbidden, "You are not authorized to delete this user")
		return
	}

	err := h.userService.DeleteUserService(paramsId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ResponseWithError(w, http.StatusNotFound, "User not found")
			return
		}

		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJson(w, http.StatusOK, utils.Response{Message: "User deleted successfully", Status: "success"})
}

func toUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
