package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zhangguo2002/golangrestapi/internal/dtos"
	"github.com/zhangguo2002/golangrestapi/internal/store"
	"github.com/zhangguo2002/golangrestapi/internal/utils"
)

func (h *Handler) CreateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//create context
		ctx := r.Context()
		//user request aka dto
		var req dtos.CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "error while hashing password")
			return
		}
		_, err = h.Queries.CreateUser(ctx, store.CreateUserParams{
			Username: req.Username,
			Email:    req.Email,
			Password: hashedPassword,
		})
		if err != nil {
			log.Printf("create user failed: %v", err)
			utils.RespondWithError(w, http.StatusInternalServerError, "error while creating user")
			return
		}
		utils.RespondWithSuccess(w, http.StatusCreated, "user created", req.Username)
	}
}
