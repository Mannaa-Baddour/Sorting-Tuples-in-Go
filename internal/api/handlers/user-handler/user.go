package user_handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/api/models"
	user_model "github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/api/models/user-model"
	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/logging"
	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/server"
)

// NewUser is an instance of User defined in user_model, used in here to handle User object data.
var NewUser user_model.User

// templatePtr is a pointer towards the templates files, specified by the path created in server.
var templatePtr = server.TemplatePtr

// Data is a struct that holds the required data to be passed to the html template.
type Data struct {
	User         user_model.User
	ErrorMessage string
}

// LoginUser is a handler function that displays the login page with method is GET,
// and logs the user in when method is POST, after checking the validity of the user's data.
func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := templatePtr.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			logging.LogError(err, "user-handler/LoginUser : template execution at method GET")
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		data := Data{}
		err := r.ParseForm()
		if err != nil {
			logging.LogError(err, "user-handler/LoginUser : parsing form at method POST")
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		user, err := user_model.LoginUser(username, password)
		if err != nil {
			if err == sql.ErrNoRows {
				data.ErrorMessage = "Incorrect username or password"
				err := templatePtr.ExecuteTemplate(w, "index.html", data)
				if err != nil {
					logging.LogError(err, "user-handler/LoginUser : executing template at method POST, incorrect credentials condition")
					http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
					return
				}
			} else {
				logging.LogError(err, "user-handler/LoginUser : fetching data from DB at method POST, incorrect credentials condition")
				http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
				return
			}
		} else {
			url := fmt.Sprintf("/users/folders/?user_id=%d", user.ID)
			http.Redirect(w, r, url, http.StatusFound)
		}
	}
}

// CreateUser is a handler function that displays the signup page when method is GET,
// and creates a new user when method is POST, after a series of validation conditions.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := templatePtr.ExecuteTemplate(w, "signup.html", nil)
		if err != nil {
			logging.LogError(err, "user-handler/CreateUser : template execution at method GET")
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		data := Data{}
		err := r.ParseForm()
		if err != nil {
			logging.LogError(err, "user-handler/CreateUser : parsing form at method POST")
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
		newUser := &NewUser
		newUser.Username = r.FormValue("username")
		newUser.Email = r.FormValue("email")
		newUser.Password = r.FormValue("password")
		confirmPassword := r.FormValue("confirmPassword")
		if newUser.Password != confirmPassword {
			data.ErrorMessage = "Passwords don't match"
			err = templatePtr.ExecuteTemplate(w, "signup.html", data)
			if err != nil {
				logging.LogError(err, "user-handler/CreateUser : template execution at method POST, passwords don't match condition")
				http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		}
		if len(newUser.Password) < 8 {
			data.ErrorMessage = "Password is too short"
			err = templatePtr.ExecuteTemplate(w, "signup.html", data)
			if err != nil {
				logging.LogError(err, "user-handler/CreateUser : template execution at method POST, password length condition")
				http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		}
		err = newUser.CreateUser()
		if err != nil {
			if err.Error() == "username taken" || err.Error() == "email already exists" {
				data.ErrorMessage = err.Error()
				err = templatePtr.ExecuteTemplate(w, "signup.html", data)
				if err != nil {
					logging.LogError(err, "user-handler/CreateUser : template execution at method POST, user already exists condition")
					http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
					return
				}
			} else {
				logging.LogError(err, "user-handler/CreateUser : returned error while creating user, user already exists condition")
				http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
				return
			}
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

// GetUserByID is a handler function that gets user's info for the settings page.
// Associated with method GET.
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	if len(parameters) != 1 || !parameters.Has("user_id") {
		http.Error(w, "Error 400: Bad Request", http.StatusBadRequest)
		return
	}
	data := Data{}
	id := r.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Error 400: Bad Request", http.StatusBadRequest)
		return
	}
	user, err := user_model.GetUserByID(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Error 404: Page Not Found", http.StatusNotFound)
			return
		} else {
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	data.User = *user
	err = templatePtr.ExecuteTemplate(w, "settings.html", data)
	if err != nil {
		logging.LogError(err, "user-handler/GetUserByID : template execution")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// UpdateUser is a handler function that updates user info after validating them.
// Associated with method PUT.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		logging.LogError(err, "user-handler/UpdateUser : converting from string user_id to int64")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	user, err := user_model.GetUserByID(userId)
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	data := Data{}
	data.User = *user
	err = r.ParseForm()
	if err != nil {
		logging.LogError(err, "user-handler/UpdateUser : parsing form")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	oldPassword := r.FormValue("oldPassword")
	if oldPassword != user.Password {
		data.ErrorMessage = "incorrect password"
		err = templatePtr.ExecuteTemplate(w, "settings.html", data)
		if err != nil {
			logging.LogError(err, "user-handler/UpdateUser : template execution, incorrect password condition")
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}
	username := r.FormValue("username")
	if username != user.Username {
		usernameExists, err := models.CheckRecords("users", []string{"username"}, username)
		if err != nil {
			logging.LogError(err, "user-handler/UpdateUser : returned error while checking records for username")
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
		if usernameExists {
			data.ErrorMessage = "username taken"
			err = templatePtr.ExecuteTemplate(w, "settings.html", data)
			if err != nil {
				logging.LogError(err, "user-handler/UpdateUser : template execution, username taken condition")
				http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		} else {
			user.Username = username
		}
	}
	email := r.FormValue("email")
	if email != user.Email {
		emailExists, err := models.CheckRecords("users", []string{"email"}, email)
		if err != nil {
			logging.LogError(err, "user-handler/UpdateUser : returned error while checking records for email")
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
		if emailExists {
			data.ErrorMessage = "email already exists"
			err = templatePtr.ExecuteTemplate(w, "settings.html", data)
			if err != nil {
				logging.LogError(err, "user-handler/UpdateUser : template execution, email already exists condition")
				http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		} else {
			user.Email = email
		}
	}
	newPassword := r.FormValue("newPassword")
	if newPassword != "" {
		if len(newPassword) < 8 {
			data.ErrorMessage = "Password is too short"
			err = templatePtr.ExecuteTemplate(w, "settings.html", data)
			if err != nil {
				logging.LogError(err, "user-handler/UpdateUser : template execution, password is too short condition")
				http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		}
		if newPassword == user.Password {
			data.ErrorMessage = "Your old password cannot be your new password"
			err = templatePtr.ExecuteTemplate(w, "settings.html", data)
			if err != nil {
				logging.LogError(err, "user-handler/UpdateUser : template execution, matching old and new passwords condition")
				http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		}
		user.Password = newPassword
	}
	err = user.UpdateUser()
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	data.User = *user
	err = templatePtr.ExecuteTemplate(w, "settings.html", data)
	if err != nil {
		logging.LogError(err, "user-handler/UpdateUser : template execution")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// DeleteUser is a handler function that deletes a user and all its related folders and files.
// Associated with method DELETE
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logging.LogError(err, "user-handler/DeleteUser : converting from string user_id to int64")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	user, err := user_model.GetUserByID(userId)
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = user.DeleteUser()
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// HandleUser is a multipurpose handler function that executes a certain user related handler function
// based on the method of the request.
func HandleUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetUserByID(w, r)
	} else {
		err := r.ParseForm()
		if err != nil {
			logging.LogError(err, "user-handler/HandleUser : parsing form at POST")
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
		method := r.FormValue("_method")
		if method == "PUT" {
			UpdateUser(w, r)
		} else if method == "DELETE" {
			DeleteUser(w, r)
		}
	}
}
