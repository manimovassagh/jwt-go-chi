package secure

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type contextKey string

const (
	userEmailContextKey contextKey = "userEmail"
	userNameContextKey  contextKey = "userName"
)

// GenerateToken generates a JWT token
func GenerateToken(email, name, secret string, expiration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"name":  name,
		"exp":   jwt.NewNumericDate(time.Now().Add(expiration)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// VerifyJWT is a middleware that verifies the JWT token in the Authorization header and stores user info in context
func VerifyJWT(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header missing", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				http.Error(w, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
				return
			}

			token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				email, emailOk := claims["email"].(string)
				name, nameOk := claims["name"].(string)
				if !emailOk || !nameOk {
					http.Error(w, "Invalid token claims", http.StatusUnauthorized)
					return
				}
				ctx := context.WithValue(r.Context(), userEmailContextKey, email)
				ctx = context.WithValue(ctx, userNameContextKey, name)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}
		})
	}
}

// GetUserFromContext retrieves the user information from the context
func GetUserFromContext(ctx context.Context) (string, string, bool) {
	email, emailOk := ctx.Value(userEmailContextKey).(string)
	name, nameOk := ctx.Value(userNameContextKey).(string)
	return email, name, emailOk && nameOk
}
