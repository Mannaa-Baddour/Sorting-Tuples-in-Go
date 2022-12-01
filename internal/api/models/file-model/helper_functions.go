package file_model

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/api/models"
	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/logging"
	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/server"
)

// ParseText is a function that reads the content of the Textarea HTML element
// and format each line as the columns specified in the struct Content,
// then returns a slice of all those Content.
func ParseText(text string) []models.Content {
	lines := strings.Split(text, "\r\n")
	var data []models.Content
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			break
		}
		var content models.Content
		columns := []*string{&content.Column0, &content.Column1, &content.Column2}
		values := strings.Split(line, ",")
		for i, value := range values {
			value = strings.TrimSpace(value)
			*columns[i] = value
		}
		data = append(data, content)
	}
	return data
}

// insertHeaderToContent adds the column names that the csv file starts with to the data.
func insertHeaderToContent(content [][]string) [][]string {
	columnsNames := []string{"column_0", "column_1", "column_2"}
	if content == nil {
		content = append(content, columnsNames)
	} else if !reflect.DeepEqual(content[0], columnsNames) {
		content = append(content[:1], content[:]...)
		content[0] = columnsNames
	}
	return content
}

// contentToStrings transforms the slice of Content into a 2D slice of string
// to be used to write the data to the csv file.
func contentToStrings(content []models.Content) [][]string {
	var stringContent [][]string
	if content != nil {
		for _, data := range content {
			line := []string{data.Column0, data.Column1, data.Column2}
			stringContent = append(stringContent, line)
		}
	}
	stringContent = insertHeaderToContent(stringContent)
	return stringContent
}

// createFile is a function that creates a new csv file according to the parameters it gets
// which specify the user folder, whether the file is in input or output, its name, and its content.
func createFile(userId int64, fileName string, fileType string, content []models.Content) error {
	userFolder := strconv.FormatInt(userId, 10)
	path := fmt.Sprintf("%s%s/%s/", server.Path, userFolder, fileType)
	err := os.Chdir(path)
	if err != nil {
		logging.LogError(err, "file-model/createFile : changing directory to the specified path")
		return err
	}
	csvFile, err := os.Create(fileName + ".csv")
	if err != nil {
		logging.LogError(err, "file-model/createFile : creating the specified file")
		return err
	}
	defer csvFile.Close()
	csvWriter := csv.NewWriter(csvFile)
	stringContent := contentToStrings(content)
	err = csvWriter.WriteAll(stringContent)
	if err != nil {
		logging.LogError(err, "file-model/createFile : writing data to the ccv file")
		return err
	}
	return nil
}

// readFile reads the content of a csv file which its path is specified by the parameters
// where userId is the user folder, fileType specifies whether it's in input or output,
// and fileName specifies the name of this file.
func readFile(userId int64, fileName string, fileType string) ([]models.Content, error) {
	filePath := fmt.Sprintf("%s%d/%s/%s.csv", server.Path, userId, fileType, fileName)
	csvFile, err := os.Open(filePath)
	if err != nil {
		logging.LogError(err, "file-model/readFile : opening file for read")
		return nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	lines, err := reader.ReadAll()
	if err != nil {
		logging.LogError(err, "file-model/readFile : reading data from the csv file")
		return nil, err
	}
	var allContent []models.Content
	for _, line := range lines {
		if !reflect.DeepEqual(line, []string{"column_0", "column_1", "column_2"}) {
			var content models.Content
			content.Column0 = line[0]
			content.Column1 = line[1]
			content.Column2 = line[2]
			allContent = append(allContent, content)
		}
	}
	return allContent, nil
}

// updateFile modifies the content of the file which its path is specified by the function parameters.
func updateFile(userId int64, newFileName string, oldFileName string, fileType string, content []models.Content) error {
	if oldFileName != newFileName {
		err := deleteFile(userId, oldFileName, fileType)
		if err != nil {
			return err
		}
	}
	filePath := fmt.Sprintf("%s%d/%s/%s.csv", server.Path, userId, fileType, newFileName)
	csvFile, err := os.Create(filePath)
	if err != nil {
		logging.LogError(err, "file-model/updateFile : creating a new csv file")
		return err
	}
	defer csvFile.Close()
	csvWriter := csv.NewWriter(csvFile)
	stringContent := contentToStrings(content)
	err = csvWriter.WriteAll(stringContent)
	if err != nil {
		logging.LogError(err, "file-model/updateFile : writing data to the csv file")
		return err
	}
	return nil
	// NOTE: This function is just like CreateFile function, with just fileType passed
	// as an argument, consider merging them unless we need to add more functionality to this
	// i.e. storing data in case of update abortion.
}

// deleteFile removes the file which its path is specified by the parameters where
// userId is the folder name, fileType is the folder its in, and the fileName which is its name
func deleteFile(userId int64, fileName string, fileType string) error {
	filePath := fmt.Sprintf("%s%d/%s/%s.csv", server.Path, userId, fileType, fileName)
	err := os.Remove(filePath)
	if err != nil {
		logging.LogError(err, "file-model/deleteFile : removing the file")
		return err
	}
	return nil
}

// DirectoryNotEmpty checks if the specified folder has files or not, returns true if it has,
// or false otherwise, in other words: true means directory not empty, false means it's empty.
// It's used to help manage the data displayed in the webpage.
func DirectoryNotEmpty(userId, folderName string) (bool, error) {
	folderPath := fmt.Sprintf("%s%s/%s", server.Path, userId, folderName)
	folder, err := os.Open(folderPath)
	if err != nil {
		logging.LogError(err, "file-model/DirectoryNotEmpty : opening folder")
		return false, err
	}
	defer folder.Close()
	_, err = folder.Readdirnames(1)
	if err != io.EOF {
		return true, nil
	}
	return false, nil
}
