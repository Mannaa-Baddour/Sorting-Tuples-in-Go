package routes

import (
	"net/http"

	file_handler "github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/api/handlers/file-handler"
	user_handler "github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/api/handlers/user-handler"
)

// RegisterRoutes prepares a pointer of ServeMux by connecting routes to certain handler function.
func RegisterRoutes(router *http.ServeMux) {
	// users routes //
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
	})
	router.HandleFunc("/users/login", user_handler.LoginUser)
	router.HandleFunc("/users/signup", user_handler.CreateUser)
	router.HandleFunc("/users/settings/", user_handler.HandleUser)

	// files routes //
	router.HandleFunc("/users/folders/", file_handler.HandleUserPage)
	router.HandleFunc("/users/folders/files/", file_handler.GetFiles)
	router.HandleFunc("/users/folders/files/file/", file_handler.HandleFile)
}
