package handlers

import (
	"conn/models"
	"conn/secure"
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Signup handles user registration
func Signup(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer r.Body.Close()
		// Hash the password before saving
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user.Password = string(hashedPassword)

		if err := db.Create(&user).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "User created successfully: %+v\n", user)
	}
}

// Login checks the user's email and password and returns JWT tokens if valid.
func Login(db *gorm.DB, cfg models.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqUser models.User
		if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var user models.User
		if err := db.Where("email = ?", reqUser.Email).First(&user).Error; err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password)); err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		accessToken, err := secure.GenerateToken(user.Email, user.Name, cfg.JWTSecret, time.Minute*15)
		if err != nil {
			http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
			return
		}

		refreshToken, err := secure.GenerateToken(user.Email, user.Name, cfg.JWTSecret, time.Hour*24*7)
		if err != nil {
			http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
			return
		}

		tokenResponse := models.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(tokenResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
