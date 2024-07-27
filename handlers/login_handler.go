package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"quote/data"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	IsSuccess bool   `json:"is_success"`
	Message   string `json:"message"`
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest

	json.NewDecoder(r.Body).Decode(&loginRequest)

	fmt.Printf("Username: %+v \n", loginRequest.Username)
	fmt.Printf("Password: %+v \n", loginRequest.Password)

	var loginResponse LoginResponse

	userRep := data.NewUserRepository(h.db)
	err := userRep.Checkuser(loginRequest.Username, loginRequest.Password)

	if err == nil {
		loginResponse = LoginResponse{
			IsSuccess: true,
			Message:   "",
		}
	} else {
		loginResponse = LoginResponse{
			IsSuccess: false,
			Message:   "Username or Password is invalid",
		}
	}

	result, _ := json.Marshal(loginResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
