package handlers

import (
	"conn/secure"
	"fmt"
	"net/http"
)

// ProtectedHandler handles the protected route
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	email, name, ok := secure.GetUserFromContext(r.Context())
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}
	fmt.Println(email, name, ok)
	w.Write([]byte(fmt.Sprintf("This is a protected route. Hello, %s (%s)!", name, email)))
}
