package user_database_service

import (
	"SmartSearch/internal/models/database"
	"SmartSearch/internal/repository/user_index"
	"errors"
)

func GetIndexStructure(userId string, indexName string) (database.JSONB, error) {
	if indexName == "" {
		return nil, errors.New("index_service index_name cannot be empty")
	}

	userIndex, err := user_index.FindByUserAndIndex(userId, indexName)
	if err != nil {
		return nil, err
	}

	return userIndex.IndexStructure, nil
}
