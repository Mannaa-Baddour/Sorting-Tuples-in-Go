package file_handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/api/models"
	file_model "github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/api/models/file-model"
	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/logging"
	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/server"
)

// NewFile is an instance of File defined in file_model, used in here to handle File object data.
var NewFile file_model.File

// templatePtr is a pointer towards the templates files, specified by the path created in server.
var templatePtr = server.TemplatePtr

// Data is a struct that holds the required data to be passed to the html template.
type Data struct {
	Notification      string
	NotificationColor string
	InputExists       bool
	OutputExists      bool
	AccountID         string
	File              file_model.File
	Files             []file_model.File
}

// DisplayFolders is a handler function that displays folders (public, input, output) if they're not empty.
func DisplayFolders(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	parameters := r.URL.Query()
	if len(parameters) != 1 || !parameters.Has("user_id") {
		http.Error(w, "Error 400: Bad Request", http.StatusBadRequest)
		return
	}
	userID := parameters.Get("user_id")
	_, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		http.Error(w, "Error 400: Bad Request", http.StatusBadRequest)
		return
	}
	userExists, err := models.CheckRecords("users", []string{"id"}, userID)
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !userExists {
		http.Error(w, "Error 404: Page Not Found", http.StatusNotFound)
		return
	}
	data.AccountID = userID
	data.InputExists, err = file_model.DirectoryNotEmpty(data.AccountID, "input")
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	data.OutputExists, err = file_model.DirectoryNotEmpty(data.AccountID, "output")
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = templatePtr.ExecuteTemplate(w, "user_page.html", data)
	if err != nil {
		logging.LogError(err, "file-handler/DisplayFolders : executing template at method GET")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// CreateInputFile is a handler function that allows the user to create a new input file, that'll be
// added to their folder in order to be able to sort it later.
func CreateInputFile(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	userID := r.URL.Query().Get("user_id")
	data.AccountID = userID
	err := r.ParseForm()
	if err != nil {
		logging.LogError(err, "file-handler/CreateInputFile : parsing form")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	newFile := &NewFile
	fileName := r.FormValue("inputFile")
	fileNameExists, err := models.CheckRecords("files", []string{"name", "user_id", "type"}, fileName, userID, "input")
	if err != nil {
		logging.LogError(err, "file-handler/CreateInputFile : returned error while checking for file name")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	if fileNameExists {
		data.Notification = "File name already exists"
		data.NotificationColor = "color: red"
		data.InputExists, err = file_model.DirectoryNotEmpty(data.AccountID, "input")
		if err != nil {
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
		data.OutputExists, err = file_model.DirectoryNotEmpty(data.AccountID, "output")
		if err != nil {
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = templatePtr.ExecuteTemplate(w, "user_page.html", data)
		if err != nil {
			logging.LogError(err, "file-handler/CreateInputFile : executing template, file name already exists condition")
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}
	newFile.Name = fileName
	newFile.Data = file_model.ParseText(r.FormValue("data"))
	userId, err := strconv.ParseInt(data.AccountID, 10, 64)
	if err != nil {
		logging.LogError(err, "file-handler/CreateInputFile : converting from string user_id to int64")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	newFile.UserID = userId
	newFile.Type = "input"
	err = newFile.CreateUserFile()
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	data.Notification = "file created successfully"
	data.NotificationColor = "color: green"
	data.InputExists, err = file_model.DirectoryNotEmpty(data.AccountID, "input")
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	data.OutputExists, err = file_model.DirectoryNotEmpty(data.AccountID, "output")
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = templatePtr.ExecuteTemplate(w, "user_page.html", data)
	if err != nil {
		logging.LogError(err, "file-handler/CreateInputFile : executing template, after file creation")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// HandleUserPage is a multipurpose handler function that executes a certain user page related handler function
// based on the method of the request.
func HandleUserPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		DisplayFolders(w, r)
	} else if r.Method == "POST" {
		CreateInputFile(w, r)
	}
}

// GetFiles is a handler function that displays the files related to a certain category (folder)
// either public, input, or output.
func GetFiles(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	if len(parameters) != 3 || !parameters.Has("user_id") || !parameters.Has("file_type") || !parameters.Has("account_id") {
		http.Error(w, "Error 400: Bad Request", http.StatusBadRequest)
		return
	}
	userID := parameters.Get("user_id")
	userExists, err := models.CheckRecords("users", []string{"id"}, userID)
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !userExists {
		http.Error(w, "Error 404: Page Not Found", http.StatusNotFound)
		return
	}
	userId, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		http.Error(w, "Error 400: Bad Request", http.StatusBadRequest)
		return
	}
	fileType := parameters.Get("file_type")
	if fileType != "input" && fileType != "output" {
		http.Error(w, "Error 400: Bad Request", http.StatusBadRequest)
		return
	}
	files, err := file_model.GetFiles(userId, fileType)
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	accountId := parameters.Get("account_id")
	accountExists, err := models.CheckRecords("users", []string{"id"}, accountId)
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !accountExists {
		http.Error(w, "Error 404: Page Not Found", http.StatusNotFound)
		return
	}
	_, err = strconv.ParseInt(accountId, 10, 64)
	if err != nil {
		http.Error(w, "Error 400: Bad Request", http.StatusBadRequest)
		return
	}
	data := Data{
		Files:     files,
		AccountID: accountId,
	}
	err = templatePtr.ExecuteTemplate(w, "files.html", data)
	if err != nil {
		logging.LogError(err, "file-handler/GetUserFiles : executing template")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// GetFileByID is a handler function that get the info of the file specified by user and file IDs.
// Associated with method GET.
func GetFileByID(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	if len(parameters) != 3 || !parameters.Has("user_id") || !parameters.Has("file_id") || !parameters.Has("account_id") {
		http.Error(w, "Error 400: Bad Request", http.StatusBadRequest)
		return
	}
	userID := r.URL.Query().Get("user_id")
	userExists, err := models.CheckRecords("users", []string{"id"}, userID)
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !userExists {
		http.Error(w, "Error 404: Page Not Found", http.StatusNotFound)
		return
	}
	userId, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		http.Error(w, "Error 400: Bad Request", http.StatusBadRequest)
		return
	}
	fileID := r.URL.Query().Get("file_id")
	fileExists, err := models.CheckRecords("files", []string{"id"}, fileID)
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !fileExists {
		http.Error(w, "Error 404: Page Not Found", http.StatusNotFound)
		return
	}
	fileId, err := strconv.ParseInt(fileID, 10, 64)
	if err != nil {
		http.Error(w, "Error 400: Bad Request", http.StatusBadRequest)
		return
	}
	file, err := file_model.GetFileByID(userId, fileId)
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	accountId := r.URL.Query().Get("account_id")
	accountExists, err := models.CheckRecords("users", []string{"id"}, accountId)
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !accountExists {
		http.Error(w, "Error 404: Page Not Found", http.StatusNotFound)
		return
	}
	_, err = strconv.ParseInt(accountId, 10, 64)
	if err != nil {
		http.Error(w, "Error 400: Bad Request", http.StatusBadRequest)
		return
	}
	data := Data{
		File:      *file,
		AccountID: accountId,
	}
	err = templatePtr.ExecuteTemplate(w, "file.html", data)
	if err != nil {
		logging.LogError(err, "file-handler/GetFileByID : executing template")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// UpdateUserFile is a handler function that modifies the current user file if it's of type input.
// Associated with method PUT.
func UpdateUserFile(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	userID := r.URL.Query().Get("user_id")
	data.AccountID = userID
	userId, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		logging.LogError(err, "file-handler/UpdateUserFile : converting from string user_id to int64")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	fileID := r.URL.Query().Get("file_id")
	fileId, err := strconv.ParseInt(fileID, 10, 64)
	if err != nil {
		logging.LogError(err, "file-handler/UpdateUserFile : converting from string file_id to int64")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	file, err := file_model.GetFileByID(userId, fileId)
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = r.ParseForm()
	if err != nil {
		logging.LogError(err, "file-handler/UpdateUserFile : parsing form")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	oldFileName := file.Name
	fileName := r.FormValue("fileName")
	if fileName != "" && fileName != file.Name {
		fileNameExists, err := models.CheckRecords("files", []string{"name", "user_id", "type"}, fileName, userID, "input")
		if err != nil {
			logging.LogError(err, "file-handler/UpdateUserFile : returned error while checking for file name")
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
		if fileNameExists {
			data.Notification = "File name already exists"
			data.NotificationColor = "color: red"
			data.File = *file
			err = templatePtr.ExecuteTemplate(w, "file.html", data)
			if err != nil {
				logging.LogError(err, "file-handler/UpdateUserFile : executing template, file name already exists condition")
				http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		}
		file.Name = fileName
	}
	formFileData := r.FormValue("data")
	fileData := file_model.ParseText(formFileData)
	file.Data = fileData
	err = file.UpdateUserFile(oldFileName)
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	data.File = *file
	err = templatePtr.ExecuteTemplate(w, "file.html", data)
	if err != nil {
		logging.LogError(err, "file-handler/UpdateUserFile : executing template")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// DeleteUserFile is a handler function that deletes the current file if it's not a public file.
// Associated with method DELETE.
func DeleteUserFile(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		logging.LogError(err, "file-handler/DeleteUserFile : converting from string user_id to int64")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	fileID := r.URL.Query().Get("file_id")
	fileId, err := strconv.ParseInt(fileID, 10, 64)
	if err != nil {
		logging.LogError(err, "file-handler/DeleteUserFile : converting from string file_id to int64")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = file_model.DeleteUserFile(userId, fileId)
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	url := fmt.Sprintf("/users/folders/?user_id=%s", userID)
	http.Redirect(w, r, url, http.StatusFound)
}

// SortFileRows is a handler function that executes sorting functionality on an input file
// and outputs the result in a file specified by the user.
func SortFileRows(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	userID := r.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		logging.LogError(err, "file-handler/SortFileRows : converting from string user_id to int64")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	fileID := r.URL.Query().Get("file_id")
	fileId, err := strconv.ParseInt(fileID, 10, 64)
	if err != nil {
		logging.LogError(err, "file-handler/SortFileRows : converting from string file_id to int64")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	accountID := r.URL.Query().Get("account_id")
	data.AccountID = accountID
	accountId, err := strconv.ParseInt(accountID, 10, 64)
	if err != nil {
		logging.LogError(err, "file-handler/SortFileRows : converting from string account_id to int64")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	data.AccountID = accountID
	err = r.ParseForm()
	if err != nil {
		logging.LogError(err, "file-handler/SortFileRows : parsing form")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	outputFile := r.FormValue("outputFile")
	fileNameExists, err := models.CheckRecords("files", []string{"name", "user_id", "type"}, outputFile, accountID, "output")
	if err != nil {
		logging.LogError(err, "file-handler/SortFileRows : returned error while checking for file name")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	if fileNameExists {
		data.Notification = "File name already exists"
		data.NotificationColor = "color: red"
		file, err := file_model.GetFileByID(userId, fileId)
		if err != nil {
			logging.LogError(err, "file-handler/SortFileRows : returned error while getting file, file name already exists condition")
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
		data.File = *file
		err = templatePtr.ExecuteTemplate(w, "file.html", data)
		if err != nil {
			logging.LogError(err, "file-handler/SortFileRows : executing template, file name already exists condition")
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}
	sortingColumn := r.FormValue("sortingColumn")
	column, err := strconv.Atoi(sortingColumn)
	if err != nil {
		logging.LogError(err, "file-handler/SortFileRows : converting from string sorting column to int")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	if column > 2 || column < 0 {
		data.Notification = "sorting column must be either 0, 1, or 2"
		data.NotificationColor = "color: red"
		err = templatePtr.ExecuteTemplate(w, "file.html", data)
		if err != nil {
			logging.LogError(err, "file-handler/SortFileRows : executing template, sorting column out of bounds condition")
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	file, err := file_model.SortFileRows(userId, fileId, column, outputFile, accountId)
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	data.File = *file
	err = templatePtr.ExecuteTemplate(w, "file.html", data)
	if err != nil {
		logging.LogError(err, "file-handler/SortFileRows : executing template")
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// HandleFile is a multipurpose handler function that executes a certain file related handler function
// based on the method of the request.
func HandleFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetFileByID(w, r)
	} else {
		err := r.ParseForm()
		if err != nil {
			logging.LogError(err, "file-handler/HandleFile : parsing form at POST")
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			return
		}
		method := r.FormValue("_method")
		if method == "PUT" {
			UpdateUserFile(w, r)
		} else if method == "DELETE" {
			DeleteUserFile(w, r)
		} else {
			SortFileRows(w, r)
		}
	}
}
