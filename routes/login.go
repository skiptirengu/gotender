package routes

import (
	"net/http"
	"encoding/json"
	"github.com/skiptirengu/gotender/models"
	"github.com/skiptirengu/gotender/util"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (LoginRequest) HandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.NewHttpError(w, http.StatusUnprocessableEntity)
		return
	}

	var user *models.User
	if user = models.FindUserByEmailOrUsername(request.Email, request.Username); user == nil {
		util.NewHttpError(w, http.StatusNotFound)
		return
	}

	if err := user.ValidatePassword(request.Password); err != nil {
		util.NewHttpError(w, http.StatusBadRequest)
		return
	}

	if token, err := models.NewToken(user.Id); err != nil {
		util.NewHttpError(w, http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(token)
	}
}
