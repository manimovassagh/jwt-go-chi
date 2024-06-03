package main

import (
	"conn/handlers"
	"conn/models"
	"conn/secure"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := models.Config{
		JWTSecret: "your_Sample_jwt_secret_key_that_should_be_very_secret__!!", // Replace with your secret key (not in production :)  and in that case you have to get it from env file or secrets)
	}

	dsn := "user:pass@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Database connected successfully")

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Post("/login", handlers.Login(db, cfg))
	r.Post("/signup", handlers.Signup(db))

	// Protected route example
	r.With(secure.VerifyJWT(cfg.JWTSecret)).Get("/protected", handlers.ProtectedHandler)

	log.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
}
