package main

import (
	"conn/database"
	"conn/handlers"
	"conn/models"
	"conn/secure"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := models.Config{
		JWTSecret: "your_Sample_jwt_secret_key_that_should_be_very_secret__!!", // Replace with your secret key
	}

	db := database.InitDB()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Post("/login", handlers.Login(db, cfg))
	r.Post("/signup", handlers.Signup(db))

	r.With(secure.VerifyJWT(cfg.JWTSecret)).Get("/protected", handlers.ProtectedHandler)

	log.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
}
