package user_database_service

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func convertRowsToStringMap(rows *sql.Rows) ([]map[string]interface{}, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var allDocs []map[string]interface{}

	for rows.Next() {
		vals := make([]interface{}, len(cols))
		valPtrs := make([]interface{}, len(cols))
		for i := range vals {
			valPtrs[i] = &vals[i]
		}

		if err := rows.Scan(valPtrs...); err != nil {
			return nil, err
		}

		doc := make(map[string]interface{})
		for i, col := range cols {
			switch v := vals[i].(type) {
			case []byte:
				doc[col] = string(v)
			case nil:
				doc[col] = nil
			default:
				doc[col] = v
			}
		}

		allDocs = append(allDocs, doc)
	}

	return allDocs, nil
}

func Extract(driverName string, dataSourceName string, tableName string) ([]map[string]interface{}, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println("Database connection error: ", err)
		return nil, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Database close connection error: ", err)
		}
	}(db)

	rows, err := db.Query("SELECT * FROM " + tableName)
	if err != nil {
		log.Println("Could not fetch the data from the table: ", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Rows close error: ", err)
		}
	}(rows)
	log.Println("Database rows are extracted")

	docs, err := convertRowsToStringMap(rows)
	if err != nil {
		return nil, err
	}
	log.Println("Rows are converted to a string map")

	return docs, nil
}
