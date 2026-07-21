package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/zhangguo2002/golangrestapi/internal/auth"
	"github.com/zhangguo2002/golangrestapi/internal/dtos"
	"github.com/zhangguo2002/golangrestapi/internal/store"
	"github.com/zhangguo2002/golangrestapi/internal/utils"
)

// login a user
func (h *Handler) LoginUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req dtos.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		}
		//Validate the request
		if err := utils.Validate(&req); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		}
		//fetch the user from the db using the store queries
		user, err := h.Queries.GetUserByUsernameOrEmail(ctx, req.Username)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "invalid credential")
			return
		}
		if !utils.ComparePassword(user.Password, req.Password) {
			utils.RespondWithError(w, http.StatusUnauthorized, "invalid credential")
			return
		}
		jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
		token, err := auth.GenerateJWT(int64(user.ID), user.Username, jwtKey)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error generating a token")
			return
		}
		utils.RespondWithSuccess(w, http.StatusOK, "Login successful", map[string]string{
			"token": token,
		})
	}
}

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
