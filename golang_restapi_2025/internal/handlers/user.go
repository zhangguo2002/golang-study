package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/zhangguo2002/golangrestapi/internal/auth"
	"github.com/zhangguo2002/golangrestapi/internal/dtos"
	"github.com/zhangguo2002/golangrestapi/internal/middlewares"
	"github.com/zhangguo2002/golangrestapi/internal/store"
	"github.com/zhangguo2002/golangrestapi/internal/utils"
)

// cleanUserSession
func (h *Handler) cleanUserSession(userID string) error {
	pattern := fmt.Sprintf("session:%s:*", userID)
	ctx := context.Background()

	//scan to iterate over all the keys matching the pattern declared
	iter := h.Redis.Scan(ctx, 0, pattern, 0).Iterator()
	//loop through each key from redis
	for iter.Next(ctx) {
		//delete the key from redis
		err := h.Redis.Del(ctx, iter.Val()).Err()
		if err != nil {
			fmt.Printf("failed to delete session")
		}
	}
	if err := iter.Err(); err != nil {
		return err
	}
	return nil
}

func extractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return ""
	}
	return parts[1]
}

// logout
func (h *Handler) LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//extract the jwt claims from the context
		claims, ok := r.Context().Value(middlewares.UserClaimsKey).(*auth.Claims)
		if !ok {
			utils.RespondWithError(w, http.StatusBadRequest, "Please login to continue")
			return
		}
		//extract the token from the auth header
		tokenString := extractTokenFromHeader(r)
		if tokenString == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "missing token")
			return
		}

		//convert expireAt to time.Time
		expirationTime := time.Unix(claims.ExpiresAt, 0)
		now := time.Now()
		ttl := expirationTime.Sub(now)
		if ttl <= 0 {
			ttl = 5 * time.Minute
		}

		//blacklist the token in redis
		err := h.Redis.Set(r.Context(), tokenString, "blacklisted", ttl).Err()
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "failed to blacklist token")
			return
		}

		//clean user session in redis
		userIDStr := fmt.Sprintf("%d", claims.UserID)
		if err := h.cleanUserSession(userIDStr); err != nil {
			fmt.Printf("Error cleaning session for %s:%v\n", userIDStr, err)
		}
		utils.RespondWithSuccess(w, http.StatusOK, "Logged out successfully", true)
	}
}

// profile
func (h *Handler) UserProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(middlewares.UserClaimsKey).(*auth.Claims)
		if !ok {
			utils.RespondWithError(w, http.StatusBadRequest, "please login to continue")
			return
		}
		userID := claims.UserID

		//check the redis first
		cacheKey := fmt.Sprintf("user:%d", userID)
		if cached, err := h.Redis.Get(r.Context(), cacheKey).Result(); err != nil {
			var user store.User
			if err := json.Unmarshal([]byte(cached), &user); err == nil {
				utils.RespondWithSuccess(w, http.StatusOK, "success (from cache/redis)", user)
				return
			}
		}
		//fallback to db
		user, err := h.Queries.GetUserProfileByUserId(r.Context(), int64(userID))
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "user not found")
			return
		}
		//set to redis
		userJSON, _ := json.Marshal(user)
		h.Redis.Set(r.Context(), cacheKey, userJSON, 5*time.Minute)

		utils.RespondWithSuccess(w, http.StatusOK, "success", user)
	}
}

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
