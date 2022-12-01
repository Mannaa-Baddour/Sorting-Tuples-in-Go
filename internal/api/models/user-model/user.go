package user_model

import (
	"database/sql"
	"errors"
	"os"
	"strconv"

	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/api/models"
	file_model "github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/api/models/file-model"
	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/logging"
	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/server"
	_ "github.com/lib/pq"
)

// db is a variable the holds the reference to the database connection made in models package.
var db = models.DB

// User is a struct that forms the fields of the users table in the database, adding to it
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// createDirectory creates a new directory named as user id on user creation,
// and then creates subdirectories within it which are input and output to contain
// the user's input and output files respectively.
func createDirectory(userId int64) error {
	err := os.Chdir(server.Path)
	if err != nil {
		logging.LogError(err, "user-model/createDirectory : returned error while changing directory")
		return err
	}
	userFolder := strconv.FormatInt(userId, 10)
	err = os.Mkdir(userFolder, os.ModePerm)
	if err != nil {
		logging.LogError(err, "user-model/createDirectory : returned error while making a directory")
		return err
	}
	err = os.Chdir(server.Path + userFolder)
	if err != nil {
		logging.LogError(err, "user-model/createDirectory : returned error while changing directory to subfolder")
		return err
	}
	for _, folder := range []string{"input", "output"} {
		err = os.Mkdir(folder, os.ModePerm)
		if err != nil {
			logging.LogError(err, "user-model/createDirectory : returned error while making subfolder")
			return err
		}
	}
	return nil
}

// deleteDirectory deletes all user related folders and files upon user account deletion.
func deleteDirectory(userId int64) error {
	err := os.Chdir(server.Path)
	if err != nil {
		logging.LogError(err, "user-model/deleteDirectory : returned error while changing directory")
		return err
	}
	target := strconv.FormatInt(userId, 10)
	err = os.RemoveAll(target)
	if err != nil {
		logging.LogError(err, "user-model/deleteDirectory : returned error while removing directory")
		return err
	}
	return nil
}

// LoginUser checks if the user exists in the database or not, to decide whether or not the user
// should be logged in or asked to signup.
func LoginUser(username, password string) (*User, error) {
	var user User
	statement, err := db.Prepare("SELECT * FROM users WHERE username = $1 AND password = $2")
	if err != nil {
		logging.LogError(err, "user-model/LoginUser : returned error while preparing statement")
		return nil, err
	}
	row := statement.QueryRow(username, password)
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser is a User associated method which creates a new user by inserting its data to the database
// and then creating user's directories.
func (user *User) CreateUser() error { //*User {
	usernameExists, err := models.CheckRecords("users", []string{"username"}, user.Username)
	if err != nil {
		logging.LogError(err, "user-model/CreateUser: returned error while checking the database for the username")
		return err
	}
	if usernameExists {
		return errors.New("username taken")
	}
	emailExists, err := models.CheckRecords("users", []string{"email"}, user.Email)
	if err != nil {
		logging.LogError(err, "user-model/CreateUser: returned error while checking the database for user email")
		return err
	}
	if emailExists {
		return errors.New("email already exists")
	}
	insertStatement, err := db.Prepare("INSERT INTO users(username, email, password) VALUES($1, $2, $3)")
	if err != nil {
		logging.LogError(err, "user-model/CreateUser : returned error while preparing statement")
		return err
	}
	defer insertStatement.Close()
	_, err = insertStatement.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		logging.LogError(err, "user-model/CreateUser : returned error while executing statement")
		return err
	}
	var userId int64
	err = db.QueryRow("SELECT id FROM users WHERE username = $1", user.Username).Scan(&userId)
	if err != nil {
		logging.LogError(err, "user-model/CreateUser : returned error while scanning query statement")
		return err
	}
	user.ID = userId
	err = createDirectory(userId)
	if err != nil {
		_, execErr := db.Exec("DELETE FROM users WHERE id = $1", user.ID)
		if execErr != nil {
			logging.LogError(execErr, "user-model/CreateUser : returned error while deleting user")
			return execErr
		}
		logging.LogError(err, "user-model/CreateUser : returned error while creating user directory")
		return err
	}
	return nil
}

// GetUserByID searches the database for a user based on its ID, queries the fields
// and returns a User object with those fields as values.
func GetUserByID(id int64) (*User, error) {
	var user User
	statement, err := db.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		logging.LogError(err, "user-model/GetUserByID : returned error while preparing statement")
		return nil, err
	}
	defer statement.Close()
	row := statement.QueryRow(id)
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			logging.LogError(err, "user-model/GetUserByID : returned error while scanning query result")
			return nil, err
		}
	}
	return &user, nil
}

// UpdateUser is a User associated method which updates the fields username, email, and/or password
// for the same already existing user.
func (user *User) UpdateUser() error { //*User {
	statement, err := db.Prepare("UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4")
	if err != nil {
		logging.LogError(err, "user-model/UpdateUser : returned error while preparing statement")
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		logging.LogError(err, "user-model/UpdateUser : returned error while executing statement")
		return err
	}
	return nil
}

// DeleteUser is a User associated method which deletes user from the database
// and then calls deleteDirectory to delete the folder associated to that user.
func (user *User) DeleteUser() error { //*User {
	retrieveStatement, err := db.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		logging.LogError(err, "user-model/DeleteUser : returned error while preparing query statement")
		return err
	}
	defer retrieveStatement.Close()
	row := retrieveStatement.QueryRow(user.ID)
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		} else {
			logging.LogError(err, "user-model/DeleteUser : returned error while scanning query result")
			return err
		}
	}
	userInputFiles, err := file_model.GetFiles(user.ID, "input")
	if err != nil {
		logging.LogError(err, "user-model/DeleteUser : returned error while retrieving input files")
		return err
	}
	userOutputFiles, err := file_model.GetFiles(user.ID, "output")
	if err != nil {
		logging.LogError(err, "user-model/DeleteUser : returned error while retrieving output files")
		return err
	}
	deleteFilesStatement, err := db.Prepare("DELETE FROM files WHERE user_id = $1")
	if err != nil {
		logging.LogError(err, "user-model/DeleteUser : returned error while preparing delete files statement")
		return err
	}
	defer deleteFilesStatement.Close()
	_, err = deleteFilesStatement.Exec(user.ID)
	if err != nil {
		logging.LogError(err, "user-model/DeleteUser : returned error while executing delete files statement")
		return err
	}
	deleteStatement, err := db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		logging.LogError(err, "user-model/DeleteUser : returned error while preparing delete user statement")
		return err
	}
	defer deleteStatement.Close()
	_, err = deleteStatement.Exec(user.ID)
	if err != nil {
		for _, file := range userInputFiles {
			insertStatement, insertErr := db.Prepare("INSERT INTO files(id, name, type, user_id) VALUES ($1, $2, $3, $4)")
			if insertErr != nil {
				logging.LogError(err, "user-model/DeleteUser : returned error while preparing insert input files statement")
				break
			}
			defer insertStatement.Close()
			_, insertExecErr := insertStatement.Exec(file.ID, file.Name, file.Type, file.UserID)
			if insertExecErr != nil {
				logging.LogError(err, "user-model/DeleteUser : returned error while executing insert input files statement")
				break
			}
		}
		for _, file := range userOutputFiles {
			insertStatement, insertErr := db.Prepare("INSERT INTO files(id, name, type, user_id) VALUES ($1, $2, $3, $4)")
			if insertErr != nil {
				logging.LogError(err, "user-model/DeleteUser : returned error while preparing insert output files statement")
				break
			}
			defer insertStatement.Close()
			_, insertExecErr := insertStatement.Exec(file.ID, file.Name, file.Type, file.UserID)
			if insertExecErr != nil {
				logging.LogError(err, "user-model/DeleteUser : returned error while executing insert output files statement")
				break
			}
		}
		logging.LogError(err, "user-model/DeleteUser : returned error while executing delete user statement")
		return err
	}
	err = deleteDirectory(user.ID)
	if err != nil {
		_ = user.CreateUser()
		return err
	}
	return nil
}
