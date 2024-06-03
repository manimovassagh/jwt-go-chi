package models

// User defines the user model
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type:varchar(100);uniqueIndex"`
	Password string `gorm:"type:varchar(100)"`
}

// Config for JWT secret
type Config struct {
	JWTSecret string
}

// TokenResponse represents the response containing the tokens
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

