package file_model

import (
	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/logging"
	_ "github.com/lib/pq"

	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/api/models"
	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/sorting"
)

// db is a variable the holds the reference to the database connection made in models package.
var db = models.DB

// File is a struct that forms the fields of the files table in the database, adding to it
// Data field that is required later to display the contents of the file.
type File struct {
	ID     int64            `json:"id"`
	Name   string           `json:"name"`
	Type   string           `json:"type"`
	Data   []models.Content `json:"data"`
	UserID int64            `json:"user_id"`
}

// GetFiles is a function that queries all files that belong to a certain user and falls under a certain
// file type, then returns a list of those files or error otherwise.
func GetFiles(userId int64, fileType string) ([]File, error) {
	var files []File
	statement, err := db.Prepare("SELECT id, name, type, user_id FROM files WHERE user_id = $1 AND type = $2")
	if err != nil {
		logging.LogError(err, "file-model/GetFiles : preparing statement")
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.Query(userId, fileType)
	if err != nil {
		logging.LogError(err, "file-model/GetFiles : querying statement")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var file File
		err = rows.Scan(&file.ID, &file.Name, &file.Type, &file.UserID)
		if err != nil {
			logging.LogError(err, "file-model/GetFiles : scanning rows")
			return nil, err
		}
		file.Data, err = readFile(userId, file.Name, file.Type)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}

// GetFileByID is a function that queries a file info based on its ID and the user ID it belongs to,
// then returns the file with the data it holds, or error otherwise.
func GetFileByID(userId, fileId int64) (*File, error) {
	file := File{
		ID:     fileId,
		UserID: userId,
	}
	statement, err := db.Prepare("SELECT name, type FROM files WHERE user_id = $1 AND id = $2")
	if err != nil {
		logging.LogError(err, "file-model/GetFileByID : preparing statement")
		return nil, err
	}
	defer statement.Close()
	row := statement.QueryRow(userId, fileId)
	err = row.Scan(&file.Name, &file.Type)
	if err != nil {
		logging.LogError(err, "file-model/GetFileByID : scanning row")
		return nil, err
	}
	file.Data, err = readFile(userId, file.Name, file.Type)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

// CreateUserFile is a method associated with the struct File,
// it adds this file to the database and calls createFile function to create it.
func (file *File) CreateUserFile() error {
	statement, err := db.Prepare("INSERT INTO files (name, type, user_id) VALUES ($1, $2, $3)")
	if err != nil {
		logging.LogError(err, "file-model/CreateUserFile : preparing statement")
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(file.Name, file.Type, file.UserID)
	if err != nil {
		logging.LogError(err, "file-model/CreateUserFile : executing statement")
		return err
	}
	err = createFile(file.UserID, file.Name, file.Type, file.Data)
	if err != nil {
		_, execErr := db.Exec("DELETE FROM files WHERE id = $1", file.ID)
		if execErr != nil {
			logging.LogError(execErr, "file-model/CreateUserFile : executing statement at returned error from creating the file")
			return execErr
		}
		return err
	}
	getFileIDStatement, err := db.Prepare("SELECT id FROM files WHERE name=$1 AND type=$2 AND user_id=$3")
	if err != nil {
		logging.LogError(err, "file-model/CreateUserFile : returned error while preparing get file id statement")
		return err
	}
	defer getFileIDStatement.Close()
	err = getFileIDStatement.QueryRow(file.Name, file.Type, file.UserID).Scan(&file.ID)
	if err != nil {
		logging.LogError(err, "file-model/CreateUserFile : returned error while scanning get file id row")
		return err
	}
	return nil
}

// UpdateUserFile is a method associated with the struct File, it modifies some changed info
// related to the file, and/or its content in the csv file.
func (file *File) UpdateUserFile(oldFileName string) error {
	statement, err := db.Prepare("UPDATE files SET name = $1 WHERE id = $2")
	if err != nil {
		logging.LogError(err, "file-model/UpdateUserFile : preparing statement")
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(file.Name, file.ID)
	if err != nil {
		logging.LogError(err, "file-model/UpdateUserFile : executing statement")
		return err
	}
	err = updateFile(file.UserID, file.Name, oldFileName, file.Type, file.Data)
	if err != nil {
		_, execErr := db.Exec("UPDATE files SET name = $1 WHERE id = $2", oldFileName, file.ID)
		if execErr != nil {
			logging.LogError(execErr, "file-model/UpdateUserFile : executing statement at returned error from updating the file")
			return execErr
		}
		return err
	}
	return nil
}

// DeleteUserFile is a function that removes the specified file from the database,
// then calls deleteFile function to delete it.
func DeleteUserFile(userId, fileId int64) error {
	file, err := GetFileByID(userId, fileId)
	if err != nil {
		logging.LogError(err, "file-model/DeleteUserFile : returned error while getting the file")
		return err
	}
	statement, err := db.Prepare("DELETE FROM files WHERE user_id = $1 AND id = $2")
	if err != nil {
		logging.LogError(err, "file-model/DeleteUserFile : preparing statement")
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(userId, fileId)
	if err != nil {
		logging.LogError(err, "file-model/DeleteUserFile : executing statement")
		return err
	}
	err = deleteFile(userId, file.Name, file.Type)
	if err != nil {
		_, execErr := db.Exec("INSERT INTO files(id, name, type, user_id) VALUES ($1, $2, $3, $4)",
			file.ID, file.Name, file.Type, file.UserID)
		if execErr != nil {
			logging.LogError(execErr, "file-model/DeleteUserFile : executing statement at returned error from deleting the file")
			return execErr
		}
		return err
	}
	return nil
}

// SortFileRows is a function that sorts a specified input file according to a specified
// sorting column, then outputs the result to a newly created output file that has the specified name.
func SortFileRows(userId int64, fileId int64, sortingColumn int, outputFileName string, accountId int64) (*File, error) {
	file, err := GetFileByID(userId, fileId)
	if err != nil {
		logging.LogError(err, "file-model/SortFileRows : returned error while getting the file")
		return nil, err
	}
	result, err := sorting.Sort(file.Data, sortingColumn)
	if err != nil {
		return nil, err
	}
	outputFile := File{
		Name:   outputFileName,
		Type:   "output",
		Data:   result,
		UserID: accountId,
	}
	err = outputFile.CreateUserFile()
	if err != nil {
		logging.LogError(err, "file-model/SortFileRows : returned error while creating the output file")
		return nil, err
	}
	return &outputFile, nil
}
