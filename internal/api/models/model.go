package models

import (
	"database/sql"
	"fmt"

	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/api/config"
	_ "github.com/lib/pq"
)

// DB is a pointer to the PostgreSQL database that opens a connection with the configured database
// to be used later in the project.
var DB *sql.DB

// Content holds the columns that represent the columns of a csv file, used to format the data correctly.
type Content struct {
	Column0 string `json:"column_0"`
	Column1 string `json:"column_1"`
	Column2 string `json:"column_2"`
}

// init function initializes a connection to the configured database and reference it in the variable DB.
func init() {
	DB = config.ConnectToDB(config.Host, config.Port, config.User, config.Password, config.DBname)
}

// CheckRecords is a function that looks through the database for a match and returns
// true if it exists, false if not, or an error in case some error emerges.
// The first value of fieldNames and values are the field and the value we're searching for a match to,
// meanwhile the rest of the fieldNames and values are just to add more conditions to the query.
func CheckRecords(tableName string, fieldNames []string, values ...interface{}) (bool, error) {
	var field string
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = $1", fieldNames[0], tableName, fieldNames[0])
	if len(fieldNames) > 1 {
		for i := 1; i < len(fieldNames); i++ {
			query += fmt.Sprintf(" AND %s = $%d", fieldNames[i], i+1)
		}
	}
	statement, err := DB.Prepare(query)
	if err != nil {
		return false, err
	}
	defer statement.Close()
	_ = statement.QueryRow(values...).Scan(&field)
	if field == values[0] {
		return true, nil
	} else {
		return false, nil
	}
}
