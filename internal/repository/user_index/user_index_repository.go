package user_index

import (
	"SmartSearch/internal/models/database"
	"SmartSearch/internal/utils"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func FindByUserAndIndex(userID, indexName string) (*database.UserIndex, error) {
	var userIndex database.UserIndex

	if userID == "" || indexName == "" {
		return nil, fmt.Errorf("userID and indexName cannot be empty")
	}

	err := utils.DB.
		Model(&database.UserIndex{}).
		Where("user_id = ? AND index_name = ?", userID, indexName).
		First(&userIndex).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user index not found for user '%s' and index '%s'", userID, indexName)
		}
		return nil, fmt.Errorf("failed to query user index: %w", err)
	}

	return &userIndex, nil
}

func CreateUserIndex(userID string, indexName string) (*database.UserIndex, error) {
	userIndex := &database.UserIndex{
		UserId:         userID,
		IndexName:      indexName,
		IndexStructure: database.JSONB{},
	}

	if err := utils.DB.Create(userIndex).Error; err != nil {
		return nil, err
	}

	return userIndex, nil
}

func UpdateUserIndexStructure(userID string, indexName string, structure interface{}) error {
	var jsonData database.JSONB

	if structure == nil {
		jsonData = database.JSONB{}
	} else {
		b, err := json.Marshal(structure)
		if err != nil {
			return fmt.Errorf("failed to marshal structure: %w", err)
		}

		var tmp map[string]interface{}
		if err := json.Unmarshal(b, &tmp); err != nil {
			return fmt.Errorf("failed to unmarshal into JSONB: %w", err)
		}
		jsonData = tmp
	}

	result := utils.DB.Model(&database.UserIndex{}).
		Where("user_id = ? AND index_name = ?", userID, indexName).
		Update("index_structure", jsonData)

	if result.Error != nil {
		return fmt.Errorf("failed to update index structure: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
