package http

import (
    "encoding/json"
    "net/http"
    "github.com/shirocola/user-service/internal/domain"
    "github.com/shirocola/user-service/internal/usecase"

    "github.com/gorilla/mux"
)

type UserHandler struct {
    UserUsecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
    return &UserHandler{
        UserUsecase: u,
    }
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    user, err := h.UserUsecase.GetUserByID(params["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user domain.User
    json.NewDecoder(r.Body).Decode(&user)
    err := h.UserUsecase.CreateUser(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(user)
}
