package user_database_service

import (
	"SmartSearch/internal/repository/user_index"
	"fmt"
	"log"
	"reflect"
)

func Update(userId string, indexName string, doc map[string]interface{}) error {
	if doc == nil {
		log.Println("Error: document cannot be nil")
		return fmt.Errorf("document cannot be nil")
	}
	log.Println("Extract data:", doc)

	fieldTypes := make(map[string]string)
	for key, value := range doc {
		if value == nil {
			fieldTypes[key] = "nil"
			continue
		}
		fieldTypes[key] = reflect.TypeOf(value).String()
	}

	err := user_index.UpdateUserIndexStructure(userId, indexName, fieldTypes)
	if err != nil {
		log.Println("Error updating UserIndex:", err)
		return err
	}

	return nil
}
